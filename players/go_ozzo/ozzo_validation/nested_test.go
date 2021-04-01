package ozzo_validation

import (
	"encoding/json"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	benchmark_validate_go "github.com/straydragon/benchmark-validate-go"
	"testing"
)

type NestedReqBench struct{}

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

func (r NestedReq) Validate() error {
	return ozzo.ValidateStruct(&r,
		ozzo.Field(&r.Name, ozzo.Required),
		ozzo.Field(&r.Address.StreetAddress, ozzo.Required),
		ozzo.Field(&r.Address.City, ozzo.Required),
		ozzo.Field(&r.Address.PostalCode, ozzo.Required, ozzo.Min(1000), ozzo.Max(2000)),
		ozzo.Field(&r.Address.State, ozzo.Required),
		ozzo.Field(&r.Address, ozzo.Required),
		ozzo.Field(&r.PhoneNumbers, ozzo.Required),
		ozzo.Field(&r.AdditionalInfo, ozzo.NilOrNotEmpty),
		ozzo.Field(&r.Remote, ozzo.Required),
		ozzo.Field(&r.Height, ozzo.Required),
		ozzo.Field(&r.FicoScore, ozzo.Required),
	)
}

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
	if err := req.Validate(); err != nil {
		//log.Println("body do not validated", err)
		return err
		//return errors.New("校验错误")
	}
	return nil
}

func BenchmarkOzzoNestedReqRun(b *testing.B) {
	b.Skipf("测试嵌套struct无法作校验 参考: https://github.com/go-ozzo/ozzo-validation/issues/136")
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
