package gatsby

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextImageCreate,
		Read:   resourceGatsbyTextImageRead,
		Update: resourceGatsbyTextImageUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alt_text": {
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

func resourceGatsbyTextImageCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextImageRead(d, m)
}

func resourceGatsbyTextImageRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	altText := d.Get("alt_text").(string)

	formattedText := fmt.Sprintf("![%s](%s)", altText, path)
	d.Set("contents", formattedText)
	d.SetId(getID())
	return nil
}

func resourceGatsbyTextImageUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextImageRead(d, m)
}
