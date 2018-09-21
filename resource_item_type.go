package main

import (
	"errors"
	"fmt"
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
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
	key := d.Get("key").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	oit := OnixItemType{
		Key:         key,
		Name:        name,
		Description: description,
	}

	_, err := oc.Put("itemtype", key, oit.GetJsonBytesReader())
	if err != nil {
		return err
	}

	item, err := oc.GetItemType(key)
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
	key := d.Get("key").(string)

	resp, err := oc.GetItemType(key)
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

	d.Set("key", resp.Key)
	d.Set("name", resp.Name)
	d.Set("description", resp.Description)

	return nil
}

func resourceItemTypeUpdate(d *schema.ResourceData, m interface{}) error {
	// As its a put, we just wash through create again.
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	oit := OnixItemType{
		Name:        name,
		Description: description,
	}

	_, err := oc.Put("itemtype", name, oit.GetJsonBytesReader())
	if err != nil {
		return err
	}

	return nil
}

func resourceItemTypeDelete(d *schema.ResourceData, m interface{}) error {
	result, err := oc.Delete("itemtype", d.Get("name").(string))
	if err != nil {
		return err
	} else if result.Error != "" {
		return errors.New(fmt.Sprintf("Onix API flagged error, %s", result.Message))
	}
	return nil
}
