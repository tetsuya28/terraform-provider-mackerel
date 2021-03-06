package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mackerelio/mackerel-client-go"
)

func TestAccMackerelDashboard_Basic(t *testing.T) {
	rName := acctest.RandomWithPrefix("TerraformTestDashboard-")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMackerelDashboardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMackerelDashboardConfigBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "title", rName),
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "url_path", fmt.Sprintf("%s/foo/bar", rName)),
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "body_markdown", "# Head1\n## Head2\n\n* List1\n* List2\n"),
				),
			},
		},
	})
}

func TestAccMackerelDashboard_Update(t *testing.T) {
	rName := acctest.RandomWithPrefix("TerraformTestDashboard-")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMackerelDashboardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMackerelDashboardConfigBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "title", rName),
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "url_path", fmt.Sprintf("%s/foo/bar", rName)),
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "body_markdown", "# Head1\n## Head2\n\n* List1\n* List2\n"),
				),
			},
			{
				Config: testAccCheckMackerelDashboardConfigUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "title", rName),
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "url_path", fmt.Sprintf("%s/bar/baz", rName)),
					resource.TestCheckResourceAttr(
						"mackerel_dashboard.foobar", "body_markdown", "# Head1\n## Head2\n\n[Link](https://terraform.io/)\n"),
				),
			},
		},
	})
}

func testAccCheckMackerelDashboardDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*mackerel.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mackerel_dashboard" {
			continue
		}

		_, err := client.FindDashboard(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Dashboard still exists")
		}
	}

	return nil
}

func testAccCheckMackerelDashboardConfigBasic(rName string) string {
	return fmt.Sprintf(`
resource "mackerel_dashboard" "foobar" {
  title         = "%s"
  url_path      = "%s/foo/bar"
	body_markdown = <<EOF
# Head1
## Head2

* List1
* List2
EOF
}`, rName, rName)
}

func testAccCheckMackerelDashboardConfigUpdate(rName string) string {
	return fmt.Sprintf(`
resource "mackerel_dashboard" "foobar" {
  title         = "%s"
  url_path      = "%s/bar/baz"
	body_markdown = <<EOF
# Head1
## Head2

[Link](https://terraform.io/)
EOF
}`, rName, rName)
}
