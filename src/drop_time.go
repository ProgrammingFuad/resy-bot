package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

func dropTimeService() {
	fmt.Println("Starting dropTimeService")
	const api_key = `ResyAPI api_key="VbWk7s3L4KiK5fzlO7JD3Q5EYolJI7n5"`
	auth_token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NzU2MzM1NjMsInVpZCI6MTk3MTE1NDksImd0IjoiY29uc3VtZXIiLCJncyI6W10sImxhbmciOiJlbi11cyIsImV4dHJhIjp7Imd1ZXN0X2lkIjo0MjM4MzcwNn19.ANkAYgaLpdYuD8KsgI_aG2KK6gzgSiZjafXKb1d754zHA83gLQxPGgntCdn4ZMneYWXuN9ptStmzkbLa5C01lCkAAdx7DtkQMNdSstb5MVrH5IYlrq5wyH4X4eHufs44CRWkPf_V3Vip7DhF3PvVi0FKQ_N_9htL5Dla5TEsmo5H71tY"
	client := httpClient()

	// drop_timew := time.Now().Add(-time.Minute)
	// confirmation_stringw := fmt.Sprintf("The drop time for me is %d:%d:00\n", drop_timew.Hour(), drop_timew.Minute())
	// fmt.Printf(confirmation_stringw + "Ending the scan.\n")

	scheduler := gocron.NewScheduler(time.UTC)

	venue := venues.carbone
	fmt.Println("The venue is:", venue.name)

	next_available_day := scanForNextDayReservationsDrop(client, api_key, auth_token, venue.id)

	fmt.Printf("Next drop day for %s is %s!\n", venue.name, next_available_day)

	scheduler.Cron("1,16,31,46 * * * *").Do(func() {
		fmt.Println("Starting the drop time scan.")
		if isNotifyOptionsEmpty(client, api_key, auth_token, next_available_day, venue.id) {
			// Found the drop time
			drop_time := time.Now().Add(-time.Minute)
			confirmation_string := fmt.Sprintf("The drop time for %s is %d:%d:00\n", venue.name, drop_time.Hour(), drop_time.Minute())
			fmt.Printf(confirmation_string + "Ending the scan.\n")
			writeToFile(confirmation_string)
			scheduler.StopBlockingChan()
			return
		} else {
			// This was not the drop time
			drop_time := time.Now().Add(-time.Minute)
			failure_string := fmt.Sprintf("The drop time for %s is NOT %d:%d:00. Continuing the scan.\n", venue.name, drop_time.Hour(), drop_time.Minute())
			fmt.Println(failure_string)
			// writeToFile(failure_string)
		}
	})

	fmt.Println(`Begininng scan at the 01, 16, 31 and 46 minutes of every hour until "notify_options" is no longer empty meaning the reservations have been dropped.`)
	scheduler.StartBlocking()
}

func scanForNextDayReservationsDrop(client *http.Client, api_key string, auth_token string, venue_id string) string {

	// Get the current time
	next_available_day := time.Now()
	for {
		formatted_next_available_day := next_available_day.Format("2006-01-02")
		formatted_next_next_available_day := next_available_day.Add(time.Hour * 24).Format("2006-01-02")
		if isNotifyOptionsEmpty(client, api_key, auth_token, formatted_next_available_day, venue_id) &&
			isNotifyOptionsEmpty(client, api_key, auth_token, formatted_next_next_available_day, venue_id) {
			// Found the next drop day
			return formatted_next_available_day
		} else {
			// restart the loop and try the next day
			fmt.Println(formatted_next_available_day, "is not the next drop day. Trying the next day.")
			next_available_day = next_available_day.Add(time.Hour * 24)
		}
	}
}

func isNotifyOptionsEmpty(client *http.Client, api_key string, auth_token string, date string, venue_id string) bool {
	url := fmt.Sprintf("https://api.resy.com/4/find?lat=0&long=0&day=%s&party_size=2&venue_id=%s", date, venue_id)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("Error occured. %+v", err)
	}

	// Only Need Auth Key
	req.Header.Add("Authorization", api_key)
	req.Header.Add("x-resy-auth-token", auth_token)

	// fmt.Println("Request URL is: ", req.URL)
	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	var data FindReservationResponse
	responseErr := json.Unmarshal(body, &data)
	if responseErr != nil {
		// handle error
		fmt.Println("FindReservationResponse json.Unmarshal error found!")
	}

	// Notify options
	if data.Results != nil && len(data.Results.Venues) > 0 && len(data.Results.Venues[0].Venue.NotifyOptions) == 0 {
		return true
	} else {
		return false
	}
}

func writeToFile(text string) {
	// Open the file with the os.O_APPEND flag, which allows you to
	// append to the file, and the os.O_WRONLY flag, which allows you
	// to write to the file. 0644 is file permissions.
	file, err := os.OpenFile("Drop_Time.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Couldnt open Drop_Time.txt. Error:", err)
		fmt.Println("Exiting write function, scan execution will continue")
	}
	defer file.Close()

	// Write the line of text to the file.
	if _, err := file.Write([]byte(text)); err != nil {
		fmt.Println("Error writing to Drop_Time.txt. Exiting write function, scan execution will continue")
	} else {
		fmt.Println("Stored correct drop time in Drop_Time.txt.")
	}
}
