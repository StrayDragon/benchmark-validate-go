package ozzo_validation

import (
	"encoding/json"
	"errors"
	"fmt"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	BenchmarkPopularInputModelValidate "github.com/straydragon/benchmark-validate-go"
	"testing"
)

func getJSONFileName(caseName string) string {
	return fmt.Sprintf("query.%s.json", caseName)
}

type QueryReq struct {
	Origin  string
	Limit   string
	Name    string
	Version string
}

func (q QueryReq) Validate() error {
	return ozzo.ValidateStruct(
		&q,
		ozzo.Field(&q.Origin, ozzo.NilOrNotEmpty),
		ozzo.Field(&q.Limit, ozzo.NilOrNotEmpty),
		ozzo.Field(&q.Name, ozzo.By(func(v interface{}) error {
			s, ok := v.(string)
			if !ok {
				return errors.New("type conv error!!!")
			}
			allNames := []string{
				"a",
				"b",
				"xxxx",
			}
			for _, name := range allNames {
				if s == name {
					return nil
				}
			}
			return errors.New("name不在范围内")
		})),
		ozzo.Field(&q.Version, ozzo.NilOrNotEmpty),
	)
}

func QueryReqValidation(caseName string) error {
	queryReq := QueryReq{}
	data := BenchmarkPopularInputModelValidate.ReadJSONFile(getJSONFileName(caseName))
	if err := json.Unmarshal(data, &queryReq); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	if err := queryReq.Validate(); err != nil {
		//log.Println("body do not validated", err)
		return errors.New("校验错误")
	}
	return nil
}

func BenchmarkOzzoQueryReqRun(b *testing.B) {
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
