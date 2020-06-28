resource "gatsby_text_italic" "test_italic_text" { 
  text = "${gatsby_text_bold.test_text.contents} <- that's bold, this isn't"
}
resource "gatsby_text_bold" "test_text" {
  text = "Make this bold"
}
