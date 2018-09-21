package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

// TODO - Look at getting the conn from meta (like the AWS provider)
var oc = OnixClient{BaseURL: "http://localhost:8080"}

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
				ForceNew: true,
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
	payload, err := json.Marshal(OnixItemType{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}

	_, err = oc.Put("itemtype", name, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	item, err := oc.GetItemType("itemtype", name)
	if err != nil {
		return err
	}

	if item.Id == 0 {
		return errors.New("ID of 0 found for saved object.")
	}

	d.SetId(strconv.Itoa(item.Id))

	return nil
}

func resourceItemTypeRead(d *schema.ResourceData, m interface{}) error {
	id := d.Get("id")

	name := d.Get("name").(string)
	resp, err := oc.GetItemType("itemtype", name)
	if err != nil {
		return err
	}

	// ID check
	if id != nil {
		idInt, err := strconv.Atoi(id.(string))
		if err != nil {
			return err
		}
		if idInt != resp.Id {
			d.SetId("")
		}
	}

	// Name check
	if resp.Name != name {
		d.SetId("")
	}

	return nil
}

func resourceItemTypeUpdate(d *schema.ResourceData, m interface{}) error {
	// As its a put, we just wash through create again.
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	payload, err := json.Marshal(OnixItemType{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}

	_, err = oc.Put("itemtype", name, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	return nil
}

func resourceItemTypeDelete(d *schema.ResourceData, m interface{}) error {
	result, err := oc.Delete("itemtype", d.Get("name").(string))
	if err != nil {
		return err
	} else if result.Error {
		return errors.New("Onix API Error")
	}
	return nil
}
