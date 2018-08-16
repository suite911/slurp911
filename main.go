package main

import (
	"fmt"
	"log"
	"os"

	"github.com/suite911/slurp911/slurp"
)

func main() {
	if len(os.Args) < 4 {
		usage()
		return
	}
	s := slurp.New(os.Args[1], os.Args[2])
	failed := false
	for _, arg := range os.Args[3:] {
		kv := strings.SplitN(arg, ":", 2)
		if len(kv) != 2 {
			badUsage(nil)
		}
		if err := s.Slurp(kv[0], kv[1]); err != nil {
			log.Println("Failed to slurp \""+kv[1]+"\"")
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
}

func badUsage(err error) {
	usage()
	if err != nil {
		log.Println(err)
	}
	os.Exit(1)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: "+os.Args[0]+" PKGNAME VARNAME KEY:PATH [KEY:PATH [...]]")
}
