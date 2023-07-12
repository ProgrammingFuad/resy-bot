package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	venue_name := os.Getenv("VENUE_NAME")
	fmt.Println("VENUE_NAME:", venue_name)

	auth_Token_owner := os.Getenv("AUTH_TOKEN_OWNER")
	fmt.Println("AUTH_TOKEN_OWNER:", auth_Token_owner)

	venue := venues_dict[venue_name]
	auth_token := owners_dict[auth_Token_owner]

	snipeReservation(venue, auth_token)

	fmt.Println("Script is done!")
	//dropTimeService()
}
