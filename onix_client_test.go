package main

import (
	"runtime/debug"
	"testing"
)

//	TODO - This is not a 'proper' test file. I used this to speed up development against the onix client
var test_oc = OnixClient{BaseURL: "http://localhost:8080"}

func CheckTestErrorFail(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf(msg)
		debug.PrintStack()
	}
}

func ErrorIfTrue(t *testing.T, condition bool, msg string) {
	if condition {
		t.Errorf(msg)
		debug.PrintStack()
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
	_, err := test_oc.Delete("itemtype", oit.Key)
	CheckTestErrorFail(t, err, "Error while deleting itemtype. (Which we were not expecting to exist.)")

	// Put in the new OnixItemType oit
	resp, err := test_oc.Put("itemtype", oit.Key, oit.GetJsonBytesReader())
	CheckTestErrorFail(t, err, "Error on Onix PUT (OnixItemType)")

	ErrorIfTrue(t, resp.HasError(), "Api returned Error.")
	ErrorIfTrue(t, !resp.Changed, "API did not report change")
	ErrorIfTrue(t, resp.Operation != "C", "Expected operation type to be 'C'")

	// Get the ItemType
	getResp, err := test_oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")
	ErrorIfTrue(t, getResp.Id == 0, "Expected ID to be not 0")
	ErrorIfTrue(t, getResp.Key != oit.Key, "Key found to be different on get.")
	ErrorIfTrue(t, getResp.Name != oit.Name, "Name found to be different on get.")
	ErrorIfTrue(t, getResp.Description != oit.Description, "Desc found to be different on get.")

	// Patch the ItemType
	oit.Description = "New Description"
	oit.Id = getResp.Id
	ErrorIfTrue(t, getResp.Description == oit.Description, "Description not updated before test.")
	resp, err = test_oc.Put("itemtype", oit.Key, oit.GetJsonBytesReader())
	CheckTestErrorFail(t, err, "Error on PUT")
	patchedGetResp, err := test_oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")

	ErrorIfTrue(t, patchedGetResp.Id == 0, "Expected ID to be not 0")
	ErrorIfTrue(t, patchedGetResp.Key != oit.Key, "Key found to be different on get.")
	ErrorIfTrue(t, patchedGetResp.Name != oit.Name, "Name found to be different on get.")
	ErrorIfTrue(t, patchedGetResp.Description != oit.Description, "Desc found to be different on get.")

	// Delete the ItemType
	deleteResp, err := test_oc.Delete("itemtype", oit.Key)
	CheckTestErrorFail(t, err, "Error while deleting itemtype.")
	ErrorIfTrue(t, deleteResp.HasError(), "Delete had API error")
	ErrorIfTrue(t, deleteResp.Operation != "D", "Delete had API error")

	// Check the ItemType is gone
	emptyResp, err := test_oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")
	ErrorIfTrue(t, emptyResp.Id != 0, "ID still found.")

}

func TestOnixClient_Item(t *testing.T) {
	oi := OnixItem{
		Key:         "TestOnixItem",
		Name:        "TestOnixItem",
		Description: "Test Onix Item",
		Type:        "",
		Status:      1,
	}
	oit := OnixItemType{
		Key:         "ParentTestType",
		Name:        "ParentTestType",
		Description: "Test Onix ItemType used to test Onix Item",
	}

	// Ensure item is not there. (We need to check this before ItemType)
	delresp, err := test_oc.Delete("item", oi.Key)
	CheckTestErrorFail(t, err, "Error while deleting item. (Which we were not expecting to exist.)")
	ErrorIfTrue(t, delresp.HasError(), "Error while deleting item. (Which we were not expecting to exist.)")

	// Ensure key is not there
	delresp, err = test_oc.Delete("itemtype", oit.Key)
	CheckTestErrorFail(t, err, "Error while deleting itemtype. (Which we were not expecting to exist.)")
	ErrorIfTrue(t, delresp.HasError(), "Error while deleting itemtype. (Which we were not expecting to exist.)")

	// Put in the new OnixItemType oit
	resp, err := test_oc.Put("itemtype", oit.Key, oit.GetJsonBytesReader())
	CheckTestErrorFail(t, err, "Error on Onix PUT (OnixItemType)")
	ErrorIfTrue(t, resp.HasError(), "Api returned Error.")
	ErrorIfTrue(t, !resp.Changed, "API did not report change")
	ErrorIfTrue(t, resp.Operation != "C", "Expected operation type to be 'C'")

	// Get the ItemType we just saved
	getResp, err := test_oc.GetItemType(oit.Key)
	CheckTestErrorFail(t, err, "Error on GET")
	ErrorIfTrue(t, getResp.Id == 0, "Expected ID to be non 0")
	ErrorIfTrue(t, getResp.Key != oit.Key, "Key found to be different on get.")
	ErrorIfTrue(t, getResp.Name != oit.Name, "Name found to be different on get.")
	ErrorIfTrue(t, getResp.Description != oit.Description, "Desc found to be different on get.")

	// Now Lets create an Item to put in
	oi.Type = getResp.Key // Put the right key on it
	resp, err = test_oc.Put("item", oi.Key, oi.GetJsonBytesReader())
	CheckTestErrorFail(t, err, "Error on PUT item")
	ErrorIfTrue(t, resp.HasError(), "Api returned Error.")
	ErrorIfTrue(t, !resp.Changed, "API did not report change")
	ErrorIfTrue(t, resp.Operation != "C", "Expected operation type to be 'C'")

	getItemResp, err := test_oc.GetItem(oi.Key)
	CheckTestErrorFail(t, err, "Error getting new Item")
	//ErrorIfTrue(t, getItemResp.Id == 0, "Expected ID to be non 0")   // Id not returned by item
	ErrorIfTrue(t, getItemResp.Key != oi.Key, "Different Key returned from GET")

	// TODO - Delete item and itemtype and check they are both gone.

}
