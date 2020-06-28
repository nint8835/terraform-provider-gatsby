package gatsby

import (
	"crypto/sha512"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func textWrapper(formatString string) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		text := d.Get("text").(string)
		formattedText := fmt.Sprintf(formatString, text)
		d.Set("contents", formattedText)
		d.SetId(getID(text))
		return nil
	}
}

func getID(text string) string {
	hash := sha512.New()
	hexString := fmt.Sprintf("%x", hash.Sum([]byte(text)))
	return hexString[:16]
}
