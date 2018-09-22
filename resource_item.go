package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TODO - Look at getting the conn from meta (like the AWS provider)
//var oc = OnixClient{ BaseURL: "http://localhost:8080"}

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
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceItemCreate(d *schema.ResourceData, m interface{}) error {
	//key := d.Get("key").(string)
	//name := d.Get("name").(string)
	//description := d.Get("description").(string)
	//itemtype, err := strconv.Atoi(d.Get("itemtype").(string))
	//if err != nil {
	//	return err
	//}
	//
	//oi := OnixItem{
	//	Key:         key,
	//	Name:        name,
	//	Description: description,
	//	Type:        itemtype,
	//}
	//_, err = oc.Put("item", name, oi.GetJsonBytesReader())
	//if err != nil {
	//	return err
	//}

	//d.SetId(name)
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
	//result , err := oc.Delete("itemtype", d.Get("name").(string))
	//if err != nil {
	//	return err
	//} else if result.Error {
	//	return errors.New("Onix API Error")
	//}
	return nil
}
