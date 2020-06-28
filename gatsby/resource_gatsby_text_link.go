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
		Delete: resourceGatsbyTextLinkDelete,

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
	url := d.Get("url")
	label, labelExists := d.GetOk("label")
	if labelExists {
		label = label.(string)
	} else {
		label = url
	}

	formattedText := fmt.Sprintf("[%s](%s)", label, url)
	d.Set("contents", formattedText)
	d.SetId(formattedText)
	return nil
}

func resourceGatsbyTextLinkUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextLinkRead(d, m)
}

func resourceGatsbyTextLinkDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
