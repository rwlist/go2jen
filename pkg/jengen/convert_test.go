package jengen

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	type testcase struct {
		GoCode string
		Jen    string
	}

	cases := []testcase{
		{
			GoCode: `package a

func hello(arg string) {
	fmt.Println("Hello", arg)		
}`,
			Jen: `package main

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
}
`,
		},
	}

	for _, test := range cases {
		fmt.Println("BEFORE =====>")
		fmt.Println(test.GoCode)

		res, err := Convert([]byte(test.GoCode))
		assert.NoError(t, err)
		fmt.Println("AFTER ======>")
		fmt.Println(string(res))

		assert.Equal(t, test.Jen, string(res))
	}
}

func TestXXX(t *testing.T) {
	c := jen.Func().Id("hello").Params().Block()

	fmt.Printf("%#v\n", c)
}
