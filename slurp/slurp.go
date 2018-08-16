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

type SlurpType struct {
	slurped map[string][]byte
	mutex   sync.RWMutex
}

func New() *SlurpType {
	return new(SlurpType).Init()
}

func (s *SlurpType) Init() *SlurpType {
	s.slurped = make(map[string][]byte)
	return s
}

func (s *SlurpType) Slurp(key string, data []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.slurped[key]; ok {
		return errors.New("Already slurped \""+key+"\"!")
	}
	s.slurped[key] = data
	return nil
}

func (s *SlurpType) SlurpDir(prefix string, pathElems ...string) error {
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

func (s *SlurpType) SlurpFile(key string, pathElems ...string) error {
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

func (s *SlurpType) SlurpURL(key, url string) error {
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
