package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var city string

func main() {
	fmt.Println("Welcome to my simple weather application.")
	checkCity()
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

	byteValue, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatalln(err)
	}

	var weather Response

	json.Unmarshal(byteValue, &weather)
	fmt.Println("Current weather is:")
	fmt.Println("Temperature in Celcius:", weather.Current.TempC)
	fmt.Println("Temperature in Fahrenheit:", weather.Current.TempF)
	fmt.Println("Condition:", weather.Current.Condition.Text)
}
