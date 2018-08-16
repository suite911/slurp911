package slurp

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Slurper struct {
	mutex sync.RWMutex
	slurp map[string][]byte

	pkgName, varName string
}

func New(opts ...string) *Slurper {
	return new(Slurper).Init(opts...)
}

func (s *Slurper) Init(opts ...string) *Slurper {
	pkgName := "main"
	varName := "Slurped"
	switch len(opts) {
	case 2:
		varName = opts[1]
		fallthrough
	case 1:
		pkgName = opts[0]
		fallthrough
	case 0:
	default:
		panic("Usage: (*Slurper).Init([pkgName string[, varName string]]")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.varName = varName
	s.pkgName = pkgName
	s.slurp = make(map[string][]byte)
	return s
}

func (s *Slurper) Read(p []byte) (n int, err error) {
	var b bytes.Buffer
	_, err = s.WriteTo(b)
	if err != nil {
		return
	}
	var n64 int64
	n64, err = io.Copy(p, b)
	n = int(n64)
	return
}

func (s *Slurper) Slurp(key string, data []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.slurp[key]; ok {
		return errors.New("Already slurp \""+key+"\"!")
	}
	s.slurp[key] = data
	return nil
}

func (s *Slurper) SlurpDir(prefix string, pathElems ...string) error {
	d, err := ioutil.ReadDir(filepath.Join(pathElems...))
	if err != nil {
		return err
	}
	for _, f := range d {
		fn := f.Name()
		key := prefix + filepath.Base(fn)
		if err := s.SlurpFile(key, fn); err != nil {
			return err
		}
	}
	return nil
}

func (s *Slurper) SlurpFile(key string, pathElems ...string) error {
	path := filepath.Join(pathElems...)
	if len(key) < 1 {
		key = path
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return s.Slurp(key, b)
}

func (s *Slurper) SlurpURL(key, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return s.Slurp(key, b)
}

func (s *Slurper) WriteTo(w io.Writer) (n int64, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	var nn int
	nn, err = io.WriteString(w, "package "+s.pkgName+"\n\nvar "+s.varName+" = map[string][]byte {")
	n += int64(nn)
	if err != nil {
		return
	}
	for k, v := range s.slurp {
		nn, err = io.WriteString(w, "\n\t\""+k+"\": []ubyte{")
		n += int64(nn)
		if err != nil {
			return
		}
		for i, b := range v {
			sep = " 0x"
			if i%8 == 0 {
				sep = "\n\t\t0x"
			}
			nn, err = io.WriteString(w, sep+strconv.FormatInt(int64(b), 16)+",")
			n += int64(nn)
			if err != nil {
				return
			}
		}
		nn, err = io.WriteString(w, "\n\t},")
		n += int64(nn)
		if err != nil {
			return
		}
	}
	nn, err = io.WriteString(w, "\n}\n")
	n += int64(nn)
	if err != nil {
		return
	}
}
