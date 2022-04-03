package main

import (
	"reflect"
	"testing"
)

func Test_stringToMinecraftServerURL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{s: "minecraft.example.com"}, want: "minecraft.example.com:25565"},
		{args: args{s: "minecraft.example.com."}, want: "minecraft.example.com.:25565"},
		{args: args{s: "minecraft.example.com:25565"}, want: "minecraft.example.com:25565"},
		{args: args{s: "minecraft.example.com:8080"}, want: "minecraft.example.com:8080"},
		{args: args{s: "minecraft.example.com.:8080"}, want: "minecraft.example.com.:8080"},
		{args: args{s: "//minecraft.example.com"}, want: "minecraft.example.com:25565"},
		{args: args{s: "//minecraft.example.com:25565"}, want: "minecraft.example.com:25565"},
		{args: args{s: "//minecraft.example.com:8080"}, want: "minecraft.example.com:8080"},
		{args: args{s: "minecraft://minecraft.example.com"}, want: "minecraft.example.com:25565"},
		{args: args{s: "minecraft://minecraft.example.com/"}, want: "minecraft.example.com:25565"},
		{args: args{s: "minecraft://minecraft.example.com:25565"}, want: "minecraft.example.com:25565"},
		{args: args{s: "minecraft://minecraft.example.com:25565/"}, want: "minecraft.example.com:25565"},
		{args: args{s: "minecraft://minecraft.example.com:8080"}, want: "minecraft.example.com:8080"},
		{args: args{s: "minecraft://minecraft.example.com:8080/"}, want: "minecraft.example.com:8080"},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			result := stringToMinecraftServerURL(tt.args.s)
			got := result.Hostname() + ":" + result.Port()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToMinecraftServerURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
