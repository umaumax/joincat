package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	numberFlag bool
	delimChar  string
)

func init() {
	flag.BoolVar(&numberFlag, "n", false, "line number flag")
	flag.StringVar(&delimChar, "d", "|", "delim char")
}

func main() {
	flag.Parse()
	args := flag.Args()

	var err error
	fs := make([]*os.File, len(args))
	scanners := make([]*bufio.Scanner, len(args))
	for i, v := range args {
		fs[i], err = os.Open(v)
		if err != nil {
			log.Fatalln(err)
		}
		defer fs[i].Close()
		scanner := bufio.NewScanner(fs[i])
		scanners[i] = scanner
	}
	if len(scanners) == 0 {
		return
	}
	w := &bytes.Buffer{}
	for nr := 1; ; nr++ {
		if numberFlag {
			fmt.Fprintf(w, "%4d ", nr)
		}
		eofCnt := 0
		for i, scanner := range scanners {
			if scanner.Scan() {
				text := scanner.Text()
				if i > 0 {
					fmt.Fprintf(w, delimChar)
				}
				fmt.Fprintf(w, "%s", text)
			} else {
				eofCnt++
			}
		}
		fmt.Fprintln(w)
		if eofCnt == len(scanners) {
			break
		}
	}
	fmt.Printf("%s", w.String())
}
