/*
Use this resource to create a Nexus privilege

Example Usage

```hcl
resource "nexus_privilege" "example" {
  name    = "example-privilige"
  actions = "all"
  type    = "repository-admin"
}
```
*/
package nexus

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePrivilege() *schema.Resource {
	return &schema.Resource{
		Create: resourcePrivilegeCreate,
		Read:   resourcePrivilegeRead,
		Update: resourcePrivilegeUpdate,
		Delete: resourcePrivilegeDelete,
		Exists: resourcePrivilegeExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"actions": {
				Description: "Actions for the privilege (browse, read, edit, add, delete, all and run)",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Type:        schema.TypeSet,
			},
			"content_selector": {
				Description: "The content selector for the privilege",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"description": {
				Description: "A description of the privilege",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"domain": {
				Description: "The domain of the privilege",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"format": {
				Description:  "The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`, `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(nexus.RepositoryFormats, false),
			},
			"name": {
				Description: "The name of the privilege",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"repository": {
				Description: "The repository of the privilege",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"type": {
				Description:  "The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`, `repository-view`, `script`, `wildcard`",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"application", "repository-admin", "repository-content-selector", "repository-view", "script", "wildcard"}, false),
			},
			"pattern": {
				Description: "The wildcard privilege pattern",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"script_name": {
				Description: "The script name related to the privilege",
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func getPrivilegeFromResourceData(d *schema.ResourceData) security.Privilege {
	privilege := security.Privilege{
		Actions: interfaceSliceToStringSlice(d.Get("actions").(*schema.Set).List()),
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		privilege.Description = description.(string)
	}

	if contentSelector, ok := d.GetOk("content_selector"); ok {
		privilege.ContentSelector = contentSelector.(string)
	}

	if domain, ok := d.GetOk("domain"); ok {
		privilege.Domain = domain.(string)
	}

	if format, ok := d.GetOk("format"); ok {
		privilege.Format = format.(string)
	}

	if repository, ok := d.GetOk("repository"); ok {
		privilege.Repository = repository.(string)
	}

	if pattern, ok := d.GetOk("pattern"); ok {
		privilege.Pattern = pattern.(string)
	}

	if scriptName, ok := d.GetOk("script_name"); ok {
		privilege.ScriptName = scriptName.(string)
	}

	return privilege
}

func setPrivilegeToResourceData(privilege *security.Privilege, d *schema.ResourceData) error {
	d.SetId(privilege.Name)
	d.Set("actions", privilege.Actions)
	d.Set("content_selector", privilege.ContentSelector)
	d.Set("description", privilege.Description)
	d.Set("domain", privilege.Domain)
	d.Set("format", privilege.Format)
	d.Set("name", privilege.Name)
	d.Set("repository", privilege.Repository)
	d.Set("type", privilege.Type)
	d.Set("pattern", privilege.Pattern)
	d.Set("script_name", privilege.ScriptName)
	return nil
}

func resourcePrivilegeCreate(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	privilege := getPrivilegeFromResourceData(d)

	if err := service.Security.Privilege.Create(privilege); err != nil {
		return err
	}

	d.SetId(privilege.Name)

	return resourcePrivilegeRead(d, m)
}

func resourcePrivilegeRead(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	privilege, err := service.Security.Privilege.Get(d.Id())
	if err != nil {
		return err
	}

	if privilege == nil {
		d.SetId("")
		return nil
	}

	return setPrivilegeToResourceData(privilege, d)
}

func resourcePrivilegeUpdate(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	privilege := getPrivilegeFromResourceData(d)
	if err := cservice.Security.Privilege.Update(d.Id(), privilege); err != nil {
		return err
	}

	return resourcePrivilegeRead(d, m)
}

func resourcePrivilegeDelete(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	if err := service.Security.Privilege.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourcePrivilegeExists(d *schema.ResourceData, m interface{}) (bool, error) {
	service := m.(nexus.NexusService)

	privilege, err := service.Security.Privilege.Get(d.Id())
	return privilege != nil, err
}
