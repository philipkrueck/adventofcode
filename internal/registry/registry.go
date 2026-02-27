package registry

type Runner = func(string) string

type key struct {
	year, day, part int
}

type entry struct {
	Run   Runner
	Input string
}

var registry = map[key]entry{}

func Register(year, day, part int, fn Runner, input string) {
	k := key{year, day, part}
	registry[k] = entry{fn, input}
}

func Get(year, day, part int) (Runner, string, bool) {
	k := key{year, day, part}
	v, ok := registry[k]
	if !ok {
		return nil, "", false
	}
	return v.Run, v.Input, true
}
