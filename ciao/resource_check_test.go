package ciao

import (
	"testing"

	resource "github.com/hashicorp/terraform/helper/resource"
)

func TestCheck(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
                    resource "ciao_check" "autonubil" {
						name ="test"
						cron ="*/5 * * * *"
						url = "https://www.autonubil.de"
						active = false
                    }
                    output "check" {
                        value = ciao_check.autonubil
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ciao_check.autonubil", "name", "test"),
				),
			},
			{
				Config: `
                    resource "ciao_check" "autonubil" {
						name ="autonubil"
						cron ="*/4 * * * *"
						url = "https://www.autonubil.de"
						active = true
                    }
                    output "check" {
                        value = ciao_check.autonubil
                    }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ciao_check.autonubil", "name", "autonubil"),
				),
			},
			{
				Config: ` # remove it all
				`,
				Check: resource.ComposeTestCheckFunc(),
			},
		},
	})
}
