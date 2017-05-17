package ibmcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMCloudCSCredentials_Basic(t *testing.T) {
	slname := "1204953_hkantare@in.ibm.com"
	slapikey := "b43c2ade072fb295d1adda8f8cf559d64e36e75db50a4799e03cc757c6e3f791"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudCSCredentials_basic(slname, slapikey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibmcloud_cs_credentials.credentials", "softlayer_username", slname),
					resource.TestCheckResourceAttr("ibmcloud_cs_credentials.credentials", "softlayer_api_key", slapikey),
				),
			},
		},
	})
}

func testAccCheckIBMCloudCSCredentials_basic(slname, slapikey string) string {
	return fmt.Sprintf(`
		data "ibmcloud_cf_org" "org" {
			org = "%s"
		}

		data "ibmcloud_cf_space" "space" {
			org    = "%s"
			space  = "%s"
		}

		data "ibmcloud_cf_account" "acc" {
			org_guid = "${data.ibmcloud_cf_org.org.id}"
		}
		
		resource "ibmcloud_cs_credentials" "credentials" {
			softlayer_username		= "%s"
			softlayer_api_key       = "%s"
			space_guid        		= "${data.ibmcloud_cf_space.space.id}"
			org_guid                = "${data.ibmcloud_cf_org.org.id}"
			account_guid            = "${data.ibmcloud_cf_account.acc.id}"
			
		}
	`, cfOrganization, cfOrganization, cfSpace, slname, slapikey)
}
