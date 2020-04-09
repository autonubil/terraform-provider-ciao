package ciao

import (
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIAO_URL", ""),
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIAO_USER", "TOKEN"),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIAO_PASSWORD", ""),
			},
			"insecure": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,

				DefaultFunc: schema.EnvDefaultFunc("CIAO_INSECURE", true),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ciao_check": resourceCheck(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	ciaoURL := d.Get("url").(string)
	username := d.Get("user").(string)
	password := d.Get("password").(string)

	urlParsed, err := url.Parse(ciaoURL)
	if err != nil {
		return nil, err
	}

	ciaoClient := NewClient(ciaoURL, username, password)

	// Enable debug mode
	ciaoClient.SetDebug(true)

	if urlParsed.Scheme == "https" {
		insecure := d.Get("insecure").(bool)
		ciaoClient.SetInsecure(insecure)
	}
	return ciaoClient, nil
}
