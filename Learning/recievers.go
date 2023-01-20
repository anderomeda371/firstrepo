package main

import "fmt"

type book struct {
	name   string
	author string
	year   int32
}

func (b book) createString() string {
	str := b.name + " , " + b.author + " , " + fmt.Sprint(b.year)
	return str
}
func main() {

	b1 := book{"20 hezar farsang zire darya", "jul vern", 1890}

	println(b1.createString())
}
