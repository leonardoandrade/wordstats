Word Stats
==========

Package to compute word stats by parsing [Wikipedia](http://wikipedia.org). The term frequency and the inverse document frequency are computed.
__Work in progress.__
 
Computing stats
--------------

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

Reading stats
-------------

__TODO__


