package collection

type Strings []string

func (m Strings) Diff(o Strings) []string {
	res := make([]string, 0)
	for _, p := range m {
		if !o.contains(p) {
			res = append(res, p)
		}
	}

	return res
}

func (m Strings) contains(player string) bool {
	for _, p := range m {
		if p == player {
			return true
		}
	}

	return false
}
