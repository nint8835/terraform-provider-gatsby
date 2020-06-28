package gatsby

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextHeading() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextHeadingCreate,
		Read:   resourceGatsbyTextHeadingRead,
		Update: resourceGatsbyTextHeadingUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"text": {
				Type:     schema.TypeString,
				Required: true,
			},
			"level": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 1 {
						errs = append(errs, fmt.Errorf("%q must be greater than or equal to 1, got: %d", key, v))
					}
					return
				},
			},
			"contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGatsbyTextHeadingCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextHeadingRead(d, m)
}

func resourceGatsbyTextHeadingRead(d *schema.ResourceData, m interface{}) error {
	text := d.Get("text").(string)
	level := d.Get("level").(int)
	formattedText := fmt.Sprintf("%s %s", strings.Repeat("#", level), text)
	d.Set("contents", formattedText)
	d.SetId(formattedText)
	return nil
}

func resourceGatsbyTextHeadingUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextHeadingRead(d, m)
}
