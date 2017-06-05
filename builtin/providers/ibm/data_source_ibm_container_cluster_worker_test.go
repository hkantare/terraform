package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMContainerClusterWorkerDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterWorkerDataSourceConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster_worker.testacc_ds_worker", "state", "normal"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterWorkerDataSourceConfig(clusterName string) string {
	return fmt.Sprintf(`
data "ibm_org" "org" {
    org = "%s"
}

data "ibm_space" "space" {
  org    = "%s"
  space  = "%s"
}

data "ibm_account" "acc" {
   org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "testacc_cluster" {
    name = "%s"
    datacenter = "dal10"
    workers = [{
    name = "worker1"
    action = "add"
  },]
	machine_type = "free"
	isolation = "public"
	public_vlan_id = "vlan"
	private_vlan_id = "vlan"

    org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"
}
data "ibm_container_cluster" "testacc_ds_cluster" {
	org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"
    cluster_name_id = "${ibm_container_cluster.testacc_cluster.id}"
}
data "ibm_container_cluster_worker" "testacc_ds_worker" {
	org_guid = "${data.ibm_org.org.id}"
	space_guid = "${data.ibm_space.space.id}"
	account_guid = "${data.ibm_account.acc.id}"
    worker_id = "${data.ibm_container_cluster.testacc_ds_cluster.workers[0]}"
}
`, cfOrganization, cfOrganization, cfSpace, clusterName)
}
