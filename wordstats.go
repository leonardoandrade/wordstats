/*
  package to compute word statistics from large document corpus
*/

package wordstats

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
	//"io"
	"runtime"
	//"strings"
	"strings"
)

const REPORT = true

type Language int

const (
	DE = iota
	PT
	EN
)

const (
	TC_TRESHOLD = 500
	MAX_PROCS   = 9
)

type termStats struct {
	tc int // term count
	dc int // document count
}

type WordStats struct {
	totalDocs int
	stats     map[string]*termStats
}

func ParseWikipedia(fileName string, language Language, outputTermFile string) {
	runtime.GOMAXPROCS(MAX_PROCS)

	textChannel := make(chan string)
	wordCountChannel := make(chan termCounts)

	go parseFile(fileName, textChannel)
	ws := NewWordStats()
	shredder := newShredder(ws)
	shredder.Run(textChannel, wordCountChannel, 4)

	//collector goroutine
	go func() {
		for tc := range wordCountChannel {
			for term, ct := range tc {
				ws.setTC(term, ws.TC(term)+ct)
				ws.setDC(term, ws.DC(term)+1)
			}
		}
	}()

	start := time.Now()
	for {
		time.Sleep(2 * time.Second)
		if REPORT {
			docs, terms := shredder.TotalDocsAndTermsProcessed()
			var elapsed time.Duration = time.Since(start)
			var kTermsPerSecond int = int((float64(terms) / elapsed.Seconds()) / 1000)
			fmt.Printf("%d docs, %d terms processed, %dk terms per second \n", docs, terms, kTermsPerSecond)

		}
		if shredder.Stopped() {
			break
		}
	}
	ws.dumpToFile(outputTermFile, TC_TRESHOLD)
}

func NewWordStats() *WordStats {
	return &WordStats{0, make(map[string]*termStats)}
}

/**
Read from file generated in the previuos function
*/
func (this *WordStats) Load(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic("cannot read file " + filepath)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	firstLine := scanner.Text()

	this.totalDocs, _ = strconv.Atoi((string)(firstLine))
	count := 1
	for scanner.Scan() {
		count = count + 1
		tok := strings.Split(scanner.Text(), ";")
		if len(tok) < 3 {
			panic("cannot parse line '" + scanner.Text() + "' in position " + strconv.Itoa(count))
		}
		tc, err := strconv.Atoi(tok[1])
		if err != nil {
			panic("cannot convert value " + tok[1] + " to int")
		}
		dc, err := strconv.Atoi(tok[2])
		if err != nil {
			panic("cannot convert value " + tok[2] + " to int")
		}

		this.setTC(tok[0], tc)
		this.setDC(tok[0], dc)
	}
}

/*
 Comma separated value (CSV) file.
 The first line contains the total documents
 The following lines contain the triple <word; term-count; document count>
*/
func (this *WordStats) dumpToFile(fileName string, thresholdTF int) {
	f, _ := os.Create(fileName)
	f.WriteString(strconv.Itoa(this.totalDocs) + "\n")
	for k, _ := range this.stats {
		if this.TC(k) >= thresholdTF {
			f.WriteString(k + ";" + strconv.Itoa(this.TC(k)) + ";" + strconv.Itoa(this.DC(k)) + "\n")
		}
	}
	f.Close()
}

func (this *WordStats) TC(term string) int {
	if val, ok := this.stats[term]; ok {
		return val.tc
	} else {
		return 0
	}
}

func (this *WordStats) setTC(key string, tc int) {
	if val, ok := this.stats[key]; ok {
		val.tc = tc
	} else {
		this.stats[key] = &termStats{tc, 0}
	}
}

func (this *WordStats) DC(term string) int {
	if val, ok := this.stats[term]; ok {
		return val.dc
	} else {
		return 0
	}
}

func (this *WordStats) setDC(term string, dc int) {
	if val, ok := this.stats[term]; ok {
		val.dc = dc
	} else {
		this.stats[term] = &termStats{dc, 0}
	}
}

func (this *WordStats) Exists(term string) bool {
	_, ok := this.stats[term]
	return ok
}

func (this *WordStats) IDF(term string) float64 {
	return math.Log(float64(this.totalDocs) / float64(this.DC(term)))
}
