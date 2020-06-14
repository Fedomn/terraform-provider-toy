package toy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"toy_example_server": resourceServer(),
		},
	}
	return provider
}
