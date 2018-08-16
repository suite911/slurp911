package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/suite911/slurp911/slurp"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	var err error
	for _, arg := range os.Args[1:] {
		if _, err = url.ParseRequestURI(arg); err == nil {
			slurp.SlurpURL(arg)
		} else {
			fi *os.FileInfo
			if fi, err = os.Stat(arg); err != nil {
				badUsage(err)
				continue
			}
			if fi.IsDir(arg) {
				slurp.SlurpDir(arg)
			} else {
				slurp.SlurpFile(arg)
			}
		}
	}
}

func badUsage(err error) {
	usage()
	log.Println(err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: "+os.Args[0]+" PATHS...")
}
