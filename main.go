package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

func main() {

	// is anything coming from stdin?
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// process from stdin
		// fmt.Println("data is being piped to stdin")
	} else {
		// process from flag
		// fmt.Println("stdin is from a terminal")
		targetFlag := flag.String("t", "", "the target")
		flag.Parse()

		if *targetFlag == "" {
			fmt.Printf("\nYou must provide a target e.g., -t http://example.com\n\n")
			os.Exit(1)
		}

		query(*targetFlag)
	}
}

func query(target string) {
	resp, err := http.Get(target)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	// print target
	fmt.Println()
	fmt.Println("Targeting", target)

	// print server software
	fmt.Println()
	fmt.Println("HTTP Server", resp.Header["Server"])

	// print headers
	fmt.Println()
	fmt.Println("Response Headers")
	fmt.Println()

	for name, value := range resp.Header {
		fmt.Println(name+":", value[0])
	}
	fmt.Println()
}