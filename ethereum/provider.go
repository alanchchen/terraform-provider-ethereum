package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ETHEREUM_ENDPOINT", nil),
				Description: "The endpoint of the Ethereum client to access.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ethereum_local_account": ResourceEthereumLocalAccount(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	// secret := d.Get("secret").(string)

	return ethclient.Dial(endpoint)
}
