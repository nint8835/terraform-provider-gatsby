package gatsby

import (
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

var resourceGatsbyTextItalicRead = textWrapper("_%s_")

func resourceGatsbyTextItalicCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextItalicRead(d, m)
}

func resourceGatsbyTextItalicUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextItalicRead(d, m)
}
