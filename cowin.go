package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Sessions struct {
	Sessions []Session `json:"sessions"`
}

type Session struct {
	CenterId          int      `json:"center_id"`
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	StateName         string   `json:"state_name"`
	DistrictName      string   `json:"district_name"`
	BlockName         string   `json:"block_name"`
	Pincode           int      `json:"pincode"`
	From              string   `json:"from"`
	To                string   `json:"to"`
	Lat               int      `json:"lat"`
	Long              int      `json:"long"`
	FeeType           string   `json:"fee_type"`
	SessionId         string   `json:"session_id"`
	Date              string   `json:"date"`
	AvailableCapacity float32  `json:"available_capacity"`
	Fee               string   `json:"fee"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

func callApi() {
	main_url := "https://cdn-api.co-vin.in/api"
	apt_url := "/v2/appointment/sessions/public/findByDistrict"
	query := "?district_id=294&date=05-05-2021"

	url := fmt.Sprintf("%s%s%s", main_url, apt_url, query)

	res, err := http.Get(url)

	// check for response error
	if err != nil {
		fmt.Print(err)
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	m := Sessions{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Printf("\n%s: %s", err, data)
		return
	}

	cnt := 0
	for i := 0; i < len(m.Sessions); i++ {
		sn := m.Sessions[i]

		if sn.MinAgeLimit == 18 {
			fmt.Printf("%s | %s | %d | %s\n", sn.Name, sn.Date, sn.AvailableCapacity, sn.Slots)
			cnt++
		}
	}

	if cnt == 0 {
		fmt.Print("\nNot Available")
	}
}

func main() {
	for {
		time.Sleep(60 * time.Second)
		go callApi()
	}
}
