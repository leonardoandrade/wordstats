package wordstats

import (
	"encoding/xml"
	"os"
	"regexp"
	"strings"
)

type redirect struct {
	Title string `xml:"title,attr"`
}

type wpPage struct {
	Title string   `xml:"title"`
	Text  string   `xml:"revision>text"`
	Redirect redirect `xml:"redirect"`
}
/*
if a page is a redirect or a special page, such as "category:" or "anexo:", is not valid for indexing
 */
func validPage(p wpPage) bool {
	return p.Redirect.Title == "" && regexp.MustCompile("[a-z]:.*").MatchString(strings.ToLower(p.Title)) == false
}

func parseFile(filename string, textChannel chan string) {

	f, err := os.Open(filename)
	if err != nil {
		panic("cannot open file '"+filename+"'")
	}
	decoder := xml.NewDecoder(f)

	for tok, _ := decoder.Token(); tok!=nil; tok,_ = decoder.Token(){

		switch element := tok.(type) {
		case xml.StartElement:
			if element.Name.Local == "page" {
				var p wpPage

				decoder.DecodeElement(&p, &element)
				if validPage(p) {
					textChannel <- p.Text
				}
			}
		}
	}
	close(textChannel)
}
