package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

/**
参考
https://golang.org/pkg/net/http/
http://polyglot.ninja/golang-making-http-requests/
*/
var baseUrl = "http://127.0.0.1:8011"

func TestHttpGet(t *testing.T) {
	resp, err := http.Get(baseUrl + "/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	//def resp.Close()
	//resp.Body
}

func TestHttpPost(t *testing.T) {
	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(baseUrl+"/post?a=1", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)

	log.Println("result: ", result)
	log.Println("result data = ", result["data"])
}
