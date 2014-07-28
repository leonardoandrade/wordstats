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
	df int
}

type WordStats struct {
	totalDocs int
	stats map[string]*termStats
}

func (this *WordStats) dumpToFile(fileName string) {
	f, _ := os.Create(fileName)
	for k, _ := range this.stats {
		f.WriteString(k + ";" + strconv.Itoa(this.TF(k)) + ";" + strconv.Itoa(this.DF(k)) + ";" + fmt.Sprintf("%.6f", this.IDF(k)) + "\n")
	}
	f.Close()
}

func ParseWikipedia(fileName string, language Language, outputTermFile string) {
	textChannel := make(chan string)
	go parseFile(fileName, textChannel)
	ws := NewWordStats()
	shredder := newShredder(ws)
	shredder.Run(textChannel, 4)
	for {
		time.Sleep(2 * time.Second)
		shredder.Report()
		if shredder.Stopped() {
			break
		}
	}
	ws.dumpToFile(outputTermFile)
}

func NewWordStats() *WordStats {
	return &WordStats{0, make(map[string]*termStats)}
}

func (this *WordStats) TF(term string) int {
	if val, ok := this.stats[term]; ok {
		return val.tf
	} else {
		return 0
	}
}

func (this *WordStats) DF(term string) int {
	if val, ok := this.stats[term]; ok {
		return val.df
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
		this.stats[key] = &termStats{tf, 0.0, 0}
	}
}

func (this *WordStats) setDF(key string, df int) {
	if val, ok := this.stats[key]; ok {
		val.df = df
	} else {
		this.stats[key] = &termStats{0, 0.0, df}
	}
}
