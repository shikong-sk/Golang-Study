package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("1,2,3", ",")
	want := []string{"1", "2", "3"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("期望的结果 => %v, 得到的结果 => %v", want, got)
	}
}
