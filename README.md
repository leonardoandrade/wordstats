Word Stats
==========

Package to compute word stats by parsing [Wikipedia](http://wikipedia.org) corpus of a given language, written in Go.
Available statistics are:

* Total docs
* Term count - number of times the term appears in corpus
* Document count - number of documents where the term appears
* Inverse document frequency - measurer the "term specificity", implemented as ***IDF = log(Total docs / Document count).*** [WP has more information on this metric](http://en.wikipedia.org/wiki/Tf%E2%80%93idf)

Useful for search engine document ranking, spellcheckers, etc...

Reading stats
-------------

After loading the csv file, the metrics for a given term can be retrieved.
Example of command line executable:


```go
package main

import (
	"os"
	"fmt"
	"github.com/leonardoandrade/wordstats"
)

func main() {
	if len(os.Args) == 3 {
		ws := wordstats.NewWordStats()
		ws.Load(os.Args[1])
		w := os.Args[2]
		fmt.Printf("Stats for word '%s':\n", w)
		if ws.Exists(w) {
			fmt.Printf("TC %d, DC: %d, IDF: %f \n", ws.TC(w), ws.DC(w), ws.IDF(w))
		} else {
			fmt.Printf("term '%s' does not exist in the corpus analized", w)
		}
	} else {
		fmt.Println("usage: " + os.Args[0] + " <csv-stats-file> <word>")
	}
}
```

Computing stats
--------------

Example of command line executable:

```go
package main

import (
	"os"
	"fmt"
	"github.com/leonardoandrade/wordstats"
)

func main() {
	if len(os.Args) == 3 {
		wordstats.ParseWikipedia(os.Args[1], wordstats.PT, os.Args[2])
	} else {
		fmt.Println("usage: "+os.Args[0]+" <file to parse> <output csv file>")
	}
}
```

The process was optimized to scale-up in a SMP machine, and explores the parallel programming capabilities of Go. This is opposed to a more common scale-out approach using map-reducers such as Hadoop.