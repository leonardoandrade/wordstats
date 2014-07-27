package wordstats

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Language int

const (
	DE = iota
	PT
	EN
)

type termStats struct {
	tf  int
	idf float64
}

type WordStats struct {
	stats map[string]termStats
}

func (this *WordStats) dumpToFile(fileName string) {
	f, _ := os.Create(fileName)
	for k, v := range this.stats {
		f.WriteString(k + ";" + strconv.Itoa(v.tf) + ";" + fmt.Sprintf("%.6f", v.idf) + "\n")
	}
	f.Close()
}

func ParseWikipedia(fileName string, language Language, outputTermFile string) {
	textChannel := make(chan string)
	go parseFile(fileName, textChannel)
	ws := &WordStats{make(map[string]termStats)}
	shredder := newShredder(ws)
	shredder.Run(textChannel, 4)
	for {
		time.Sleep(1 * time.Second)
		shredder.Report()
		if shredder.Stopped() {
			break
		}
	}
	ws.dumpToFile(outputTermFile)
}

func NewWordStats(termFile string) *WordStats {
	//TODO
	return nil
}

func (this *WordStats) TF(term string) int {
	if val, ok := this.stats[term]; ok {
		return val.tf
	} else {
		return 0
	}
}

func (this *WordStats) IDF(term string) float32 {
	//TODO
	return 0.0
}

func (this *WordStats) setTF(key string, tf int) {
	if val, ok := this.stats[key]; ok {
		val.tf = tf
	} else {
		this.stats[key] = termStats{tf, 0.0}
	}
}
