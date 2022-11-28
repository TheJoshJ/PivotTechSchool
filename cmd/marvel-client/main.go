package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var ErrNotFound = errors.New("Not found")

const baseURL = "https://gateway.marvel.com"

const (
	k  = "d516e7d9e1a3fc70cb300928603e3ab940270b03"
	pk = "1f4b8f992212e1e3fc00e6b7779e54ed"
)

type MarvelResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int          `json:"offset"`
		Limit   int          `json:"limit"`
		Total   int          `json:"total"`
		Count   int          `json:"count"`
		Results []characters `json:"results"`
	} `json:"data"`
}

type characters struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Modified    string `json:"modified"`
	Thumbnail   struct {
		Path      string `json:"path"`
		Extension string `json:"extension"`
	}
	Comics struct {
		Available int `json:"available"`
	}
}

func addQueryAuth(q url.Values) url.Values {
	// add auth parameters to query
	ts := time.Now().UnixNano()
	hashInput := fmt.Sprintf("%d%s%s", ts, k, pk)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(hashInput)))
	q.Add("ts", fmt.Sprintf("%d", ts))
	q.Add("apikey", pk)
	q.Add("hash", hash)

	return q
}

func makeRequest(url string) []byte {
	// Fill in the correct arguments here
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Leave this part alone to handle adding the necessary auth parameters
	q := addQueryAuth(req.URL.Query())
	req.URL.RawQuery = q.Encode()

	// here, you'd need to actually send the request, read the response, and return it
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return nil
	}

	return body
}

type CharacterDetail struct {
	ID           int    `json:"id"`
	Description  string `json:"name"`
	ThumbnailURL string
	ComicCount   int
}

type Client struct{}

func (c *Client) GetFirst25() []string {

	var MarRes MarvelResponse
	var First25 []string

	res := makeRequest(baseURL + "/v1/public/characters?limit=25")

	err := json.Unmarshal(res, &MarRes)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, character := range MarRes.Data.Results {
		First25 = append(First25, character.Name)
	}
	return First25
}

func (c *Client) GetCharacterDetail(name string) (CharacterDetail, error) {
	var MarRes MarvelResponse
	var charDet CharacterDetail

	res := makeRequest(baseURL + "/v1/public/characters?name=" + name)
	err := json.Unmarshal(res, &MarRes)
	if err != nil {
		return charDet, nil
	}

	if len(MarRes.Data.Results) == 0 {
		return CharacterDetail{}, ErrNotFound
	}

	for _, character := range MarRes.Data.Results {
		charDet.ID = character.ID
		charDet.Description = character.Description
		charDet.ComicCount = character.Comics.Available
		charDet.ThumbnailURL = (character.Thumbnail.Path + "." + character.Thumbnail.Extension)
	}

	return charDet, nil
}

func main() {
	c := Client{}
	log.Println(c.GetFirst25())

	/*
		OUTPUT:
		[3-D Man A-Bomb (HAS) A.I.M. Aaron Stack Abomination (Emil Blonsky) Abomination (Ultimate) Absorbing Man
		Abyss Abyss (Age of Apocalypse) Adam Destine Adam Warlock Aegis (Trey Rollins) Aero (Aero)
		Agatha Harkness Agent Brand Agent X (Nijo) Agent Zero Agents of Atlas Aginar
		Air-Walker (Gabriel Lan) Ajak Ajaxis Akemi Alain Albert Cleary]
	*/
}
