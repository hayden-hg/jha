package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	apiUrl = "https://api.openweathermap.org/data/2.5/onecall?lat=%v&lon=%v&appid=%s"
	apiKey = ""
)

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Current struct {
	Temp      float64   `json:"temp"`
	FeelsLike float64   `json:"feels_like"`
	Weather   []Weather `json:"weather"`
}

type WeatherApiResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        Current `json:"current"`
}

type demoApiResponse struct {
	Temp    string    `json:"temp"`
	Weather []Weather `json:"weather"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getWeather", getWeather)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

/*
Write an HTTP server that uses the Open Weather API that exposes an endpoint that takes in lat/long coordinates. This endpoint should return the weather conditions in that area (snow, rain, etc), whether it’s hot, cold, or moderate outside (use your own discretion on what temperature equates to each type), and whether there are any weather alerts in that area along with the weather conditions related to the alert.
*/

func getWeather(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	getUrl := fmt.Sprintf(apiUrl, lat, long, apiKey)

	req, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := http.DefaultClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	weatherResponse := WeatherApiResponse{}
	jsonErr := json.Unmarshal(body, &weatherResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	wtemp := weatherResponse.Current.Temp
	var restemp string

	if wtemp > 80 {
		restemp = "hot"
	} else if wtemp < 40 {
		restemp = "cold"
	} else {
		restemp = "moderate"
	}

	demoResp := demoApiResponse{Temp: restemp, Weather: weatherResponse.Current.Weather}
	jsonResponse, jsonError := json.Marshal(demoResp)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
