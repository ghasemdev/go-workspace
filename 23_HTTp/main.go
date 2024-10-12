package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	response, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("ERROR: can't call httpbin.org")
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)

	job := &Job{
		User:   "Amir",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer // in memory io reader
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("ERROR: can't encode job - %s", err)
	}

	response2, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("ERROR: can't call httpbin.org")
	}

	defer response2.Body.Close()

	io.Copy(os.Stdout, response2.Body)
}
