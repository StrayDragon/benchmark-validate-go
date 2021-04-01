package validator

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	benchmark_validate_go "github.com/straydragon/benchmark-validate-go"
	"testing"
)

type NestedReq struct {
	FicoScore string `json:"ficoScore" validate:"required"`
	Address   struct {
		StreetAddress string `json:"streetAddress" validate:"required"`
		City          string `json:"city" validate:"required"`
		PostalCode    int    `json:"postalCode" validate:"min=10000,max=20000"`
		State         string `json:"state" validate:"required"`
	} `json:"address" validate:"required"`
	Name           string                 `json:"name" validate:"required"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
	// 坑! bool用required要注意: https://github.com/go-playground/validator/issues/319
	Remote       bool     `json:"remote"`
	PhoneNumbers []string `json:"phoneNumbers" validate:"required"`
	Height       float64  `json:"height" validate:"required"`
}

type NestedReqBench struct{}

func (n NestedReqBench) GetJSONFileName(caseName string) string {
	return benchmark_validate_go.GetJSONFileName("nested", caseName)
}

func (n NestedReqBench) ReqModelValidation(caseName string) error {
	req := NestedReq{}
	data := benchmark_validate_go.ReadJSONFile(n.GetJSONFileName(caseName))
	if err := json.Unmarshal(data, &req); err != nil {
		// log.Println("Unmarshal Json Error", err)
		return err
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		// from here you can create your own error messages in whatever language you wish
		return err
	}
	return nil
}

func BenchmarkGoPlayGroundValidatorNestedReqRun(b *testing.B) {
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
