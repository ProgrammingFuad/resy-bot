package main

var venues = Venues{
	carbone:        Venue{id: "6194", name: "Carbone", days_out: 30, drop_time: DropTime{10, 14, 59, 500}},
	i_sodi:         Venue{id: "443", name: "I Sodi", days_out: 13, drop_time: DropTime{23, 59, 59, 500}},
	don_angie:      Venue{id: "1505", name: "Don Angie", days_out: 6, drop_time: DropTime{8, 59, 59, 500}},
	bocca_di_bacco: Venue{id: "59685", name: "Bocca Di Bacco", days_out: 10, drop_time: DropTime{10, 14, 59, 500}},
	via_corota:     Venue{id: "2567", name: "Via Corota", days_out: 30, drop_time: DropTime{9, 59, 59, 500}},
	misi:           Venue{id: "3015", name: "Misi", days_out: 27, drop_time: DropTime{9, 59, 59, 500}},
	carbone_miami:  Venue{id: "43225", name: "Carbone Miami", days_out: 30, drop_time: DropTime{10, 14, 59, 500}},
}

var venues_dict = map[string]Venue{
	"carbone":        venues.carbone,
	"i_sodi":         venues.i_sodi,
	"don_angie":      venues.don_angie,
	"bocca_di_bacco": venues.bocca_di_bacco,
	"via_corota":     venues.via_corota,
	"misi":           venues.misi,
	"carbone_miami":  venues.carbone_miami,
}

var owners_dict = map[string]string{
	"fuad":  "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NzU3OTIwOTUsInVpZCI6Mzc4MzMwMDYsImd0IjoiY29uc3VtZXIiLCJncyI6W10sImxhbmciOiJlbi11cyIsImV4dHJhIjp7Imd1ZXN0X2lkIjoxMjY0MTEyMDV9fQ.AJGrAow3gtLQf6UIDwuB-l-Or82FJAUI6n7SxesvrM88TPcT5D-GZH4qW4p5OKfxHq2ZTtPcAryNM0Dn25zyFy0XACOWsya5U2Rl9KH9qyyFp88A__e8yLcCad66Zzqsd8MSETPfkk5uI-w8VmrZgTgxfIymGeRRIIRoEjiXfwxBWR3y",
	"aaron": "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NzU2MzM1NjMsInVpZCI6MTk3MTE1NDksImd0IjoiY29uc3VtZXIiLCJncyI6W10sImxhbmciOiJlbi11cyIsImV4dHJhIjp7Imd1ZXN0X2lkIjo0MjM4MzcwNn19.ANkAYgaLpdYuD8KsgI_aG2KK6gzgSiZjafXKb1d754zHA83gLQxPGgntCdn4ZMneYWXuN9ptStmzkbLa5C01lCkAAdx7DtkQMNdSstb5MVrH5IYlrq5wyH4X4eHufs44CRWkPf_V3Vip7DhF3PvVi0FKQ_N_9htL5Dla5TEsmo5H71tY",
	"steve": "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NzY2Njc2NzYsInVpZCI6MzgxMzk5MTcsImd0IjoiY29uc3VtZXIiLCJncyI6W10sImxhbmciOiJlbi11cyIsImV4dHJhIjp7Imd1ZXN0X2lkIjoxMjY5OTY2NDh9fQ.AZA0iPw62uS16hOOQbm2R-KowhG1xubkDq_HRZUgkeXj8iC_Ei2utNtCnItvm10L5bGehS-l7VGitNYuenS6PfEXAY2eLIcXE12cEOeVWUVQC6qC7lGFk0Qujv5gKjBfCe6hRoYJkbK8kcUUErGq5JovGbPrKQEGW0ba9moyIUDgkI5b",
	"john":  "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NzY2ODcwMDIsInVpZCI6Mzc4NzkwODEsImd0IjoiY29uc3VtZXIiLCJncyI6W10sImV4dHJhIjp7Imd1ZXN0X2lkIjoxMjY1MDc4MzZ9fQ.AEhv7BWui7ErMd4pCHT0cUOSzFwWLyH9MQkJQtYwWkaFgefMw6KlalKNH_m43cdtt9SDDlJlPLssii9pYdrpJEhfAEwRboQfWlUNeGLXp5t1IDm2fmV07xr52ZyqM4g_nKFGGI-WBeWzlV3aCLikyOPXvkLoxImG_kOA847hrDeRlli0",
	"jack":  "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NzY2OTU5NzMsInVpZCI6MzgyMDczMDUsImd0IjoiY29uc3VtZXIiLCJncyI6W10sImV4dHJhIjp7Imd1ZXN0X2lkIjoxMjcxMTUyNDN9fQ.AKJ_UheEUt1ukobS3uPCDBn6DXfe14gZXH5U76rpofgt4_K2q4TATPCOF5gjKsDVQ1SEjOdp-w1JC3tJwfsMhEDaAKxXhV4sLkqQYhnktbCitN7anYJaqQmWecMjUQiCmQeJA0Pe8ZzrYZ2uRYEYFV112fFzg2WSciCTl8PxXGxR8YLz",
}

