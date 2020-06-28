package gatsby

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextItalic() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextItalicCreate,
		Read:   resourceGatsbyTextItalicRead,
		Update: resourceGatsbyTextItalicUpdate,
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

func resourceGatsbyTextItalicCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextItalicRead(d, m)
}

func resourceGatsbyTextItalicRead(d *schema.ResourceData, m interface{}) error {
	text := d.Get("text").(string)
	formattedText := fmt.Sprintf("_%s_", text)
	d.Set("contents", formattedText)
	d.SetId(formattedText)
	return nil
}

func resourceGatsbyTextItalicUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextItalicRead(d, m)
}
