package ciao

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

const pemCertReqType = "CERTIFICATE REQUEST"

func resourceCheck() *schema.Resource {
	return &schema.Resource{
		Create: createCheck,
		Update: updateCheck,
		Delete: deleteCheck,
		Read:   readCheck,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique name of the check",
			},

			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Target URL to check",
			},

			"cron": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "*/5 * * *",
				Description: "Cron pattern for checks",
			},

			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "activation status of the check",
			},
		},
	}
}

func apply(chk *Check, d *schema.ResourceData) {
	d.SetId(strconv.Itoa(chk.ID))
	d.Set("name", chk.Name)
	d.Set("url", chk.URL)
	d.Set("cron", chk.Cron)
	d.Set("active", chk.Active)
}

func createCheck(d *schema.ResourceData, meta interface{}) error {
	chk := &Check{
		Name:   d.Get("name").(string),
		URL:    d.Get("url").(string),
		Cron:   d.Get("cron").(string),
		Active: d.Get("active").(bool),
	}
	chk, err := meta.(*Client).NewCheck(chk)
	if err != nil {
		return err
	}
	apply(chk, d)
	return nil
}

func updateCheck(d *schema.ResourceData, meta interface{}) error {
	chk := &Check{
		Name:   d.Get("name").(string),
		URL:    d.Get("url").(string),
		Cron:   d.Get("cron").(string),
		Active: d.Get("active").(bool),
	}
	chk, err := meta.(*Client).UpdateCheck(d.Id(), chk)
	if err != nil {
		return err
	}
	apply(chk, d)
	return nil
}

func deleteCheck(d *schema.ResourceData, meta interface{}) error {
	err := meta.(*Client).DeleteCheck(d.Id())
	d.SetId("")
	return err
}

func readCheck(d *schema.ResourceData, meta interface{}) error {
	// d.GetId()
	chk, err := meta.(*Client).ReadCheck(d.Id())
	if err != nil {
		return err
	}
	apply(chk, d)
	return nil
}
