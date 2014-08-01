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
	m *Mutex
}

func newShredder(wordStats *WordStats) *Shredder {
	return &Shredder{0, 0, true, wordStats, &Mutex{}}
}

func (this *Shredder) worker(textChannel chan string) {
	this.stopped = false
	for text := range textChannel {
		//fmt.Println("doclen:"+strconv.Itoa(len(text)))

		this.documentCount++
		this.wordStats.totalDocs=this.documentCount
		tmp := strings.Split(ExtractContentFromWikitext(text), " ")
		this.termCount = this.termCount + len(tmp)

		termSet := make(map[string]bool)

		for _, tok := range tmp {
			//fmt.Println(tok)
			termSet[tok] = true
			this.m.Lock()
			this.wordStats.setTC(tok, this.wordStats.TC(tok)+1)
			this.m.Unlock()
		}

		for k, _ := range termSet {
			this.m.Lock()
			this.wordStats.setDC(k, this.wordStats.DC(k)+1)
			this.m.Unlock()
		}


	}
	this.stopped = true
}

func (this *Shredder) Run(textChannel chan string, numThreads int) {
	for i := 0; i < numThreads; i++ {
		go this.worker(textChannel)
	}
}

func (this *Shredder) TotalDocsAndTermsProcessed() (int, int) {
	return this.documentCount, this.termCount
}


func (this *Shredder) Stopped() bool {
	return this.stopped
}
