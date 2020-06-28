package gatsby

import (
	"fmt"

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

func resourceGatsbyTextBoldCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextBoldRead(d, m)
}

func resourceGatsbyTextBoldRead(d *schema.ResourceData, m interface{}) error {
	text := d.Get("text").(string)
	formattedText := fmt.Sprintf("**%s**", text)
	d.Set("contents", formattedText)
	d.SetId(formattedText)
	return nil
}

func resourceGatsbyTextBoldUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextBoldRead(d, m)
}
