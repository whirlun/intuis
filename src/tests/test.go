package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)
const APIROOT = "http://imdbapi.net/api"
type APIRequest struct {
	Key string
	Id string
	Type string
}

func main() {
	req := APIRequest{
		Key: "UZYjbJ6hkt8PgSEsc7bFPxlZmbhr5m",
		Id: "4466494",
		Type: "json",
	}
	res, err := http.Post(APIROOT,"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("key=%s&id=%s&type=%s",req.Key,req.Id,req.Type)))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Println(fmt.Sprintf("key=%s&id=%s&type=%s",req.Key,req.Id,req.Type))
}