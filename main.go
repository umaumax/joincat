package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	width      int
	numberFlag bool
	delimChar  string
)

func init() {
	flag.IntVar(&width, "w", -1, "width")
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
		line_buffer := &bytes.Buffer{}
		if numberFlag {
			fmt.Fprintf(line_buffer, "%4d ", nr)
		}
		eofCnt := 0
		for i, scanner := range scanners {
			text := ""
			if scanner.Scan() {
				text = scanner.Text()
			} else {
				eofCnt++
			}
			if i > 0 {
				fmt.Fprintf(line_buffer, delimChar)
			}
			if width < 0 {
				fmt.Fprintf(line_buffer, "%s", text)
			} else {
				fmt.Fprintf(line_buffer, "%"+strconv.Itoa(width)+"s", text)
			}
		}
		fmt.Fprintln(line_buffer)
		if eofCnt == len(scanners) {
			break
		}
		fmt.Fprintf(w, "%s", line_buffer.String())
	}
	fmt.Printf("%s", w.String())
}
