---
layout: "ibmcloud"
page_title: "IBM Cloud: cs_credentials"
sidebar_current: "docs-ibmcloud-resource-cs-credentials"
description: |-
  Manages IBM Cloud infrastructure credentials.
---

# ibmcloud\_cs_credentials

Set Bluemix Infrastructure account credentials for your Bluemix account. These credentials allow you to access the Bluemix Infrastructure portfolio through your Bluemix account

## Example Usage

In the following example, you can create a cs credentials:

```hcl
resource "ibmcloud_cs_credentials" "credentials" {
    softlayer_username = "test"
	softlayer_api_key = "h4njklkjln4frtg"
    org_guid = "test"
	space_guid = "test_space"
	account_guid = "test_acc"
}
}
```

## Argument Reference

The following arguments are supported:


* `softlayer_username` - (Required) The SoftLayer user name.
* `softlayer_api_key` - (Required)  The SoftLayer api key.
* `org_guid` - (Required) The GUID for the Bluemix organization.
* `space_guid` - (Required) The GUID for the Bluemix space.
* `account_guid` - (Required) The GUID for the Bluemix account.

    
## Attributes Reference

The following attributes are exported:

* `id` - ID of the credentials.