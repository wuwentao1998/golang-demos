package tool

import (
	"testing"
)

func TestFuncName(t *testing.T) {
	tests := []struct {
		name string
		f    interface{}
		want string
	}{
		{
			"1",
			PrintStackTrace,
			"util.PrintStackTrace",
		},
		{
			"2",
			f1,
			"util.f1",
		},
		{
			"3",
			nil,
			"",
		},
		{
			"4",
			"",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuncName(tt.f); got != tt.want {
				t.Errorf("FuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncPath(t *testing.T) {
	tests := []struct {
		name string
		f    interface{}
		want string
	}{
		{
			"1",
			PrintStackTrace,
			"github.com/wuwentao1998/golang-demos/tool.PrintStackTrace",
		},
		{
			"2",
			1,
			"",
		},
		{
			"3",
			nil,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuncPath(tt.f); got != tt.want {
				t.Errorf("FuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}
