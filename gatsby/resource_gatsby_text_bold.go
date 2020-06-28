package gatsby

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextBold() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextBoldCreate,
		Read:   resourceGatsbyTextBoldRead,
		Update: resourceGatsbyTextBoldUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"text": {
				Type:     schema.TypeString,
				Required: true,
			},
			"contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

var resourceGatsbyTextBoldRead = textWrapper("**%s**")

func resourceGatsbyTextBoldCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextBoldRead(d, m)
}

func resourceGatsbyTextBoldUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextBoldRead(d, m)
}
