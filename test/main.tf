resource "gatsby_text_italic" "so_cool" {
  text = "so cool"
}
resource "gatsby_text_bold" "bold_text" {
  text = "make this bold"
}

resource "gatsby_text_link" "label_free_link" {
  url = "https://google.ca"
}

resource "gatsby_text_link" "labeled_link" {
  url   = "https://google.ca"
  label = "I have a label!"
}

resource "gatsby_text_heading" "demo_heading" {
  text = "Hello world"
}

resource "gatsby_text_heading" "subheading" {
  text  = "Look at these cool links"
  level = 2
}

output "demo_text" {
  value = <<EOF
${gatsby_text_heading.demo_heading.contents}

Look at my cool website, isn't it ${gatsby_text_italic.so_cool.contents}?

I can even do things like ${gatsby_text_bold.bold_text.contents}!

${gatsby_text_heading.subheading.contents}
${gatsby_text_link.label_free_link.contents}
${gatsby_text_link.labeled_link.contents}
    EOF
}
