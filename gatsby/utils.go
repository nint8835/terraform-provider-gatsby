package gatsby

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func textWrapper(formatString string) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		text := d.Get("text").(string)
		formattedText := fmt.Sprintf(formatString, text)
		d.Set("contents", formattedText)
		d.SetId(getID())
		return nil
	}
}

func getID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
