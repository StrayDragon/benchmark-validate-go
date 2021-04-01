package jio

import (
	"encoding/json"
	"github.com/faceair/jio"
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
	} `json:"address"`
	Name           string                 `json:"name"`
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
	_, err := jio.ValidateJSON(
		&data, jio.Object().Keys(
			jio.K{
				"name": jio.String().Required(),
				"address": jio.Object().Keys(jio.K{
					"streetAddress": jio.String().Required(),
					"city":          jio.String().Required(),
					"state":         jio.String().Required(),
					"postalCode":    jio.Number().Integer().Min(10000).Max(20000).Required(),
				}),
				"phoneNumbers":   jio.Array().Items(jio.String().Required()),
				"additionalInfo": jio.Object().Optional(),
				"remote":         jio.Bool().Required(),
				"height":         jio.Number().Required(),
				"ficoScore":      jio.String().Required(),
			}),
	)
	if err != nil {
		return err
	}
	return nil
}

func BenchmarkJioNestedReqRun(b *testing.B) {
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
