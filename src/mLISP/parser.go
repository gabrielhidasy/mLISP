package main

import (
	"fmt"
	"errors"
)

func parser(t []ltoken) (ast, error) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Uncatch parser error", r)
        }
    }()
	ret := *new(ast)
	if len(t) == 0 {
		return *new(ast), errors.New("Invalid expression")
	}
	ret.parameters = make([]ast,0)
	if t[0].kind=="lparen" && t[len(t)-1].kind=="rparen" {
		t = t[1:len(t)-1]
	} else {
		return *new(ast), errors.New("Invalido")
	}
	//O primeiro valor e o nome da operacao, tem que ser
	ret.status = false
	ret.operation = t[0].value
	ret.kind = t[0].kind
	ret.value = nil
	t = t[1:]
	for len(t) != 0 {
		if t[0].kind == "lparen" {
			j := 1
			for pn := 1; pn!=0; j++ {
				if j > len(t) {
				err := errors.New("Mismatch par!")
				return *new(ast), err
				}
				if t[j].kind == "lparen" {
					pn++
				}
				if t[j].kind == "rparen" {
					pn--
				}
			}
			tmp, _ := parser(t[0:j])
			t = t[j:]
			ret.parameters = append(ret.parameters,tmp)
		} else {
			tmp := *new(ast)
			tmp.status = true
			tmp.value = t[0].value
			tmp.kind = t[0].kind
			t = t[1:]
			ret.parameters = append(ret.parameters,tmp)
		}
	}
	return ret, nil
}
