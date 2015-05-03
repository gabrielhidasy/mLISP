package main

import (
	"fmt"
	"errors"
)

func eval(AST ast) (interface{},error) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Uncatch eval error", r)
        }
    }()
	if AST.status == true {
		return AST.value, nil
	}
	if AST.operation == "+" {
		var ret int64
		for i := 0; i < len(AST.parameters); i++ {
			tmp, err := eval(AST.parameters[i])
			if err != nil {
				return 0, err
			}
			ret = ret + tmp.(int64)
		}
		return ret, nil
	}
	if AST.operation == "-" {
		var ret int64
		tmp, err := eval(AST.parameters[0])
		if err != nil {
			return 0, err
		}
		ret = tmp.(int64)
		for i := 1; i < len(AST.parameters); i++ {
			tmp, err := eval(AST.parameters[i])
			if err != nil {
				return 0, err
			}
			ret = ret - tmp.(int64)
		}
		return ret, nil
	}
	if AST.operation == "*" {
		var ret int64
		ret = 1
		for i := 0; i < len(AST.parameters); i++ {
			tmp, err := eval(AST.parameters[i])
			if err != nil {
				return 0, err
			}
			ret = ret * tmp.(int64)
		}
		return ret, nil
	}
	if AST.operation == "/" {
		var ret int64
		tmp, err := eval(AST.parameters[0])
		if err != nil {
			return 0, err
		}
		ret = tmp.(int64)
		for i := 1; i < len(AST.parameters); i++ {
			tmp, err := eval(AST.parameters[i])
			if err != nil {
				return 0, err
			}
			ret = ret / tmp.(int64)
		}
		return ret, nil
	}
	return 0, errors.New("Undefined operation")
}
