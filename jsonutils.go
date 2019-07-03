package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type jsonType map[string]interface{}

func jsonGet(url string, result interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &result)
}

func jsonPost(url string, content jsonType, result interface{}) {
	contentEncoded, err := json.Marshal(content)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(contentEncoded))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &result)
}
