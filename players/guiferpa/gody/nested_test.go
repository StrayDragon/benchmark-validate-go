package gody

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/guiferpa/gody/v2"
	benchmark_validate_go "github.com/straydragon/benchmark-validate-go"
	"testing"
)

type NestedReq struct {
	FicoScore string `json:"ficoScore" validate:"not_empty"`
	Address   struct {
		StreetAddress string `json:"streetAddress" validate:"not_empty"`
		City          string `json:"city" validate:"not_empty"`
		PostalCode    int    `json:"postalCode" validate:"min=10000 max=20000"`
		State         string `json:"state" validate:"not_empty"`
	} `json:"address"`
	Name           string                 `json:"name" validate:"not_empty"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
	Remote         bool                   `json:"remote"`
	PhoneNumbers   []string               `json:"phoneNumbers"`
	Height         float64                `json:"height"`
}

type NestedReqBench struct {
}

func (n NestedReqBench) GetJSONFileName(caseName string) string {
	return benchmark_validate_go.GetJSONFileName("nested", caseName)
}

func (n NestedReqBench) ReqModelValidation(caseName string) error {
	req := NestedReq{}
	data := benchmark_validate_go.ReadJSONFile(n.GetJSONFileName(caseName))
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	spew.Dump(req)
	validator := gody.NewValidator()
	if _, err := validator.Validate(req); err != nil {
		return err
	}
	return nil
}

func BenchmarkGodyNestedReqRun(b *testing.B) {
	b.Skipf("测试嵌套struct中string field校验错误")
	benchmarks := []struct {
		name     string
		caseName string
		debug    bool
	}{
		{"OK", "ok", true},
		//{"Err1", "err1", false},
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
