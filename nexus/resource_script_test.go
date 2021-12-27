package nexus

import (
	"fmt"
	"testing"

	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceScript(t *testing.T) {
	resName := "nexus_script.acceptance"

	script := nexus.Script{
		Name:    acctest.RandString(10),
		Content: "log.info('Hello, World!')",
		Type:    "groovy",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScriptConfig(script),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "id", script.Name),
					resource.TestCheckResourceAttr(resName, "name", script.Name),
					resource.TestCheckResourceAttr(resName, "type", script.Type),
					resource.TestCheckResourceAttr(resName, "content", script.Content),
				),
			},
			{
				ResourceName:      resName,
				ImportState:       true,
				ImportStateId:     script.Name,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceScriptConfig(script nexus.Script) string {
	return fmt.Sprintf(`
resource "nexus_script" "acceptance" {
	name    = "%s"
	content = "%s"
	type    = "%s"
}
`, script.Name, script.Content, script.Type)
}
