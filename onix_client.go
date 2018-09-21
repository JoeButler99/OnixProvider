package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Handy!! -  https://mholt.github.io/curl-to-go/

// CheckOnixError provides error checking functions for OnixClient. This is because it may need handling with respect to the terraform provider.
func CheckOnixError(err error) {
	if err != nil {
		fmt.Println("[OnixError] ", err)
		log.Printf("[OnixError] %s", err)
		panic(err) // <-- TODO Check how this should really be done for a terraform provider.
	}
}

type OnixClient struct {
	BaseURL string
	//ClientId, Username, Password string
}

// TODO - OnixClient may want to cache some of the GET requests to improve TF speed. (Or implent itemtype lookup by key)

type OnixItemType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OnixItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Itemtype    int    `json:"itemtype"` // TODO - Should this link to the Type above?
}

type OnixApiResponseItem struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OnixApiPutDeleteResponse struct {
	Changed   bool   `json:"changed"`
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	Operation string `json:"operation"`
}

type OnixApiGetResponse struct {
	Items []OnixApiResponseItem
}

func (o *OnixClient) Put(elementName, key string, payload io.Reader) (*OnixApiPutDeleteResponse, error) {

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", o.BaseURL, elementName, key), payload)
	CheckOnixError(err)

	req.Header.Set("Contenttype", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	CheckOnixError(err)
	defer resp.Body.Close()

	onixResponse := new(OnixApiPutDeleteResponse)
	CheckOnixError(json.NewDecoder(resp.Body).Decode(onixResponse))

	return onixResponse, err
}

func (o *OnixClient) Delete(elementName, key string) (*OnixApiPutDeleteResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", o.BaseURL, elementName, key), nil)
	CheckOnixError(err)

	resp, err := http.DefaultClient.Do(req)
	CheckOnixError(err)
	defer resp.Body.Close()

	onixResponse := new(OnixApiPutDeleteResponse)
	CheckOnixError(json.NewDecoder(resp.Body).Decode(onixResponse))

	return onixResponse, err
}

func (o *OnixClient) GetItemType(elementName, key string) (OnixItemType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", o.BaseURL, elementName), nil)
	CheckOnixError(err)

	resp, err := http.DefaultClient.Do(req)
	CheckOnixError(err)
	defer resp.Body.Close()

	onixResponse := new(OnixApiGetResponse)
	CheckOnixError(json.NewDecoder(resp.Body).Decode(onixResponse))

	for _, item := range onixResponse.Items {
		if item.Key == key {
			return OnixItemType{
				Id:          item.Id,
				Name:        item.Name,
				Description: item.Description,
			}, nil
		}
	}

	return OnixItemType{}, err
}
