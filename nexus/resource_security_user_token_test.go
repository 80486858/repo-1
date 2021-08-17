package nexus

import (
	"fmt"
	"strconv"
	"testing"

	nexus "github.com/datadrivers/go-nexus-client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceSecurityUserToken(t *testing.T) {
	resName := "nexus_security_user_token.acceptance"

	token := nexus.UserTokenConfiguration{
		Enabled:        true,
		ProtectContent: false,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSecurityUserTokenConfig(token),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "enabled", strconv.FormatBool(token.Enabled)),
					resource.TestCheckResourceAttr(resName, "protect_content", strconv.FormatBool(token.ProtectContent)),
				),
			},
		},
	})
}

func testAccResourceSecurityUserTokenConfig(token nexus.UserTokenConfiguration) string {
	return fmt.Sprintf(`
resource "nexus_security_user_token" "acceptance" {
	enabled         = %t
	protect_content = %t
}
`, token.Enabled, token.ProtectContent)
}
