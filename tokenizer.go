package wordstats

import (
	"regexp"
	"strings"
)


var file_link  *regexp.Regexp = regexp.MustCompile("\\[\\[ficheiro:.*\\]\\]")

var image_link  *regexp.Regexp = regexp.MustCompile("\\[\\[imagem:.*\\]\\]")

var reference  *regexp.Regexp = regexp.MustCompile("<ref.*>.*</ref>")

var elements_to_convert_to_space *regexp.Regexp = regexp.MustCompile("[\\-,\\.*\\n'\\\"|\\(\\):;]+")

var tags_to_remove  *regexp.Regexp = regexp.MustCompile("\\{\\{[^\\{\\}]*\\}\\}")

var link_markers *regexp.Regexp = regexp.MustCompile("[\\[\\]]+")

var header_markers *regexp.Regexp = regexp.MustCompile("[=]+")



func ExtractContentFromWikitext(input string) (string) {
	var ret string
	ret =strings.ToLower(input)
	ret = file_link.ReplaceAllString(ret, " ")
	ret = image_link.ReplaceAllString(ret, " ")
	ret = reference.ReplaceAllString(ret, " ")
	ret = elements_to_convert_to_space.ReplaceAllString(ret, " ")
	ret = file_link.ReplaceAllString(ret, " ")
	ret = tags_to_remove.ReplaceAllString(ret, " ")
	ret = tags_to_remove.ReplaceAllString(ret, " ")
	ret = tags_to_remove.ReplaceAllString(ret, " ")
	ret = link_markers.ReplaceAllString(ret, " ")
	ret = header_markers.ReplaceAllString(ret, " ")
	return ret
}
