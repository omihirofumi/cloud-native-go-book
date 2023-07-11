package main

type duck struct{}

func (d duck) Says() string {
	return "クワッ！"
}

var Animal duck
