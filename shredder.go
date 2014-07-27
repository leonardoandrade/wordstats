package wordstats

import (
	"fmt"
	"strconv"
	"strings"
)

type Shredder struct {
	termCount     int
	documentCount int
	stopped       bool
	wordStats     *WordStats
}

func newShredder(wordStats *WordStats) *Shredder {
	return &Shredder{0, 0, true, wordStats}
}

func (this *Shredder) worker(textChannel chan string) {
	this.stopped = false
	for text := range textChannel {
		//fmt.Println("doclen:"+strconv.Itoa(len(text)))
		this.termCount = this.termCount + len(text)
		this.documentCount++
		tmp := strings.Split(strings.Replace(strings.ToLower(text), "\n", "", -1), " ")

		for _, tok := range tmp {
			this.wordStats.setTF(tok, this.wordStats.TF(tok)+1)
			//this.wordStats.stats[tok].idf = this.wordStats.stats[tok].idf + 1
		}
		if this.documentCount > 3000 {
			break
		}
	}
	this.stopped = true
}

func (this *Shredder) Run(textChannel chan string, numThreads int) {
	for i := 0; i < numThreads; i++ {
		go this.worker(textChannel)
	}
}

func (this *Shredder) Report() {
	fmt.Println("docs: " + strconv.Itoa(this.documentCount) + " terms: " + strconv.Itoa(this.termCount))
}

func (this *Shredder) Stopped() bool {
	return this.stopped
}
