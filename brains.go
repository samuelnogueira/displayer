package main

import "github.com/alteamc/minequery/ping"

type stringList []string

func newStringListFromPlayerNames(response ping.Response) stringList {
	var names stringList
	for _, p := range response.Players.Sample {
		names = append(names, p.Name)
	}

	return names
}

func (m stringList) diff(o stringList) []string {
	res := make([]string, 0)
	for _, p := range m {
		if !o.contains(p) {
			res = append(res, p)
		}
	}

	return res
}

func (m stringList) contains(player string) bool {
	for _, p := range m {
		if p == player {
			return true
		}
	}

	return false
}
