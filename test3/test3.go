package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type data struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type datas []data

func main() {
	response, err := GetListData()
	if err != nil {
		log.Println("Error : ", err.Error())
		panic(err)
	}
	log.Println(JsonEncode(response))
}

func GetListData() (result datas, err error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	request, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return
	}
	response, err := client.Do(request.Request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 300 {
		err = errors.New("something error : " + response.Status)
		return
	}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

func JsonEncode(obj interface{}) string {
	json, _ := json.MarshalIndent(obj, "", "  ")
	return string(json)
}
