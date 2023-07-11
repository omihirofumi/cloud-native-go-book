package main

type fox struct{}

func (f fox) Says() string {
	return "コンコン"
}

var Animal fox
