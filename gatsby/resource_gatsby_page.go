package gatsby

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/yaml.v2"
)

func resourceGatsbyPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyPageCreate,
		Read:   resourceGatsbyPageRead,
		Update: resourceGatsbyPageUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"contents": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frontmatter": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Default:  make(map[string]interface{}),
			},
			"page_text": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGatsbyPageCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyPageRead(d, m)
}

func resourceGatsbyPageRead(d *schema.ResourceData, m interface{}) error {
	contents := d.Get("contents").(string)
	frontmatter := d.Get("frontmatter")

	data, err := yaml.Marshal(frontmatter)
	if err != nil {
		return err
	}
	formattedText := fmt.Sprintf("---\n%s---\n%s", data, contents)
	d.SetId(getID(formattedText))
	d.Set("page_text", formattedText)
	return nil
}

func resourceGatsbyPageUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyPageRead(d, m)
}
