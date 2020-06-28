package gatsby

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextLinkCreate,
		Read:   resourceGatsbyTextLinkRead,
		Update: resourceGatsbyTextLinkUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGatsbyTextLinkCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextLinkRead(d, m)
}

func resourceGatsbyTextLinkRead(d *schema.ResourceData, m interface{}) error {
	url := d.Get("url").(string)
	label, labelExists := d.GetOk("label")
	var labelText string

	if labelExists {
		labelText = label.(string)
	} else {
		labelText = url
	}

	formattedText := fmt.Sprintf("[%s](%s)", labelText, url)
	d.Set("contents", formattedText)
	d.SetId(getID(labelText))
	return nil
}

func resourceGatsbyTextLinkUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextLinkRead(d, m)
}
