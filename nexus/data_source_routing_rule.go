/*
Use this data source to work with routing rules

Example Usage

```hcl
data "nexus_routing_rule" "stop_leaks" {
  name = "stop-leaks"
}
```
*/
package nexus

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRoutingRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoutingRuleRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the routing rule",
				Required:    true,
				Type:        schema.TypeString,
			},
			"description": {
				Computed:    true,
				Description: "The description of the routing rule",
				Type:        schema.TypeString,
			},
			"mode": {
				Computed:    true,
				Description: "The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW`",
				Type:        schema.TypeString,
			},
			"matchers": {
				Computed:    true,
				Description: "Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above mode)",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeSet,
			},
		},
	}
}

func dataSourceRoutingRuleRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("name").(string))
	return resourceRoutingRuleRead(d, m)
}
