package slurp911

import "github.com/suite911/slurp911/slurp"

func Main(programName string, pairs []string, opts ...string) error {
	var s slurp.Slurper
	s.Init(opts...)
	failed := false
	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		var k, v string
		switch len(kv) {
		case 1:
			k, v = "", kv[0]
		case 2:
			k, v = kv[0], kv[1]
		default:
			return errors.New("Bad usage in argument \""+pair+"\"")
		}
		if err := s.Slurp(k, v); err != nil {
			return err
		}
	}
	return nil
}
