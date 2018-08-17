package slurp

import "io"

func init() {
	self.Init()
}

func Read(p []byte) (n int, err error) {
	n, err = self.Read(p)
	return
}

func Slurp(key, path string) error {
	return self.Slurp(key, path)
}

func Slurped(key string, data []byte) error {
	return self.Slurped(key, data)
}

func SlurpDir(prefix, path string) error {
	return self.SlurpDir(prefix, path)
}

func SlurpFile(key, path string) error {
	return self.SlurpFile(key, path)
}

func SlurpURL(key string, url string) error {
	return self.SlurpURL(key, url)
}

func WriteTo(w io.Writer) (n int64, err error) {
	n, err = self.WriteTo(w)
	return
}

var self Slurper
