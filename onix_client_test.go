package main

import (
	"testing"
)

//	TODO - This is not a 'proper' test file. I used this to speed up development against the onix client
var test_oc = OnixClient{BaseURL: "http://localhost:8080"}

func CheckTestErrorFail(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf(msg)
	}
}

func ErrorIfTrue(t *testing.T, condition bool, msg string) {
	if condition {
		t.Errorf(msg)
	}
}

func TestOnixClient_ItemType(t *testing.T) {
	oit := OnixItemType{
		Key:         "TESTKEY1",
		Name:        "TestT",
		Description: "Test Item Type 1",
	}
	ErrorIfTrue(t, oit.Id != 0, "New struct ID not 0")

	// Ensure key is not there
	_, err := oc.Delete("itemtype", oit.Key)
	CheckTestErrorFail(t, err, "Error while deleting itemtype. (Which we were not expecting to exist.)")

	// Put in the new OnixItemType oit
	resp, err := test_oc.Put("itemtype", oit.Key, oit.GetJsonBytesReader())
	CheckTestErrorFail(t, err, "Error on Onix PUT (OnixItemType)")

	ErrorIfTrue(t, resp.HasError(), "Api returned Error.")
	ErrorIfTrue(t, !resp.Changed, "API did not report change")
	ErrorIfTrue(t, resp.Operation != "C", "Expected operation type to be 'C'")

	// Get the ItemType
	getResp, err := oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")
	ErrorIfTrue(t, getResp.Id == 0, "Expected ID to be not 0")
	ErrorIfTrue(t, getResp.Key != oit.Key, "Key found to be different on get.")
	ErrorIfTrue(t, getResp.Name != oit.Name, "Name found to be different on get.")
	ErrorIfTrue(t, getResp.Description!= oit.Description, "Desc found to be different on get.")

	// Patch the ItemType
	oit.Description = "New Description"
	oit.Id = getResp.Id
	ErrorIfTrue(t, getResp.Description == oit.Description, "Description not updated before test.")
	resp, err = test_oc.Put("itemtype", oit.Key, oit.GetJsonBytesReader())
	CheckTestErrorFail(t, err, "Error on PUT")
	patchedGetResp, err := oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")

	ErrorIfTrue(t, patchedGetResp.Id == 0, "Expected ID to be not 0")
	ErrorIfTrue(t, patchedGetResp.Key != oit.Key, "Key found to be different on get.")
	ErrorIfTrue(t, patchedGetResp.Name != oit.Name, "Name found to be different on get.")
	ErrorIfTrue(t, patchedGetResp.Description != oit.Description, "Desc found to be different on get.")

	// Delete the ItemType
	deleteResp, err := oc.Delete("itemtype", oit.Key)
	CheckTestErrorFail(t, err, "Error while deleting itemtype.")
	ErrorIfTrue(t, deleteResp.HasError(), "Delete had API error")
	ErrorIfTrue(t, deleteResp.Operation != "D", "Delete had API error")

	// Check the ItemType is gone
	emptyResp, err := oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")
	ErrorIfTrue(t, emptyResp.Id != 0, "ID still found.")

}



//func TestOnixClient_Put(t *testing.T) {
//	oti := OnixItemType{
//		Name:        "TestT",
//		Key:         "TestKey",
//		Description: "Test Item Type 1",
//	}
//
//	resp, err := test_oc.Put("itemtype", oti.Name, oti.GetJsonBytesReader())
//	fmt.Println(resp)
//	if err != nil {
//		t.Errorf("OnixClient Error")
//	}
//
//	if resp.HasError() {
//
//		t.Errorf("Onix API returned Error")
//	}
//
//}
//
//func TestOnixClient_Delete(t *testing.T) {
//	TestOnixClient_Put(t)
//
//	resp, err := test_oc.Delete("itemtype", "TestT")
//	if err != nil {
//		t.Errorf("OnixClient Error")
//	}
//
//	if resp.Error != "" {
//		t.Errorf("Onix API returned Error")
//	}
//
//}
//
//func TestOnixClient_GetItemType(t *testing.T) {
//	TestOnixClient_Put(t)
//
//	resp, err := test_oc.GetItemType("itemtype", "TestT")
//	if err != nil {
//		t.Errorf("OnixClient Error")
//	}
//
//	if resp.Name != "TestT" {
//		t.Errorf("Did not find TestT in itemtype response.")
//	}
//
//	if resp.Id == 0 {
//		t.Errorf("Got ID 0 for saved object")
//	}
//
//	TestOnixClient_Delete(t)
//}