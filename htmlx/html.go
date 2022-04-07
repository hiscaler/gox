package htmlx

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var rxHTML = regexp.MustCompile(`(?s)<sty(.*)/style>|<scr(.*)/script>|<link(.*)/>|<meta(.*)/>|<!--(.*)-->`)

// Strip Clean html tags
// https://stackoverflow.com/questions/55036156/how-to-replace-all-html-tag-with-empty-string-in-golang
func Strip(html string) string {
	html = strings.TrimSpace(html)
	if html != "" {
		html = rxHTML.ReplaceAllString(html, "")
	}
	if html == "" {
		return ""
	}

	const (
		htmlTagStart = 60 // Unicode `<`
		htmlTagEnd   = 62 // Unicode `>`
	)
	// Setup a string builder and allocate enough memory for the new string.
	var builder strings.Builder
	builder.Grow(len(html) + utf8.UTFMax)

	in := false // True if we are inside an HTML tag.
	start := 0  // The index of the previous start tag character `<`
	end := 0    // The index of the previous end tag character `>`

	for i, c := range html {
		// If this is the last character and we are not in an HTML tag, save it.
		if (i+1) == len(html) && end >= start && c != htmlTagStart && c != htmlTagEnd {
			builder.WriteString(html[end:])
		}

		// Keep going if the character is not `<` or `>`
		if c != htmlTagStart && c != htmlTagEnd {
			continue
		}

		if c == htmlTagStart {
			// Only update the start if we are not in a tag.
			// This make sure we strip out `<<br>` not just `<br>`
			if !in {
				start = i
			}
			in = true

			// Write the valid string between the close and start of the two tags.
			builder.WriteString(html[end:start])
			continue
		}
		// else c == htmlTagEnd
		in = false
		end = i + 1
	}
	return strings.TrimSpace(builder.String())
}

func Tag(tag, content string, attributes, styles map[string]string) string {
	var sb strings.Builder
	sb.Grow(len(tag)*2 + len(content) + 5)
	sb.WriteString("<")
	sb.WriteString(tag)
	if len(attributes) > 0 {
		for k, v := range attributes {
			sb.WriteString(" ")
			sb.WriteString(k)
			sb.WriteString(`="`)
			sb.WriteString(v)
			sb.WriteString(`"`)
		}
	}
	if len(styles) > 0 {
		sb.WriteString(` style="`)
		for k, v := range styles {
			sb.WriteString(k)
			sb.WriteString(":")
			sb.WriteString(v)
			sb.WriteString(`;`)
		}
		sb.WriteString(`"`)
	}
	sb.WriteString(">")
	sb.WriteString(content)
	sb.WriteString("</")
	sb.WriteString(tag)
	sb.WriteString(">")
	return sb.String()
}
