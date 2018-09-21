package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TODO - Provider inputs

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"onix_itemtype": resourceItemType(),
			// TODO onix_item
			// TODO onix_link
		},
	}
}
