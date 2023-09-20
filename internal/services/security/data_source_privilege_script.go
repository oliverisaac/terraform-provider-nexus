package security

import (
	"github.com/datadrivers/terraform-provider-nexus/internal/schema/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceSecurityPrivilegeScript() *schema.Resource {
	return &schema.Resource{
		Description: `Use this data source to get a privilege for a script`,

		Read: dataSourceSecurityPrivilegeScriptRead,
		Schema: map[string]*schema.Schema{
			"id": common.DataSourceID,
			"name": {
				Description: "Name of the script privilege",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the script privilege",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"readonly": {
				Description: "Whether it is readonly or not",
				Computed:    true,
				Type:        schema.TypeBool,
			},
			"script_name": {
				Description: "The script the privilege applies to",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"actions": {
				Description: "A list of allowed actions. For a list of applicable values see https://help.sonatype.com/repomanager3/nexus-repository-administration/access-control/privileges#Privileges-PrivilegeTypes",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"ADD", "READ", "DELETE", "RUN", "BROWSE", "EDIT"}, false),
				},
			},
		},
	}
}

func dataSourceSecurityPrivilegeScriptRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("name").(string))
	return resourceSecurityPrivilegeScriptRead(d, m)
}
