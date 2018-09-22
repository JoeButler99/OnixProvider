package main

import (
	"bytes"
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
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (oit *OnixItemType) GetJsonBytesReader() *bytes.Reader {
	jsonBytes, err := json.Marshal(oit)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(jsonBytes)
}

type OnixItem struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"` // TODO - Should this link to the Type above?
	Status      int    `json:"status"`
}

func (oi *OnixItem) GetJsonBytesReader() *bytes.Reader {
	jsonBytes, err := json.Marshal(oi)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(jsonBytes)
}

type OnixApiPutDeleteResponse struct {
	Changed    bool        `json:"changed"`
	Error      interface{} `json:"error"` // TODO - Tidy up in onix? this is normally shown as bool.
	Message    string      `json:"message"`
	Operation  string      `json:"operation"`
	StatusCode int
}

func (o *OnixApiPutDeleteResponse) HasError() bool {
	switch v := o.Error.(type) {
	case bool:
		return o.Error.(bool)
	case string:
		if o.Error.(string) == "" {
			return false
		} else {
			return true
		}
	default:
		fmt.Errorf("Found type %s, thats odd", v)
		// Default shouldn't happen, but if it did, its an error.
		return true
	}
}

type OnixApiGetResponse struct {
	Items []OnixItem
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
	onixResponse.StatusCode = resp.StatusCode
	CheckOnixError(json.NewDecoder(resp.Body).Decode(onixResponse))

	return onixResponse, err
}

func (o *OnixClient) GetItemType(key string) (OnixItemType, error) {
	// TODO - this will be better when we can directly ask for the itemtype/key
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/itemtype", o.BaseURL), nil)
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
				Key:         item.Key,
				Name:        item.Name,
				Description: item.Description,
			}, nil
		}
	}

	return OnixItemType{}, err
}

func (o *OnixClient) GetItem(key string) (OnixItem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/item/%s", o.BaseURL, key), nil)
	CheckOnixError(err)

	resp, err := http.DefaultClient.Do(req)
	CheckOnixError(err)
	defer resp.Body.Close()

	onixResponse := new(OnixItem)
	CheckOnixError(json.NewDecoder(resp.Body).Decode(onixResponse))

	return *onixResponse, nil
}
