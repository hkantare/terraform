package ibmcloud

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMCloudCSCredentials() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCloudCSCredentialsCreate,
		Read:     resourceIBMCloudCSCredentialsRead,
		Delete:   resourceIBMCloudCSCredentialsDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"softlayer_username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The softlayer user name",
			},
			"softlayer_api_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The softlayer API key",
			},
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceIBMCloudCSCredentialsCreate(d *schema.ResourceData, meta interface{}) error {
	clusterClient := meta.(ClientSession).ClusterClient()
	softlayerUserName := d.Get("softlayer_username").(string)
	softlayerAPIKey := d.Get("softlayer_api_key").(string)

	targetEnv := getClusterTargetHeader(d)

	err := clusterClient.SetCredentials(softlayerUserName, softlayerAPIKey, targetEnv)
	if err != nil {
		return err
	}

	return resourceIBMCloudCSCredentialsRead(d, meta)
}

func resourceIBMCloudCSCredentialsRead(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("softlayer_username").(string))

	return nil
}

func resourceIBMCloudCSCredentialsDelete(d *schema.ResourceData, meta interface{}) error {
	clusterClient := meta.(ClientSession).ClusterClient()

	targetEnv := getClusterTargetHeader(d)

	err := clusterClient.UnsetCredentials(targetEnv)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
