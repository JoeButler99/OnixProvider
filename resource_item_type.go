package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
)

// TODO - Look at getting the conn from meta (like the AWS provider)
var oc = OnixClient{ BaseURL: "http://localhost:8080"}

func resourceItemType() *schema.Resource {
	return &schema.Resource{
		Create: resourceItemTypeCreate,
		Read:   resourceItemTypeRead,
		Update: resourceItemTypeUpdate,
		Delete: resourceItemTypeDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceItemTypeCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	d.Set("name", name)
	d.Set("description", description)

	payload, err := json.Marshal(OnixItemType{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}

	_ , err = oc.Put("itemtype", name, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	d.SetId(name)
	return nil
}

func resourceItemTypeRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resp, err := oc.GetItemType("itemtype", name)
	if err != nil { return err}

	if resp.Name != name {
		// If we don't find the item type its not saved
		d.SetId("")
	}
	return nil
}

func resourceItemTypeUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceItemTypeDelete(d *schema.ResourceData, m interface{}) error {
	result , err := oc.Delete("itemtype", d.Get("name").(string))
	if err != nil {
		return err
	} else if result.Error {
		return errors.New("Onix API Error")
	}
	return nil
}
