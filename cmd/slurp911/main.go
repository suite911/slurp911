package main

import (
	"fmt"
	"log"
	"os"

	"github.com/suite911/slurp911"
)

func main() {
	if len(os.Args) < 4 {
		usage()
		if len(os.Args) >= 2 {
			os.Exit(1)
		}
		return
	}

	if err := slurp911.Main(os.Args[0], os.Args[3:], os.Args[1], os.Args[2]); err != nil {
		usage()
		log.Fatalln(err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: "+os.Args[0]+" PKGNAME VARNAME KEY:PATH [KEY:PATH [...]]")
}
