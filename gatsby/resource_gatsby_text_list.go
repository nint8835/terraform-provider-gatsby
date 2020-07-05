package gatsby

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextList() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextListCreate,
		Read:   resourceGatsbyTextListRead,
		Update: resourceGatsbyTextListUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "- ",
			},
			"separator": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "\n",
			},
			"items": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGatsbyTextListCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextListRead(d, m)
}

func resourceGatsbyTextListRead(d *schema.ResourceData, m interface{}) error {
	prefix := d.Get("prefix").(string)
	separator := d.Get("separator").(string)

	items := d.Get("items").([]interface{})

	prefixedStrings := make([]string, len(items))

	for index, value := range items {
		if value == nil {
			value = ""
		}
		prefixedStrings[index] = fmt.Sprintf("%s%s", prefix, value)
	}

	formattedText := strings.Join(prefixedStrings, separator)
	d.Set("contents", formattedText)
	d.SetId(getID())
	return nil
}

func resourceGatsbyTextListUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextListRead(d, m)
}
