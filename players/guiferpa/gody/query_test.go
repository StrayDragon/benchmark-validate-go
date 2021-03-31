package gody

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guiferpa/gody/v2"
	"github.com/guiferpa/gody/v2/rule"
	BenchmarkPopularInputModelValidate "github.com/straydragon/benchmark-validate-go"
	"testing"
)

func getJSONFileName(caseName string) string {
	return fmt.Sprintf("query.%s.json", caseName)
}

type QueryReq struct {
	Origin  string `json:"origin" validate:"not_empty"`
	Limit   string `json:"limit" validate:"not_empty"`
	Name    string `json:"name" validate:"name_valid"`
	Version string `json:"version" validate:"not_empty"`
}

type QueryReqNameValidation struct {
}

func (QueryReqNameValidation) Name() string {
	return "name_valid"
}

func (q *QueryReqNameValidation) Validate(_, v, _ string) (bool, error) {
	allNames := []string{
		"a",
		"b",
		"xxxx",
	}
	for _, name := range allNames {
		if v == name {
			return true, nil
		}
	}
	return false, errors.New("name不在范围内")
}

func QueryReqValidation(caseName string) error {
	queryReq := QueryReq{}
	data := BenchmarkPopularInputModelValidate.ReadJSONFile(getJSONFileName(caseName))
	if err := json.Unmarshal(data, &queryReq); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	validator := gody.NewValidator()
	_ = validator.AddRules(rule.NotEmpty, &QueryReqNameValidation{})
	if _, err := validator.Validate(queryReq); err != nil {
		// log.Println("body do not validated", err)
		return errors.New("校验错误")
	}
	return nil
}

func BenchmarkGodyQueryReqRun(b *testing.B) {
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
