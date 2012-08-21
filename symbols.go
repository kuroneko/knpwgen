package main

import (
	"math/rand"
)

var (
	validSymbols = []string {
		"%", "$", "*", "#", "@", "!", "+", "-", ";",
		":", ".", ",", "?" }
)

func RandomSymbol() string {
	return validSymbols[rand.Intn(len(validSymbols))]
}

func RandomDigit() string {
	var outbuf = []byte { '0' }
	outbuf[0] += byte(rand.Intn(10))
	return string(outbuf)
}