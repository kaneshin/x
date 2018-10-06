package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kaneshin/x"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: suddendeath [OPTION]... [FILE]
＿人人人人人人＿
＞　突然の死　＜
￣ＹＹＹＹＹＹ￣
`)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func run() error {
	var name string
	if args := flag.Args(); len(args) > 0 {
		name = args[0]
	}

	var r io.Reader
	switch name {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r = f
	}

	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	os.Stdout.Write(x.SuddenDeath(string(src)).Bytes())
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
