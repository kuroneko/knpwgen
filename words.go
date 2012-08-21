package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strings"
	"unicode"
	"math/rand"
	"flag"
)

var (
	words		[]string
	wordFile	string = "/usr/share/dict/words"
)

func init() {
	flag.StringVar(&wordFile, "wordfile", "/usr/share/dict/words", "File to read word list from")
}

func NonLetter(c rune) bool {
	return !unicode.IsLetter(c)
}

func initWords() (err error) {
	fh, err := os.Open(wordFile)
	if err != nil {
		return err
	}
	defer fh.Close()
	bufreader := bufio.NewReader(fh)
	words = make([]string, 0)

	var line []byte
	for {
		line, _, err = bufreader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		thisline := string(line)
		if -1 != strings.IndexFunc(thisline, NonLetter) {
			continue
		}
		words = append(words, thisline)
	}

	return nil
}

func RandomWord() string {
	if words == nil {
		err := initWords()
		if err != nil {
			fmt.Printf("ERR: %s reading word list", err)
			os.Exit(1)
		}
	}
	return words[rand.Intn(len(words))]
}
