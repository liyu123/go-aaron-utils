package com_utils

import "fmt"

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
