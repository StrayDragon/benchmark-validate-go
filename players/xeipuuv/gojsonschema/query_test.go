package gojsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	BenchmarkPopularInputModelValidate "github.com/straydragon/benchmark-validate-go"
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

func getJSONFileName(caseName string) string {
	return fmt.Sprintf("query.%s.json", caseName)
}

func getJSONSchemaFileName(caseName string) string {
	return "query.schema.json"
}

type QueryReq struct {
	Origin  string `json:"origin"`
	Limit   string `json:"limit"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func QueryReqValidation(caseName string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(BenchmarkPopularInputModelValidate.GetTargetJSONFileStandardPath(getJSONSchemaFileName(caseName)))
	documentLoader := gojsonschema.NewReferenceLoader(BenchmarkPopularInputModelValidate.GetTargetJSONFileStandardPath(getJSONFileName(caseName)))
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if result.Valid() {
		queryReq := QueryReq{}
		data := BenchmarkPopularInputModelValidate.ReadJSONFile(getJSONFileName(caseName))
		if err := json.Unmarshal(data, &queryReq); err != nil {
			// log.Println("Unmarshal Json Error", err)
			return err
		}
	} else {
		return errors.New("校验错误")
	}
	return nil
}

func BenchmarkXeipuuvGojsonschemaQueryReqRun(b *testing.B) {
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
