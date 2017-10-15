// Program pathescape escapes stdin so it can be safely used in a URL path segment.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

var flagReverse = flag.Bool("u", false, "unescape instead")

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *flagReverse {
		u, err := url.PathUnescape(string(b))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Print(u)
	} else {
		fmt.Print(url.PathEscape(string(b)))
	}
}
