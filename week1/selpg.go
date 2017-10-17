package main

import (
	"bufio"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage  int
	endPage    int
	inFilename string
	pageLen    int
	pageType   int
	printDest  string
}

var progname string

func main() {
	var sa selpgArgs

	progname = os.Args[0]

	sa.startPage = -1
	sa.endPage = -1
	sa.pageLen = 72
	sa.pageType = 'l'

	processArgs(&sa)
	processInput(&sa)
}

func processArgs(psa *selpgArgs) {

	start := flag.IntP("startPage", "s", -1, "start page")
	end := flag.IntP("endPage", "e", -1, "end page")
	pageSeparator := flag.BoolP("filter", "f", false, "seperate by \\f")
	pageLen := flag.IntP("pageLenth", "l", 72, "page length")
	dest := flag.StringP("destFile", "d", "", "destination to receive the command")
	flag.Parse()

	if *start < 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid start page\n", progname)
		usage()
		os.Exit(1)
	}
	psa.startPage = *start

	if *end < 1 || *end < psa.startPage {
		fmt.Fprintf(os.Stderr, "%s: invalid end page \n", progname)
		usage()
		os.Exit(2)
	}
	psa.endPage = *end

	if *pageLen < 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid page length\n", progname)
		usage()
		os.Exit(3)
	}
	psa.pageLen = *pageLen

	if *pageSeparator {
		psa.pageType = 'f'
	}

	if *dest != "" {
		psa.printDest = *dest
	}
	if len(flag.Args()) != 0 {
		file := flag.Arg(0)
		psa.inFilename = file
		f, err := os.Open(psa.inFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n",
				progname, psa.inFilename)
			os.Exit(4)
		}
		f.Close()
	}
}

func processInput(sa *selpgArgs) {
	var err error
	fin := os.Stdin
	if sa.inFilename != "" {
		fin, err = os.Open(sa.inFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n",
				progname, sa.inFilename)
			os.Exit(12)
		}
	}

	fout := os.Stdout
	//  =========== disable ===============
	// if sa.printDest != "" {
	// 	outPipe := sa.printDest
	// 	cmd := exec.Command(outPipe)
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	err := cmd.Start()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n",
	// 			progname, outPipe)
	// 		os.Exit(13)
	// 	}
	// 	fout, err = cmd.StdinPipe()
	// }

	var lineCtr int
	var pageCtr int
	inputReader := bufio.NewReader(fin)

	if sa.pageType == 'l' {
		lineCtr = 0
		pageCtr = 1

		for true {
			line, _, err := inputReader.ReadLine()
			if err != nil || line == nil {
				break
			}
			lineCtr = lineCtr + 1
			if lineCtr > sa.pageLen {
				pageCtr = pageCtr + 1
				lineCtr = 1
			}
			if pageCtr >= sa.startPage && pageCtr <= sa.endPage {
				fmt.Fprintf(fout, "%s\n", line)
			}
		}
	} else {
		pageCtr = 1
		for true {
			c, _, err := inputReader.ReadRune()
			if err != nil {
				break
			}

			if c == '\f' {
				pageCtr = pageCtr + 1
			}
			if pageCtr >= sa.startPage && pageCtr <= sa.endPage {
				fmt.Fprintf(fout, "%c", c)
			}
		}
	}

	if pageCtr < sa.startPage {
		fmt.Fprintf(os.Stderr,
			"\n%s: start_page (%d) greater than total pages (%d), no output written\n",
			progname, sa.startPage, pageCtr)
	} else if pageCtr < sa.endPage {
		fmt.Fprintf(os.Stderr,
			"\n%s: end_page (%d) greater than total pages (%d), less output than expected\n",
			progname, sa.endPage, pageCtr)
	}

	fmt.Fprintf(os.Stdout, "%s: done\n", progname)
	fin.Close()
	fout.Close()
}

func usage() {
	fmt.Fprintf(os.Stderr,
		"\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n",
		progname)
}
