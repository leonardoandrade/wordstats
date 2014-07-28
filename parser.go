package wordstats

import (
	"encoding/xml"
	"os"
	//"fmt"
	//"strconv"
)

type redirect struct {
	Title string `xml:"title,attr"`
}

type wpPage struct {
	Title string   `xml:"title"`
	Text  string   `xml:"revision>text"`
	Redirect redirect `xml:"redirect"`
}

func validPage(p wpPage) bool {
	return p.Redirect.Title == "" && !(len(p.Title) >= len("Categoria:") && p.Title[:len("Categoria:")] == "Categoria:")
}

func parseFile(filename string, textChannel chan string) {

	f, _ := os.Open(filename)
	decoder := xml.NewDecoder(f)

	for tok, _ := decoder.Token(); tok!=nil; tok,_ = decoder.Token(){

		switch element := tok.(type) {
		case xml.StartElement:
			if element.Name.Local == "page" {
				var p wpPage

				decoder.DecodeElement(&p, &element)
				//fmt.Println("title:"+p.Title+" valid:"+strconv.FormatBool(validPage(p)))
				if validPage(p) {
					textChannel <- p.Text
				}
			}
		}
	}
	close(textChannel)
}
