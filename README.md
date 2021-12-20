# Push ID

PushID is a Go library which generates chronological, 20-character unique IDs. It 
implements Firebase's push IDs, what are Push IDs? 

*Push IDs are string identifiers that are generated client-side. They are a
 combination of a timestamp and some random bits. The timestamp ensures they
 are ordered chronologically, and the random bits ensure that each ID is
 unique, even if thousands of people are creating push IDs at the same time.*

```
{
    "messages": {
        "-JhLeOlGIEjaIOFHR0xd": "Hello there!",
        "-JhQ76OEK_848CkIFhAq": "Push IDs are pretty magical.",
        "-JhQ7APk0UtyRTFO9-TS": "Look a white rabbit!"
    }
}
```

# Installation

```
$ go get github.com/umahmood/pushid
```

# Usage

```
package main

import (
    "fmt"
    
    "github.com/umahmood/pushid"
)

func main() {
    id := pushid.New()
    for i := 0; i < 10; i++ { 
        pid, err := id.Generate()
        if err != nil {
            //...
        }
        fmt.Println(pid)
    }
}
```
Output:
```
-MrOKZ8QmUqMN34sp0Ne
-MrOKZA2OJ1HB_vkglmr
-MrOKZBfU9x7emWkX5DX
-MrOKZDIh7Voefkx01eK
-MrOKZEsM13vVmJjq4r-
-MrOKZGU5AYs_lWs7Z8t
-MrOKZI6VGY4I_rUqtyw
-MrOKZJirH5rg1F40C8R
-MrOKZLLqOUswv_-zPE6
-MrOKZMz_9QG-zOI3eqF
```

# Documentation

> https://pkg.go.dev/github.com/umahmood/pushid

# Resources

[https://firebase.googleblog.com/2015/02/the-2120-ways-to-ensure-unique_68.html]()

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
