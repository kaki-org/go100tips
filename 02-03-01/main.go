package main

import "fmt"

/*
次の順序で処理が実行される
  1. パッケージレベル変数の初期化 - var a = func()...
  2. init関数の実行 - func init()
  3. main関数の実行 - func main()
*/

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
