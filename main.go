package main

import (
	"fmt"
)

// ***型スイッチ***

// interface型は型演算できない、演算するには型を復元する必要があります
func anything(a interface{}) {
	// 型スイッチ
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

// nilかどうかを判断する関数 (https://zenn.dev/jboy_blog/articles/b47ccfa0629206)
func isNil(x interface{}) bool {
	// reflect.ValueOf(x).IsNil()は、xがnilかどうかを判定する
	// return x == nil || reflect.ValueOf(x).IsNil()
	return x == nil
}

func main() {
	anything("aaa")

	// ***型アサーション(型を上書き、TSのas相当)***
	var x interface{} = 3
	// interfaceをint型に復元
	i := x.(int)
	fmt.Println(i + 2)

	// int型はfloat64型で復元できないので、runtimeエラーになります
	// f := x.(float64)
	// fmt.Println(f)

	// runtimeエラーにしない方法
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	// isNilの関数を通じると警告は消えます
	/*
		if isNil(x) {
			fmt.Println("None")
		} else if i, isInt := x.(int); isInt {
			fmt.Println(i, "x is Int")
		} else if s, isString := x.(string); isString {
			fmt.Println(s, isString)
		} else {
			fmt.Println("I don't Know")
		}
	*/

	// ***ここ警告が出ました***
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't Know")
	}

	// 型によるswitch文
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
