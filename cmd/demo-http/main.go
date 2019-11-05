package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "hello world")
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET params were:", r.URL.Query())
	log.Println("GET params a = ", r.URL.Query().Get("a"))

	var request_data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&request_data)
	log.Println("request json body : ", request_data)

	request_data["data"] = "ok"
	bytesResult, err := json.Marshal(request_data)
	if err != nil {
		log.Fatalln(err)
	}
	//w.Write([]byte("bye bye ,this is v2 httpServer"))
	w.Write(bytesResult)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/post", PostHandler)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8011", nil))
}
