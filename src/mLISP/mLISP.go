package main


//Isso é um LISP, um LISP é composto de espressões. uma espressão
//pode ser colocada em forma de uma árvore a ser processada (quem sabe
//rola até eu fazer lazy evaluation aqui)

type ast struct {
	status bool //False = precisa eval, True = já eval
	kind string
	operation interface{}
	parameters  []ast
	value interface{}
}

type ltoken struct {
	kind string
	value interface{}
}


func parse_and_lex(s string) (ast,error) {
	tokens, err := lexer(s)
	//fmt.Println(tokens)
	AST, err := parser(tokens)
	return AST,err
}

