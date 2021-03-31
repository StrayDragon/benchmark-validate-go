package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/validate"
	BenchmarkPopularInputModelValidate "github.com/straydragon/benchmark-validate-go"
	"testing"
)

func getJSONFileName(caseName string) string {
	return fmt.Sprintf("query.%s.json", caseName)
}

type QueryReq struct {
	Origin  string `json:"origin" validate:"required"`
	Limit   string `json:"limit" validate:"required"`
	Name    string `json:"name" validate:"nameValid" message:"name不在范围内"`
	Version string `json:"version" validate:"required"`
}

func (r QueryReq) NameValid(v string) bool {
	allNames := []string{
		"a",
		"b",
		"xxxx",
	}
	for _, name := range allNames {
		if v == name {
			return true
		}
	}
	return false
}

func QueryReqValidation(caseName string) error {
	queryReq := QueryReq{}
	data := BenchmarkPopularInputModelValidate.ReadJSONFile(getJSONFileName(caseName))
	if err := json.Unmarshal(data, &queryReq); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	v := validate.Struct(queryReq)
	if !v.Validate() {
		return errors.New("校验错误")
	}
	return nil
}

func BenchmarkGookitValidateQueryReqRun(b *testing.B) {
	benchmarks := []struct {
		name     string
		caseName string
	}{
		{"OK", "ok"},
		{"Err1", "err1"},
		{"ErrAll", "errall"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = QueryReqValidation(bm.caseName)
			}
		})
	}
}
