package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceItem() *schema.Resource {
	return &schema.Resource{
		Create: resourceItemCreate,
		Read:   resourceItemRead,
		Update: resourceItemUpdate,
		Delete: resourceItemDelete,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"itemtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"meta": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

func resourceItemCreate(d *schema.ResourceData, m interface{}) error {
	key := d.Get("key").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	itemtype := d.Get("itemtype").(string)
	meta := d.Get("meta").(string)

	oi := OnixItem{
		Key:         key,
		Name:        name,
		Description: description,
		Type:        itemtype,
		Meta:        meta,
	}

	_, err := oc.Put("item", key, oi.GetJsonBytesReader())
	if err != nil {
		return err
	}

	item, err := oc.GetItem(key)
	if err != nil {
		return err
	}

	d.SetId(item.Key)

	return nil
}

func resourceItemRead(d *schema.ResourceData, m interface{}) error {
	//name := d.Get("name").(string)
	//resp, err := oc.GetItem("item", name)
	//if err != nil { return err}
	//
	//if resp.Name != name {
	//	// If we don't find the item type its not saved
	//	d.SetId("")
	//}
	return nil
}

func resourceItemUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceItemDelete(d *schema.ResourceData, m interface{}) error {
	result, err := oc.Delete("item", d.Get("key").(string))
	if err != nil {
		return err
	} else if result.Error != "" {
		return errors.New(fmt.Sprintf("Onix API flagged error, %s", result.Message))
	}
	return nil
}
