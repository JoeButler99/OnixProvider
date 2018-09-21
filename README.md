
About
=====
* This is a work in Progress

Terraform provider for the Onix CMDB 

https://github.com/gatblau/onix 

OnixProvider aims to make it possible to both define Onix layouts and add items dynamically from within Terrafrom.


Terraform resources
===================

* See the main.tf example for more information

    resource "onix_itemtype" "AWS VPC" {
      name        = "AWS VPC"
      description = "AWS VPC"
    }


Terraform Links
===============

* https://www.terraform.io/docs/extend/writing-custom-providers.html

