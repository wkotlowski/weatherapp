package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var city string

func main() {
	fmt.Println("Welcome to my simple weather application.")
	checkCity()
	fmt.Println("Weather for", city, "is:")
	checkWeather()
}

func checkCity() {
	fmt.Println("Please, provide your city and hit ENTER:")
	fmt.Scanln(&city)
	fmt.Println("Your city was set to:", city)
}

func checkWeather() {
	resp, err := http.Get(fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=7e8316c662c342c38e880612210612&q=%s&aqi=yes", city))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)
	fmt.Println(stringBody)
}
