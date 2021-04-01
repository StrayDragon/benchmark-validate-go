package jio

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/faceair/jio"
	BenchmarkPopularInputModelValidate "github.com/straydragon/benchmark-validate-go"
	"testing"
)

func getJSONFileName(caseName string) string {
	return fmt.Sprintf("query.%s.json", caseName)
}

type QueryReq struct {
	Origin  string `json:"origin"`
	Limit   string `json:"limit"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func QueryReqValidation(caseName string) error {
	queryReq := QueryReq{}
	data := BenchmarkPopularInputModelValidate.ReadJSONFile(getJSONFileName(caseName))
	if err := json.Unmarshal(data, &queryReq); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	_, err := jio.ValidateJSON(&data, jio.Object().Keys(
		jio.K{
			"origin": jio.String().Required(),
			"limit":  jio.String().Required(),
			"name": jio.String().Transform(func(ctx *jio.Context) {
				allNames := []string{
					"a",
					"b",
					"xxxx",
				}
				for _, name := range allNames {
					if ctx.Value == name {
						ctx.Skip()
					}
				}
				ctx.Abort(errors.New("name不在范围内"))
			}),
			"version": jio.String().Required(),
		}))
	if err != nil {
		return errors.New("校验错误")
	}
	return nil
}

func BenchmarkJioQueryReqRun(b *testing.B) {
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
