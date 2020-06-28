resource "gatsby_text_italic" "test_italic_text" { 
  text = "${gatsby_text_bold.test_text.contents} <- that's bold, this isn't"
}
resource "gatsby_text_bold" "test_text" {
  text = "Make this bold"
}

resource "gatsby_text_link" "label_free_link" {
    url = "https://google.ca"
}

resource "gatsby_text_link" "labeled_link" {
    url = "https://google.ca"
    label = "Hello world"
}