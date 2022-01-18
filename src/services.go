package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var endpoint = "https://api.nasa.gov/planetary/apod?"
var formatDate = "date="
var formatKey = "&api_key="

func ObtemRecursosPorData(data string) []byte {
	var url = endpoint + (formatDate + data) + (formatKey + os.Getenv("nasa_key"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
