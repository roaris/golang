//とほほのGo言語入門 https://www.tohoho-web.com/ex/golang.html
//実行 go run main.go
//コンパイル go build main.go → 実行ファイルmainができる
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("hello, world")
	num := 123
	str := "ABC"
	fmt.Print("num=", num, " str=", str, "\n") //改行無し、空白無し、フォーマット無し
	fmt.Println("num=", num, " str=", str)     //改行有り、空白有り、フォーマット無し
	fmt.Printf("num=%d str=%s\n", num, str)    //改行無し、空白無し、フォーマット有り

	//型
	//int8/int16/int32/int64 nビット整数
	//uint8/uint16/uint32/uint64 nビット非負整数
	//float32/float64 nビット浮動小数点数
	//byte 1バイトデータ(uint8)
	var x uint16 = 1234
	var y uint32 = uint32(x) //型変換
	fmt.Println(y)

	//配列(個数が変更不可)
	color1 := [3]string{}
	color1[0] = "Red"
	color1[1] = "Green"
	color1[2] = "Blue"
	fmt.Println(color1)
	//初期化時に値を設定することもできる
	//color1 := [3]string{"Red", "Green", "Blue"}

	//スライス(個数が変更可能)
	color2 := []string{}
	color2 = append(color2, "Red")
	color2 = append(color2, "Green")
	color2 = append(color2, "Blue")
	fmt.Println(color2)

	//cap: メモリ上に確保されている数 len: 実際に使用されている数
	l := []int{}
	for i := 0; i < 10; i++ {
		l = append(l, i)
		fmt.Println(cap(l), len(l))
	}

	//スライスのメモリ確保にはmake(スライス型, 初期個数, 初期容量)を使うことができる
	//容量超過時の再確保を減らして速度を速めることができる
	//l = make(int[], 0, 1024)

	//配列、スライスのループ処理
	for i, color := range color1 {
		fmt.Printf("%d: %s\n", i, color)
	}

	//マップ
	m := map[string]int{
		"x": 100,
		"y": 200,
	}
	fmt.Println(m["x"]) //参照
	m["z"] = 300        //追加
	fmt.Println(m)
	delete(m, "z") //削除
	fmt.Println(m)
	fmt.Println(len(m)) //長さ

	//要素が存在するかのチェック
	_, ok := m["z"]
	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not Exist")
	}

	//ループ処理
	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}

	//関数呼び出し
	fmt.Println(add(5, 3))
	fmt.Println(addMinus(5, 3))

	//構造体
	var p1 Person
	p1.setPerson("Yamada", 26)
	fmt.Println(p1.getPerson())
	//初期化
	p2 := Person{"Tanaka", 32}
	fmt.Println(p2.getPerson())

	//インタフェース
	b := Book{"吾輩は猫である"}
	printOut1(p1)
	printOut1(b)
	printOut2(p2)
	printOut2("hoge")
	//interface{}を使うと、任意の型を持つ、配列、リスト、マップを作ることができる
	//l := []interface{}{123, "hoge"}

	//ポインタ
	var value int
	var pointer *int //ポインタ変数の定義
	pointer = &value
	*pointer = 123
	fmt.Println(value)
	//値渡しと参照渡し
	n1, n2 := 123, 123
	change(n1, &n2)
	fmt.Println(n1, n2)

	//ゴルーチン 並行処理を実現する
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Println("M")
		time.Sleep(20 * time.Millisecond)
	}
}

//関数定義
func add(x int, y int) int {
	return x + y
}

//複数の値を返却することもできる
func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

//構造体 メンバ変数のみを定義
type Person struct {
	name string
	age  int
}

//クラスメソッド 関数の前に、(thisに相当する変数 *構造体名) をつける
func (p *Person) setPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) getPerson() (string, int) {
	return p.name, p.age
}

//インタフェース
func (p Person) toString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) toString() string {
	return b.title
}

type Printable interface {
	toString() string
}

func printOut1(p Printable) {
	fmt.Println(p.toString())
}

func printOut2(p interface{}) { //どんな型でも受け取る
	q, ok := p.(Printable) //.(型名)でinterfact{}型を他の型に変換
	if ok {
		fmt.Println(q.toString())
	} else {
		fmt.Println("Not printable.")
	}
}

//値渡しと参照渡し
func change(x int, y *int) {
	x = 456
	*y = 456
}

//遅延実行 関数から戻る直前に処理を実行 リソースの解放によく用いられる
func fileOpen() {
	fp, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	defer fp.Close()
}

//ゴルーチン
func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Println("A")
		time.Sleep(10 * time.Millisecond)
	}
}
