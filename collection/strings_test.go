package collection

import (
	"reflect"
	"testing"
)

func Test_Strings_diff(t *testing.T) {
	type args struct {
		players []string
	}
	tests := []struct {
		name string
		m    Strings
		args args
		want []string
	}{
		{name: "negative Diff", m: Strings{}, args: args{players: []string{"nakiski"}}, want: []string{}},
		{name: "one", m: Strings{"nakiski", "DataDisruptor"}, args: args{players: []string{"nakiski"}}, want: []string{"DataDisruptor"}},
		{name: "no Diff", m: Strings{"nakiski"}, args: args{players: []string{"nakiski"}}, want: []string{}},
		{name: "empty arg", m: Strings{"nakiski"}, args: args{players: []string{}}, want: []string{"nakiski"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Diff(tt.args.players); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
