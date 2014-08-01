package wordstats

import (
	"regexp"
	"strings"
)

var generic_link  *regexp.Regexp = regexp.MustCompile("\\[\\[[a-zaàá]+:.*\\]\\]")

var tables *regexp.Regexp = regexp.MustCompile("\\{\\|[^\\{\\}]*\\|\\}")

var html_comment  *regexp.Regexp = regexp.MustCompile("<!--.*-->")

var reference  *regexp.Regexp = regexp.MustCompile("<ref.*>.*</ref>")

var elements_to_convert_to_space *regexp.Regexp = regexp.MustCompile("[,\\.*\\n'\\\"|\\(\\):;]+")

var tags_to_remove  *regexp.Regexp = regexp.MustCompile("\\{\\{[^\\{\\}]*\\}\\}")

var link_markers *regexp.Regexp = regexp.MustCompile("[\\[\\]]+")

var header_markers *regexp.Regexp = regexp.MustCompile("[=]+")


func ExtractContentFromWikitext(input string) (string) {
	var ret string
	ret =strings.ToLower(input)
	ret = generic_link.ReplaceAllString(ret, " ")
	ret = html_comment.ReplaceAllString(ret, " ")
	ret = reference.ReplaceAllString(ret, " ")
	ret = elements_to_convert_to_space.ReplaceAllString(ret, " ")
	ret = tags_to_remove.ReplaceAllString(ret, " ")
	ret = tags_to_remove.ReplaceAllString(ret, " ")
	ret = tags_to_remove.ReplaceAllString(ret, " ")
	ret = tables.ReplaceAllString(ret, " ")
	ret = tables.ReplaceAllString(ret, " ")
	ret = tables.ReplaceAllString(ret, " ")
	ret = link_markers.ReplaceAllString(ret, " ")
	ret = header_markers.ReplaceAllString(ret, " ")
	return ret
}
