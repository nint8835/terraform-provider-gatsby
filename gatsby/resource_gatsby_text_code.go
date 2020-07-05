package gatsby

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGatsbyTextCode() *schema.Resource {
	return &schema.Resource{
		Create: resourceGatsbyTextCodeCreate,
		Read:   resourceGatsbyTextCodeRead,
		Update: resourceGatsbyTextCodeUpdate,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"multiline": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"code": {
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

func resourceGatsbyTextCodeCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextCodeRead(d, m)
}

func resourceGatsbyTextCodeRead(d *schema.ResourceData, m interface{}) error {
	multiline := d.Get("multiline").(bool)
	language := d.Get("language").(string)
	code := d.Get("code").(string)

	var formattedText string
	if multiline {
		formattedText = fmt.Sprintf("```%s\n%s\n```", language, code)
	} else {
		formattedText = fmt.Sprintf("`%s`", code)
	}

	d.Set("contents", formattedText)
	d.SetId(getID())
	return nil
}

func resourceGatsbyTextCodeUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGatsbyTextCodeRead(d, m)
}
