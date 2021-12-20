/*
Package pushid generates chronological, 20-character unique IDs

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
*/
package pushid
