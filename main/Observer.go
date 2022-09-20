package main

type Observer interface {
	eventListener(v []string)
}
