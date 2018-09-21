package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TODO - Provider inputs (the client URL)
// TODO -

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"onix_itemtype": resourceItemType(),
			"onix_item":     resourceItem(),
			// TODO onix_link
		},
	}
}
