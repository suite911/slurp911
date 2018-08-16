package main

import (
	"fmt"
	"log"
	"os"

	"github.com/suite911/slurp911"

	"github.com/suite911/term911/vt"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(usage)
		if len(os.Args) >= 2 {
			os.Exit(1)
		}
		return
	}

	if err := slurp911.Main(os.Args[0], os.Args[3:], os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(usage)
		log.Fatalln(err)
	}
}

var usage string

func init() {
	usage = "usage: " + os.Args[0] + " " + vt.U("PKGNAME") + " " + vt.U("VARNAME")  " " +
		vt.U("KEY") + ":" + vt.U("PATH") + " " + vt.S("[") +
		vt.U("KEY") + ":" + vt.U("PATH") + " " + vt.S("[...]]")
}
