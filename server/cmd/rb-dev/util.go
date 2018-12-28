package main

import (
	"fmt"
	"runtime"

	"golang.org/x/exp/rand"
)

func randString(len int) string {
	abc := "abcdefghijklmnopqrstuvwxyz"

	s := ""
	for i := 0; i < len; i++ {
		s += string(abc[rand.Int()%26])
	}

	return s
}

func pe(e error) {
	if e != nil {
		panic(fmt.Sprintf("%s: %s", WAI(2), e.Error()))
	}
}

// WAI ...
func WAI(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("%s:%d", file, line)
}
