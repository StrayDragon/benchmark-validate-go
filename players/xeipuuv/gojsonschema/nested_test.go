package gojsonschema

import (
	"encoding/json"
	"errors"
	benchmark_validate_go "github.com/straydragon/benchmark-validate-go"
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

type NestedReq struct {
	FicoScore string `json:"ficoScore"`
	Address   struct {
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		PostalCode    int    `json:"postalCode"`
		State         string `json:"state"`
	} `json:"address"`
	Name           string      `json:"name"`
	AdditionalInfo interface{} `json:"additionalInfo"`
	Remote         bool        `json:"remote"`
	PhoneNumbers   []string    `json:"phoneNumbers"`
	Height         float64     `json:"height"`
}

type NestedReqBench struct{}

func (n NestedReqBench) GetJSONFileName(caseName string) string {
	return benchmark_validate_go.GetJSONFileName("nested", caseName)
}
func (n NestedReqBench) GetJSONSchemaFileName() string {
	return "nested.schema.json"
}

func (n NestedReqBench) ReqModelValidation(caseName string) error {
	queryReq := QueryReq{}
	data := benchmark_validate_go.ReadJSONFile(n.GetJSONFileName(caseName))
	if err := json.Unmarshal(data, &queryReq); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	schemaLoader := gojsonschema.NewReferenceLoader(benchmark_validate_go.GetTargetJSONFileStandardPath(n.GetJSONSchemaFileName()))
	documentLoader := gojsonschema.NewReferenceLoader(benchmark_validate_go.GetTargetJSONFileStandardPath(n.GetJSONFileName(caseName)))
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if !result.Valid() {
		return errors.New(result.Errors()[0].String())
	}
	return nil
}

func BenchmarkXeipuuvGojsonschemaNestedReqRun(b *testing.B) {
	benchmarks := []struct {
		name     string
		caseName string
		debug    bool
	}{
		{"OK", "ok", false},
		{"Err1", "err1", false},
		//{"ErrAll", "errall"},
	}
	theme := NestedReqBench{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := theme.ReqModelValidation(bm.caseName)
				if bm.debug && err != nil {
					b.Log(err)
				}
			}
		})
	}
}
