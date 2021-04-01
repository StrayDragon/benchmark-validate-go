package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	v "github.com/gobuffalo/validate"
	benchmark_validate_go "github.com/straydragon/benchmark-validate-go"
	"testing"
)

type NestedReq struct {
	FicoScore string `json:"ficoScore"`
	Address   struct {
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		PostalCode    int    `json:"postalCode"`
		State         string `json:"state"`
	} `json:"address"` // https://github.com/go-ozzo/ozzo-validation/issues/136
	Name           string                 `json:"name"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
	Remote         bool                   `json:"remote"`
	PhoneNumbers   []string               `json:"phoneNumbers"`
	Height         float64                `json:"height"`
}

func (n *NestedReq) IsValid(errors *v.Errors) {
	// PostalCode 范围限制
	if n.Address.PostalCode < 10000 || n.Address.PostalCode > 20000 {
		errors.Add("address?.postalCode", "PostalCode 必须范围在[10000,20000]中")
	}
	// AdditionalInfo Optional
	// 其他字段 Required
	targets := []interface{}{
		n.FicoScore, n.Address, n.Name, n.Remote, n.PhoneNumbers, n.Height,
		n.Address.StreetAddress, n.Address.City, n.Address.State,
	}
	msg := "该字段不可为空"
	for idx, t := range targets {
		switch t := t.(type) {
		case string:
			if t == "" {
				key := fmt.Sprintf("#{%d}", idx)
				errors.Add(key, msg)
			}
		case []string:
			if len(t) == 0 {
				key := fmt.Sprintf("#{%d}", idx)
				errors.Add(key, msg)
			}
		case float64:
			if t == 0 {
				key := fmt.Sprintf("#{%d}", idx)
				errors.Add(key, msg)
			}
		}
	}
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
	validate := v.Validate(&req)
	if len(validate.Errors) != 0 {
		return errors.New(validate.Error())
	}
	return nil
}

func BenchmarkGoBuffaioValidateNestedReqRun(b *testing.B) {
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
