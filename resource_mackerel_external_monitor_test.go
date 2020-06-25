package mackerel

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mackerelio/mackerel-client-go"
)

func TestAccMackerelExternalMonitor_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMackerelExternalMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMackerelExternalMonitorConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "name", "terraform_for_mackerel_test_foobar"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "url", "https://terraform.io/"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "service", "Web"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "notification_interval", "10"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_duration", "5"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_warning", "500"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_critical", "1000"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "contains_string", "terraform"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "max_check_attempts", "2"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "certification_expiration_warning", "30"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "certification_expiration_critical", "10"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "skip_certificate_verification", "false"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "method", "GET"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "memo", "XXX"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "request_body", "{\"request\": \"body\"}"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "headers.API-Key", "xxxxxx"),
				),
			},
		},
	})
}

func TestAccMackerelExternalMonitor_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMackerelExternalMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMackerelExternalMonitorConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "name", "terraform_for_mackerel_test_foobar"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "url", "https://terraform.io/"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "service", "Web"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "notification_interval", "10"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_duration", "5"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_warning", "500"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_critical", "1000"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "contains_string", "terraform"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "max_check_attempts", "2"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "certification_expiration_warning", "30"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "certification_expiration_critical", "10"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "skip_certificate_verification", "false"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "method", "GET"),
				),
			},
			{
				Config: testAccCheckMackerelExternalMonitorConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "name", "terraform_for_mackerel_test_foobar_upd"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "url", "https://terraform.io/"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "service", "Web"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "notification_interval", "10"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_duration", "10"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_warning", "800"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "response_time_critical", "900"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "contains_string", "terraform"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "max_check_attempts", "3"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "certification_expiration_warning", "60"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "certification_expiration_critical", "30"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "skip_certificate_verification", "true"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "method", "POST"),
					resource.TestCheckNoResourceAttr(
						"mackerel_external_monitor.foobar", "headers.API-Key"),
				),
			},
		},
	})
}

func TestAccMackerelExternalMonitor_Minimum(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMackerelExternalMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMackerelExternalMonitorConfig_minimum,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "name", "terraform_for_mackerel_test_foobar_min"),
					resource.TestCheckResourceAttr(
						"mackerel_external_monitor.foobar", "url", "https://terraform.io/"),
				),
			},
		},
	})
}

func testAccCheckMackerelExternalMonitorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*mackerel.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mackerel_external_monitor" {
			continue
		}

		monitors, err := client.FindMonitors()
		if err != nil {
			return err
		}
		for _, monitor := range monitors {
			if monitor.MonitorID() == rs.Primary.ID {
				return fmt.Errorf("Monitor still exists")
			}
		}
	}

	return nil
}

const testAccCheckMackerelExternalMonitorConfig_basic = `
resource "mackerel_service" "web" {
  name = "Web"
}

resource "mackerel_external_monitor" "foobar" {
  name                   = "terraform_for_mackerel_test_foobar"
  url                    = "https://terraform.io/"
  service                = mackerel_service.web.name
  notification_interval  = 10
  response_time_duration = 5
  response_time_warning  = 500
  response_time_critical = 1000
  contains_string        = "terraform"
  max_check_attempts     = 2

  certification_expiration_warning  = 30
  certification_expiration_critical = 10

  skip_certificate_verification = false

  request_body = "{\"request\": \"body\"}"
  headers = {
      "Content-Type" = "application/json",
      "API-Key" = "xxxxxx",
  }

  memo = "XXX"
}`

const testAccCheckMackerelExternalMonitorConfig_update = `
resource "mackerel_service" "web" {
  name = "Web"
}

resource "mackerel_external_monitor" "foobar" {
  name                   = "terraform_for_mackerel_test_foobar_upd"
  url                    = "https://terraform.io/"
  method                 = "POST"
  service                = mackerel_service.web.name
  notification_interval  = 10
  response_time_duration = 10
  response_time_warning  = 800
  response_time_critical = 900
  contains_string        = "terraform"
  max_check_attempts     = 3

  certification_expiration_warning  = 60
  certification_expiration_critical = 30

  skip_certificate_verification = true
}`

const testAccCheckMackerelExternalMonitorConfig_minimum = `
resource "mackerel_external_monitor" "foobar" {
  name                   = "terraform_for_mackerel_test_foobar_min"
  url                    = "https://terraform.io/"
}`
