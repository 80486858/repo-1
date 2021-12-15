/*
Use this resource to create a Nexus Repository.

Example Usage

Example Usage - Apt hosted repository

```hcl

resource "nexus_repository_yum_hosted" "yum" {
  name = "yummy"
}

resource "nexus_repository_yum_hosted" "yum1" {
  deploy_policy  = "STRICT"
  name = "yummy1"
  online = true
  repodata_depth = 4

  cleanup {
    policy_names = ["policy"]
  }

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
    write_policy                   = "ALLOW"
  }

}
```
*/
package nexus

import (
	"strings"

	nexus "github.com/datadrivers/go-nexus-client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRepositoryYumHosted() *schema.Resource {
	return &schema.Resource{
		Create: resourceYumRepositoryCreate,
		Read:   resourceYumRepositoryRead,
		Update: resourceYumRepositoryUpdate,
		Delete: resourceYumRepositoryDelete,
		Exists: resourceYumRepositoryExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "A unique identifier for this repository",
				Required:    true,
				Type:        schema.TypeString,
			},
			"online": {
				Default:     true,
				Description: "Whether this repository accepts incoming requests",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"repodata_depth": {
				Default:     0,
				Description: "Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"deploy_policy": {
				Default:     "STRICT",
				Description: "Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"cleanup": {
				Description: "Cleanup policies",
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_names": {
							Description: "List of policy names",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
							Set: func(v interface{}) int {
								return schema.HashString(strings.ToLower(v.(string)))
							},
							Type: schema.TypeSet,
						},
					},
				},
			},
			"storage": {
				Description: "The storage configuration of the repository",
				DefaultFunc: repositoryStorageDefault,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"blob_store_name": {
							Default:     "default",
							Description: "Blob store used to store repository contents",
							Optional:    true,
							Type:        schema.TypeString,
						},
						"strict_content_type_validation": {
							Default:     true,
							Description: "Whether to validate uploaded content's MIME type appropriate for the repository format",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"write_policy": {
							Description: "Controls if deployments of and updates to assets are allowed",
							Default:     "ALLOW",
							Optional:    true,
							Type:        schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{
								"ALLOW",
								"ALLOW_ONCE",
								"DENY",
							}, false),
						},
					},
				},
			},
			/* This is remains here until the go client has been adpated */
			"yum": {
				Description: "Yum specific configuration of the repository",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repodata_depth": {
							Description: "Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
						},
						"deploy_policy": {
							Description:  "Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"STRICT", "PERMISSIVE"}, false),
						},
					},
				},
			},
			"type": {
				Description: "Repository type",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func getYumRepositoryFromResourceData(d *schema.ResourceData) nexus.Repository {
	repo := nexus.Repository{
		Format: "yum",
		Name:   d.Get("name").(string),
		Online: d.Get("online").(bool),
		Type:   "hosted",
	}

	if _, ok := d.GetOk("cleanup"); ok {
		cleanupList := d.Get("cleanup").([]interface{})
		cleanupConfig := cleanupList[0].(map[string]interface{})
		repo.RepositoryCleanup = &nexus.RepositoryCleanup{
			PolicyNames: interfaceSliceToStringSlice(cleanupConfig["policy_names"].(*schema.Set).List()),
		}
	} else {
		repo.RepositoryCleanup = &nexus.RepositoryCleanup{
			PolicyNames: []string{"string"},
		}
	}

	var storageList []interface{}

	if _, ok := d.GetOk("storage"); ok {
		storageList = d.Get("storage").([]interface{})
		storageConfig := storageList[0].(map[string]interface{})

		repo.RepositoryStorage = &nexus.RepositoryStorage{
			BlobStoreName:               storageConfig["blob_store_name"].(string),
			StrictContentTypeValidation: storageConfig["strict_content_type_validation"].(bool),
		}

		writePolicy := storageConfig["write_policy"].(string)
		repo.RepositoryStorage.WritePolicy = &writePolicy
	} else {
		repo.RepositoryStorage = &nexus.RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		}
		writePolicy := "ALLOW"
		repo.RepositoryStorage.WritePolicy = &writePolicy
	}

	repo.RepositoryYum = &nexus.RepositoryYum{
		RepodataDepth: d.Get("repodata_depth").(int),
		DeployPolicy:  d.Get("deploy_policy").(string),
	}

	return repo
}

func setYumRepositoryToResourceData(repo *nexus.Repository, d *schema.ResourceData) error {
	d.SetId(repo.Name)
	d.Set("format", repo.Format)
	d.Set("name", repo.Name)
	d.Set("online", repo.Online)
	d.Set("type", "hosted")

	if repo.RepositoryCleanup != nil {
		if err := d.Set("cleanup", flattenRepositoryCleanup(repo.RepositoryCleanup)); err != nil {
			return err
		}
	}

	if repo.RepositoryYum != nil {
		if err := d.Set("yum", flattenRepositoryYum(repo.RepositoryYum)); err != nil {
			return err
		}
	}

	if err := d.Set("storage", flattenRepositoryStorage(repo.RepositoryStorage, d)); err != nil {
		return err
	}

	return nil
}

func resourceYumRepositoryCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(nexus.Client)

	repo := getYumRepositoryFromResourceData(d)

	if err := client.RepositoryCreate(repo); err != nil {
		return err
	}

	if err := setYumRepositoryToResourceData(&repo, d); err != nil {
		return err
	}

	return resourceRepositoryRead(d, m)
}

func resourceYumRepositoryRead(d *schema.ResourceData, m interface{}) error {
	nexusClient := m.(nexus.Client)

	repo, err := nexusClient.RepositoryRead(d.Id())
	if err != nil {
		return err
	}

	if repo == nil {
		d.SetId("")
		return nil
	}

	return setYumRepositoryToResourceData(repo, d)
}

func resourceYumRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(nexus.Client)

	repoName := d.Id()
	repo := getYumRepositoryFromResourceData(d)

	if err := client.RepositoryUpdate(repoName, repo); err != nil {
		return err
	}

	if err := setYumRepositoryToResourceData(&repo, d); err != nil {
		return err
	}

	return resourceRepositoryRead(d, m)
}

func resourceYumRepositoryDelete(d *schema.ResourceData, m interface{}) error {
	nexusClient := m.(nexus.Client)

	return nexusClient.RepositoryDelete(d.Id())
}

func resourceYumRepositoryExists(d *schema.ResourceData, m interface{}) (bool, error) {
	nexusClient := m.(nexus.Client)

	repo, err := nexusClient.RepositoryRead(d.Id())
	return repo != nil, err
}
