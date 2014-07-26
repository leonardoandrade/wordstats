package wordstats

import (
	"fmt"
	"strconv"

)

type Shredder struct {
	termCount int
	documentCount int
}

func newShredder() (*Shredder) {
	return &Shredder{0,0}
}

func (this *Shredder) worker(textChannel chan string) {
	for text := range(textChannel) {
		//fmt.Println("doclen:"+strconv.Itoa(len(text)))
		this.termCount = this.termCount + len(text)
		this.documentCount++
	}
}

func (this *Shredder) Run(textChannel chan string, numThreads int) {
	for i := 0; i< numThreads; i++ {
		go this.worker(textChannel)
	}
}

func (this *Shredder) Report() {
	fmt.Println("docs: "+strconv.Itoa(this.documentCount)+" terms: "+strconv.Itoa(this.termCount))
}
