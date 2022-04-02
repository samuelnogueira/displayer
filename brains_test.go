package main

import (
	"reflect"
	"testing"
)

func Test_brains_contains(t *testing.T) {
	type args struct {
		player string
	}
	tests := []struct {
		name string
		m    stringList
		args args
		want bool
	}{
		{name: "empty", m: stringList{}, args: args{player: "nakiski"}, want: false},
		{name: "not contains", m: stringList{"nakiski"}, args: args{player: "DataDisruptor"}, want: false},
		{name: "contains", m: stringList{"nakiski"}, args: args{player: "nakiski"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.contains(tt.args.player); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_brains_diff(t *testing.T) {
	type args struct {
		players []string
	}
	tests := []struct {
		name string
		m    stringList
		args args
		want []string
	}{
		{name: "negative diff", m: stringList{}, args: args{players: []string{"nakiski"}}, want: []string{}},
		{name: "one", m: stringList{"nakiski", "DataDisruptor"}, args: args{players: []string{"nakiski"}}, want: []string{"DataDisruptor"}},
		{name: "no diff", m: stringList{"nakiski"}, args: args{players: []string{"nakiski"}}, want: []string{}},
		{name: "empty arg", m: stringList{"nakiski"}, args: args{players: []string{}}, want: []string{"nakiski"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.diff(tt.args.players); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
