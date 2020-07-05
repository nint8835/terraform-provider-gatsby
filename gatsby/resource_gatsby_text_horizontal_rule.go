package gatsby

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextHorizontalRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextHorizontalRuleCreate,
		Read:   resourceGatsbyTextHorizontalRuleRead,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGatsbyTextHorizontalRuleCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextHorizontalRuleRead(d, m)
}

func resourceGatsbyTextHorizontalRuleRead(d *schema.ResourceData, m interface{}) error {
	formattedText := "---"
	d.Set("contents", formattedText)
	d.SetId(getID())
	return nil
}
