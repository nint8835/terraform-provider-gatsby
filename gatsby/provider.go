package gatsby

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"gatsby_text_bold":            resourceGatsbyTextBold(),
			"gatsby_text_italic":          resourceGatsbyTextItalic(),
			"gatsby_text_link":            resourceGatsbyTextLink(),
			"gatsby_text_heading":         resourceGatsbyTextHeading(),
			"gatsby_text_list":            resourceGatsbyTextList(),
			"gatsby_text_horizontal_rule": resourceGatsbyTextHorizontalRule(),
			"gatsby_text_image":           resourceGatsbyTextImage(),
		},
	}
}
