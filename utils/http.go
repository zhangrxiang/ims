package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	client = http.Client{}
	TUrl   = url.URL{
		Scheme: "http",
		Host:   "127.0.0.1:8081",
	}
)

func Get(path string) []byte {
	TUrl.Path = path
	resp, err := client.Get(TUrl.String())

	var all []byte
	if err != nil {
		log.Fatalln("Get", err)
		return all
	}
	log.Println(TUrl.Query().Encode())
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("ReadAll", err)
		}
		log.Println("\n" + string(all))
		return all
	}
	return all
}

func Post(path string, params interface{}) interface{} {

	TUrl.Path = path
	var b []byte
	var err error
	switch params.(type) {
	case string:
		b = []byte(params.(string))
	default:
		b, err = json.Marshal(params)
		if err != nil {
			log.Fatalln("ReadAll", err)
		}
	}
	log.Println("\n" + string(b))

	resp, err := client.Post(TUrl.String(), "application/x-www-form-urlencoded", bytes.NewReader(b))
	data := map[string]interface{}{}
	if err != nil {
		log.Fatalln("Get", err)
		return data
	}
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("ReadAll", err)
		}
		log.Println("\n" + string(all))
		err = json.Unmarshal(all, &data)
		if err != nil {
			log.Fatalln("Unmarshal", err)
		}
		return data
	}
	return data
}
