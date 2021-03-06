package jengen

import (
	"bytes"
	"fmt"
	"go/format"

	"github.com/aloder/tojen/gen"
	"github.com/aloder/tojen/run"
)

var ErrGeneratedCodeNotMatch = fmt.Errorf("generated code differs from source")

func Convert(goCode []byte) ([]byte, error) {
	file := gen.GenerateFile(goCode, "main", true)
	resultB := &bytes.Buffer{}
	err := file.Render(resultB)
	if err != nil {
		return nil, err
	}

	jenCode := resultB.Bytes()
	return JenFormat(jenCode)
}

func Validate(goCode, jenCode []byte) ([]byte, error) {
	goCode, err := format.Source(goCode)
	if err != nil {
		return nil, err
	}

	ret, err := run.Exec(string(jenCode))
	if err != nil {
		return nil, err
	}
	fmtBytes, err := format.Source([]byte(*ret))
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(goCode, fmtBytes) {
		return fmtBytes, ErrGeneratedCodeNotMatch
	}
	return fmtBytes, nil
}