type DropTime struct {
	hour        int
	minute      int
	second      int
	millisecond int
}

type Venue struct {
	id        string
	name      string
	days_out  int
	drop_time DropTime
}

type Venues struct {
	carbone        Venue
	i_sodi         Venue
	don_angie      Venue
	bocca_di_bacco Venue
	via_corota     Venue
	misi           Venue
	carbone_miami  Venue
}

/*
	// Carbone 7:00pm - 8:30pm
	var config_tokens = [4]string{
		"rgs://resy/6194/1808469/3/2023-01-26/2023-01-26/19:00:00/2/Indoor Dining",
		"rgs://resy/6194/1808469/3/2023-01-26/2023-01-26/19:30:00/2/Indoor Dining",
		"rgs://resy/6194/1808469/3/2023-01-26/2023-01-26/20:00:00/2/Indoor Dining",
		"rgs://resy/6194/1808469/3/2023-01-26/2023-01-26/20:30:00/2/Indoor Dining",
	}

	// Bocca di Bacco 7:00pm - 8:30pm
	var config_tokens = [4]string{
		"rgs://resy/59685/1648000/2/2023-01-26/2023-01-26/19:00:00/2/Main Dining",
		"rgs://resy/59685/1648000/2/2023-01-26/2023-01-26/19:30:00/2/Main Dining",
		"rgs://resy/59685/1648000/2/2023-01-26/2023-01-26/20:00:00/2/Main Dining",
		"rgs://resy/59685/1648000/2/2023-01-26/2023-01-26/20:30:00/2/Main Dining",
	}

	// Don Angie 9:00pm - 10:00pm
	var config_tokens = [4]string{
		"rgs://resy/1505/1108282/2/2023-01-03/2023-01-03/21:00:00/2/Indoor",
		"rgs://resy/1505/1108282/2/2023-01-03/2023-01-03/21:30:00/2/Indoor",
		"rgs://resy/1505/1108282/2/2023-01-03/2023-01-03/22:00:00/2/Indoor",
		"rgs://resy/1505/1108282/2/2023-01-03/2023-01-03/22:30:00/2/Indoor",
	}

	// I Sodi 9:00pm - 10:00pm
	var config_tokens = [4]string{
		"rgs://resy/443/1164715/2/2023-01-11/2023-01-11/21:00:00/2/Dinning Room",
		"rgs://resy/443/1164715/2/2023-01-11/2023-01-11/21:30:00/2/Dinning Room",
		"rgs://resy/443/1164715/2/2023-01-11/2023-01-11/22:00:00/2/Dinning Room",
		"rgs://resy/443/1164715/2/2023-01-11/2023-01-11/22:30:00/2/Dinning Room",
	}
*/
