package gatsby

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyPageCreate,
		Read:   resourceGatsbyPageRead,
		Update: resourceGatsbyPageUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"contents": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGatsbyPageCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyPageRead(d, m)
}

func resourceGatsbyPageRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)

	d.SetId(path)
	return nil
}

func resourceGatsbyPageUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyPageRead(d, m)
}
