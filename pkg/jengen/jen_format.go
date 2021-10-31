package jengen

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type walker struct {
	prev dst.Node
}

func (w *walker) shouldIndent(expr *dst.CallExpr) bool {
	if len(expr.Args) > 1 {
		return true
	}

	if sel, ok := expr.Fun.(*dst.SelectorExpr); ok {
		if sel.Sel.Name == "Block" {
			return true
		}
	}

	return false
}

func (w *walker) Visit(node dst.Node) dst.Visitor {
	if node == nil {
		return w
	}

	// Use `fmt.Printf("%T/%+v\n", node, node)` to print debug info

	switch n := node.(type) {
	case *dst.CallExpr:
		if w.shouldIndent(n) {
			w.indentArgs(n.Args)
		}
	}

	w.prev = node
	return w
}

func (w *walker) indentArgs(args []dst.Expr) {
	for _, arg := range args {
		arg.Decorations().Before = dst.NewLine
		arg.Decorations().After = dst.NewLine
	}
}

func JenFormat(src []byte) ([]byte, error) {
	f, err := decorator.Parse(src)
	if err != nil {
		return nil, err
	}

	w := &walker{}
	dst.Walk(w, f)

	buf := bytes.NewBuffer(nil)
	err = decorator.Fprint(buf, f)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
