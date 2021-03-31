package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	BenchmarkPopularInputModelValidate "github.com/straydragon/benchmark-validate-go"
	"strings"
	"testing"
)
import v "github.com/gobuffalo/validate"

func getJSONFileName(caseName string) string {
	return fmt.Sprintf("query.%s.json", caseName)
}

type QueryReq struct {
	Origin  string `json:"origin"`
	Limit   string `json:"limit"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
type NotEmptyValidator struct {
	Field string
	Value string
}

func (v *NotEmptyValidator) IsValid(errors *v.Errors) {
	if v.Value == "" {
		errors.Add(strings.ToLower(v.Field), fmt.Sprintf("%s must not be empty!", v.Field))
	}
}

type NameValidator struct {
	Field string
	Value string
}

func (n *NameValidator) IsValid(errors *v.Errors) {
	allNames := []string{
		"a",
		"b",
		"xxxx",
	}
	ok := false
	for _, name := range allNames {
		if n.Value == name {
			ok = true
			break
		}
	}
	if !ok {
		errors.Add(strings.ToLower(n.Field), "name 不在范围内")
	}
}

func QueryReqValidation(caseName string) error {
	queryReq := QueryReq{}
	data := BenchmarkPopularInputModelValidate.ReadJSONFile(getJSONFileName(caseName))
	if err := json.Unmarshal(data, &queryReq); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	if err := v.Validate(
		&NotEmptyValidator{"Origin", queryReq.Origin},
		&NotEmptyValidator{"Limit", queryReq.Limit},
		&NameValidator{"Name", queryReq.Name},
		&NotEmptyValidator{"Version", queryReq.Version},
	); err != nil {
		//log.Println("body do not validated", err)
		return errors.New("校验错误")
	}
	return nil
}

func BenchmarkGoBuffaioValidateQueryReqRun(b *testing.B) {
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
