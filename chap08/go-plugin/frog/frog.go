package main

type frog struct{}

func (f frog) Says() string {
	return "ゲロゲロ"
}

var Animal frog
