package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

func lexer(s string) ([]ltoken,error) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Uncatch lexer error", r)
        }
    }()
	i := 0
	tokens := make([]ltoken,0)
	for i<len(s) {
		fmt.Println(i,len(s),tokens)
		for s[i]==' ' {
			i++
		}
		if s[i] == '(' {
			tok := ltoken{"lparen",0}
			tokens = append(tokens,tok)
			i++
			continue
		}
		if s[i] == ')' {
			tok := ltoken{"rparen",0}
			tokens = append(tokens,tok)
			i++
			continue
		}
		if s[i] == '"' {
			j := i+1
			for s[j] != '"' {
				j++
			}
			j++
			tok := ltoken{"string",string(s[i:j])}
			tokens = append(tokens,tok)
			i = j
			continue
		}
		//Number parsing
		if strings.Contains("0123456789.",string(s[i])) {
			j := i
			for strings.Contains("0123456789.",string(s[j])) {
				j++
			}
			if strings.Contains(string(s[i:j]),".") {
				value, err := strconv.ParseFloat(s[i:j],64)
				if err!=nil {
					fmt.Println(err)
				}
				tok := ltoken{"float",value}
				tokens = append(tokens,tok)
			} else {
				value, err := strconv.ParseInt(s[i:j],10,64)
				if err!=nil {
					fmt.Println(err)
				}
				tok := ltoken{"int",value}
				tokens = append(tokens,tok)
			}
			i = j
			continue
		}
		j := i
		for s[j] != ' ' && s[j] != ')' {
			j++
			if j > len(s) {
				return tokens, errors.New("Malformed exp")
			}
		}
		tok := ltoken{"name",s[i:j]}
		tokens = append(tokens,tok)
		i = j
		continue
	}
	return tokens, nil
}
