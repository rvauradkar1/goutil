package main

import (
	"reflect"
	"testing"
)

func Test_s1(t *testing.T) {
	tests := []struct {
		name string
		want chan string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s1(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("s1() = %v, want %v", got, tt.want)
			}
		})
	}
}
