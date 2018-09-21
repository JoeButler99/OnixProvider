package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)


//	TODO - This is not a 'proper' test file. I used this to speed up development against the onix client

func TestOnixClient_Put(t *testing.T) {
	oc := OnixClient{BaseURL: "http://localhost:8080"}

	oti := OnixItemType{
		Name:        "TestT",
		Description: "Test Item Type 1",
	}

	payload, err := json.Marshal(oti)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := oc.Put("itemtype", oti.Name, bytes.NewReader(payload))
	if err != nil {
		t.Errorf("OnixClient Error")
	}

	if resp.Error {
		t.Errorf("Onix API returned Error")
	}

}

func TestOnixClient_Delete(t *testing.T) {
	TestOnixClient_Put(t)

	oc := OnixClient{BaseURL: "http://localhost:8080"}

	resp, err := oc.Delete("itemtype", "TestT")
	if err != nil {
		t.Errorf("OnixClient Error")
	}

	if resp.Error {
		t.Errorf("Onix API returned Error")
	}

}

func TestOnixClient_GetItemType(t *testing.T) {
	TestOnixClient_Put(t)

	oc := OnixClient{BaseURL: "http://localhost:8080"}

	resp, err := oc.GetItemType("itemtype", "TestT")
	if err != nil {
		t.Errorf("OnixClient Error")
	}

	if resp.Name != "TestT" {
		t.Errorf("Did not find TestT in itemtype response.")
	}

	TestOnixClient_Delete(t)
}