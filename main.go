package main

import (
	"math/rand"
	"time"
	"fmt"
	"flag"
)

type PasswordSection int

const (
	S_WORD		= PasswordSection(iota)
	S_PUNCT
	S_DIGIT
	S_ANY		// must be last.
)

var (
	mustWords int	= 2
	mustDigits int	= 1
	mustPunct int	= 0
	totalSections int = 5
)

func init() {
	flag.IntVar(&mustWords, "words", 2, "minimum number of words to include")
	flag.IntVar(&mustDigits, "digits", 1, "minimum number of digits to include")
	flag.IntVar(&mustPunct, "symbols", 0, "minimum number of symbols/punctuation to include")
	flag.IntVar(&totalSections, "sections", 5, "number of sections to use in any password")
}

func (section PasswordSection) Generate() (string) {
	switch section {
	case S_WORD:
		return RandomWord()
	case S_PUNCT:
		return RandomSymbol()
	case S_DIGIT:
		return RandomDigit()
	case S_ANY:
		return PasswordSection(rand.Intn(int(S_ANY))).Generate()
	}
	return ""
}

func SeedRandom() {
	now := time.Now()
	rand.Seed(now.UnixNano())
}

func MakePassword() (password string) {
	/* initialise the sections list */
	var pwsections = make([]PasswordSection, totalSections)
	var ctr int = 0
	for i := 0; i < mustWords; i++ {
		pwsections[ctr] = S_WORD
		ctr++
	}
	for i := 0; i < mustDigits; i++ {
		pwsections[ctr] = S_DIGIT
		ctr++
	}
	for i := 0; i < mustPunct; i++ {
		pwsections[ctr] = S_PUNCT
		ctr++
	}
	for ; ctr < totalSections; ctr++ {
		pwsections[ctr] = S_ANY
	}
	
	/* now, loop */
	for len(pwsections) > 0 {
		sectcount := len(pwsections)
		thisSect := rand.Intn(sectcount)
		password += pwsections[thisSect].Generate()
		
		// remove the section we just used...
		copy(pwsections[thisSect:], pwsections[thisSect+1:])
		pwsections = pwsections[:sectcount-1]
	}
	return password
}

func main() {
	flag.Parse()
	SeedRandom()
	
	if totalSections < (mustWords + mustDigits + mustPunct) {
		fmt.Printf("Warning: totalSections is less than the sum of the mandatory elements.  increasing it to the sum.")
		totalSections = mustWords + mustDigits + mustPunct
	}
	pwout := MakePassword()
	fmt.Printf("%s\n", pwout)
	
}
