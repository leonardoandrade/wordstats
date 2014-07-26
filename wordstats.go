package wordstats

import (
	"time"
)

type Language int

const(
  DE = iota
  PT
  EN
)

type termStats struct {
  tf int
  idf float32
}

type WordStats struct {
   stats map[string]termStats
}

func ParseWikipedia(fileName string, language Language, outputTermFile string) {
	textChannel := make(chan string)
	go parseFile(fileName, textChannel)

	shredder := newShredder()
	shredder.Run(textChannel, 4)
	for {
		time.Sleep(1 * time.Second)
		shredder.Report()
	}
}

func NewWordStats(termFile string) (*WordStats) {
  //TODO
  return nil
}

func (this *WordStats) TF(term string) (int){
  //TODO
  return 0
}

func (this *WordStats) IDF(term string) (float32){
  //TODO
  return 0.0
}
