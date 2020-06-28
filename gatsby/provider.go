package gatsby

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"gatsby_text_bold": resourceGatsbyTextBold(),
		},
	}
}
