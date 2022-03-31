package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Labels2map(labels []Label) map[string]string {
	kv := make(map[string]string)
	if labels == nil {
		return kv
	}
	for _, label := range labels {
		kv[label.Key] = label.Value
	}
	return kv
}

func HttpGet(uri string) (string, error) {
	// uri := "https://httpbin.org/delay/3"
	client := &http.Client{}
	client.Timeout = time.Millisecond * 100
	resp, err := client.Get(uri)
	if err != nil {
		log.Fatalf("http.Get() failed with '%s'\n", err)
	}
	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
		return "", err
	}
	return string(d), nil
}

func HttpPost(uri string, data string) (string, error) {
	client := &http.Client{}
	client.Timeout = time.Second * 15

	// uri := "https://httpbin.org/post"
	body := bytes.NewBufferString(data)
	resp, err := client.Post(uri, "application/json; charset=utf-8", body)
	if err != nil {
		log.Fatalf("client.Post() failed with '%s'\n", err)
		return "", err
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("http.Get() failed with '%s'\n", err)
		return "", err
	}
	return string(d), nil
}

func HttpPut(uri string, data string) {
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
	}

	client := &http.Client{}
	client.Timeout = time.Second * 15

	body := bytes.NewBuffer(d)
	req, err := http.NewRequest(http.MethodPut, uri, body)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
	}

	defer resp.Body.Close()
	d, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	fmt.Printf("Response status code: %d, text:\n%s\n", resp.StatusCode, string(d))
}
