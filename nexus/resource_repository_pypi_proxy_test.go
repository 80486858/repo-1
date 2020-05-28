package nexus

import (
	"fmt"
	"testing"

	nexus "github.com/datadrivers/go-nexus-client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccRepositoryPyPiProxy(t *testing.T) {
	repoName := fmt.Sprintf("test-repo-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{

		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTfStmtForResourcePypiProxy(repoName),
				Check: resource.ComposeTestCheckFunc(
					// Base and common repo props
					// Identity fields
					resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "id", repoName),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "name", repoName),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "format", nexus.RepositoryFormatPyPi),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "type", nexus.RepositoryTypeProxy),
					),
					// Common fields
					resource.ComposeAggregateTestCheckFunc(
						// Online
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "online", "true"),
						// Storage
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "storage.#", "1"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "storage.0.blob_store_name", "default"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "storage.0.strict_content_type_validation", "true"),
						// FIXME: (BUG) Write policy can not be set to ALLOW is not set
						// resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "storage.0.write_policy", "ALLOW"),
					),
					// No fields related to other repo types
					// Format
					resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "maven.#", "0"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "apt.#", "0"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "apt_signing.#", "0"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "bower.#", "0"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "docker.#", "0"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "docker_proxy.#", "0"),
					),
					// Type
					resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "group.#", "0"),
					),
					// Fields related to this format and type
					// Format
					// - No special fields
					// Type
					resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "proxy.#", "1"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "negative_cache.#", "1"),
						resource.TestCheckResourceAttr("nexus_repository.pypi_proxy", "http_client.#", "1"),
					),
				),
			},
			{
				ResourceName:      "nexus_repository.pypi_proxy",
				ImportStateId:     repoName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createTfStmtForResourcePypiProxy(name string) string {
	return fmt.Sprintf(`
resource "nexus_repository" "pypi_proxy" {
	format = "%s"
	name   = "%s"
	online = true
	type   = "%s"

	proxy {
		remote_url  = "https://pypi.org"
	}

	http_client {
	}

	negative_cache {
		enabled = true
		ttl     = 1440
	}

	storage {

	}
}`, nexus.RepositoryFormatPyPi, name, nexus.RepositoryTypeProxy)
}
