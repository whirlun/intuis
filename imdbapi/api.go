package imdbapi

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/astaxie/beego"
	"errors"
	"io/ioutil"
	"encoding/json"
)

type APIRequest struct {
	Key string
	Id int
	Type string
}

type APIResponse struct {
	Imdb_id string  `json:"imdb_id"`
	Title string    `json:"title"`
	Year string     `json:"year"`
	Rated string    `json:"rated"`
	Released string `json:"released"`
	Runtime string  `json:"runtime"`
	genre string    `json:"genre"`
	Director string `json:"director"`
	Writers string  `json:"writers"`
	Actors  string  `json:"actors"`
	Plot  string    `json:"plot"`
	Country string  `json:"country"`
	Language string `json:"language"`
	Metascore string`json:"metascore"`
	Poster string   `json:"poster"`
	Rating string   `json:"rating"`
	Votes  string   `json:"votes"`
	Budget string   `json:"budget"`
	OpeningWd string`json:"opening_weekend"`
	Gross string    `json:"gross"`
	Production string`json:"production"`
	Type string     `json:"type"`
	Status string   `json:"status"`
}

const APIROOT = "http://imdbapi.net/api"

func (r *APIRequest) Get(id int) error  {
	req := APIRequest{
		Key: beego.AppConfig.String("imdbapi::apikey"),
		Id: id,
		Type: "json",
	}
	res, err := http.Post(APIROOT,"application/x-www-form-urlencoded",
						strings.NewReader(fmt.Sprintf("key=%s&id=&d&type=%s",req.Key,req.Id,req.Type)))
	if err != nil {
		return errors.New("apierror")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("apierror")
	}
	apires := APIResponse{}
	err = json.Unmarshal(body,&apires)
	if err != nil {
		return errors.New("apierror")
	}
}