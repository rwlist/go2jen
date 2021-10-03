package jengen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJenFormat(t *testing.T) {
	const src = `package main

import (
	"fmt"
	jen "github.com/dave/jennifer/jen"
)

func genFunchello() jen.Code {
	return jen.Func().Id("hello").Params(jen.Id("arg").Id("string")).Block(jen.Id("fmt").Dot("Println").Call(jen.Lit("Hello"), jen.Id("arg")))
}
func genFile() *jen.File {
	ret := jen.NewFile("a")
	ret.Add(genFunchello())
	return ret
}
func main() {
	ret := genFile()
	fmt.Printf("%#v", ret)
}`

	res, err := JenFormat([]byte(src))
	assert.NoError(t, err)
	fmt.Println(string(res))
}
