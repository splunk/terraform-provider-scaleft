package scaleft

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"scaleft_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALEFT_KEY", nil),
				Description: "ScaleFT API key.",
			},

			"scaleft_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALEFT_KEY_SECRET", nil),
				Description: "ScaleFT API secret.",
			},

			"scaleft_team": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALEFT_TEAM", nil),
				Description: "ScaleFT Team.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scaleft_project":          resourceScaleftProject(),
			"scaleft_enrollment_token": resourceScaleftToken(),
			"scaleft_assign_group":     resourceScaleftAssignGroup(),
			"scaleft_create_group":     resourceScaleftCreateGroup(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		key:    d.Get("scaleft_key").(string),
		secret: d.Get("scaleft_secret").(string),
		team:   d.Get("scaleft_team").(string),
	}

	return config.Authorization()

}
