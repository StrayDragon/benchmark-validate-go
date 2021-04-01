package benchmark_validate_go

type BenchmarkPlayer interface {
	GetJSONFileName(caseName string) string
	ReqModelValidation(caseName string) error
}
