package slurp

import (
	"sync"

	"github.com/suite911/error911/onfail"
)

func init() {
	self.Init()
}

func Read(p []byte) (n int, err error) {
	n, err = self.Read(p)
	return
}

func Slurp(key string, data []byte) error {
	return self.Slurp(key, data)
}

func SlurpDir(key string, pathElems ...string) error {
	return self.SlurpDir(key, pathElems...)
}

func SlurpFile(key string, pathElems ...string) error {
	return self.SlurpFile(key, pathElems...)
}

func SlurpURL(key string, url string) error {
	return self.SlurpURL(key, url)
}

func WriteTo(w io.Writer) (n int64, err error) {
	n, err = self.WriteTo(w)
	return
}

var self Slurper
