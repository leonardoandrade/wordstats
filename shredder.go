package wordstats

import (
	"strings"
	. "sync"
)

type Shredder struct {
	termCount     int
	documentCount int
	stopped       bool
	wordStats     *WordStats
	m             *Mutex
}

type termCounts map[string]int

func newShredder(wordStats *WordStats) *Shredder {
	return &Shredder{0, 0, true, wordStats, &Mutex{}}
}

func (this *Shredder) worker(textChannel chan string, wordCountChannel chan termCounts) {
	this.stopped = false
	for text := range textChannel {
		this.documentCount++
		this.wordStats.totalDocs = this.documentCount
		tmp := strings.Split(ExtractContentFromWikitext(text), " ")
		this.termCount = this.termCount + len(tmp)

		termCounts := make(map[string]int)
		for _, tok := range tmp {
			termCounts[tok] = termCounts[tok] + 1
		}
		wordCountChannel <- termCounts

		if this.wordStats.totalDocs > 5000 {
			break
		}

	}
	this.stopped = true
}

func (this *Shredder) Run(textChannel chan string, wordCountChannel chan termCounts, numThreads int) {
	for i := 0; i < numThreads; i++ {
		go this.worker(textChannel, wordCountChannel)
	}
}

func (this *Shredder) TotalDocsAndTermsProcessed() (int, int) {
	return this.documentCount, this.termCount
}

func (this *Shredder) Stopped() bool {
	return this.stopped
}
