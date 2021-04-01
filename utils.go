package benchmark_validate_go

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _     = runtime.Caller(0)
	basePath       = filepath.Dir(b)
	modelsJSONPath = path.Join(basePath, "models")
)

func BasePath() string {
	return basePath
}

func ReadJSONFile(name string) []byte {
	targetPath := GetTargetJSONFilePath(name)
	bytesValue, err := ioutil.ReadFile(targetPath)
	if err != nil {
		panic(err.Error())
	}
	return bytesValue
}

func GetTargetJSONFilePath(name string) string {
	if filepath.Ext(name) != ".json" {
		panic("Cannot read non-json file!")
	}
	targetPath := path.Join(modelsJSONPath, name)
	return targetPath
}

func GetTargetJSONFileStandardPath(name string) string {
	if filepath.Ext(name) != ".json" {
		panic("Cannot read non-json file!")
	}
	targetPath := path.Join(modelsJSONPath, name)
	return "file://" + targetPath
}

func GetJSONFileName(exampleName, caseName string) string {
	return fmt.Sprintf("%s.%s.json", exampleName, caseName)
}
