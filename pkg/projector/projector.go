package projector

import "path"

type Data struct {
	Projector map[string]map[string]string `json:"projector`
}

type Projector struct {
	config Config
	data   Data
}

func (p *Projector) GetValue(key string) (string, bool) {

	curr := p.config.Pwd
	prev := ""

	out := ""
	found := false
	for curr != prev {

		if dir, ok := p.data.Projector[curr]; ok {
			if value, ok := dir[key]; ok {
				out = value
				found = true
				break
			}
		}

		prev = curr
		curr = path.Dir(curr)
	}
	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	out := map[string]string{}

	paths := []string{}

	curr := p.config.Pwd
	prev := ""

	for curr != prev {
		paths = append(paths, curr)
		prev = curr
		curr = path.Dir(curr)
	}

	for i := len(paths) - 1; i >= 0; i-- {

		if dir, ok := p.data.Projector[paths[i]]; ok {
			for k, v := range dir {
				out[k] = v
			}
		}
	}

	return out
}
