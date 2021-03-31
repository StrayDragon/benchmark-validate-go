# Benchmark 一下热门的validator

Date Created: Jan 7, 2021 1:15 PM Property: Mar 30, 2021 2:18 PM Status: 进行中

# 待测面板

[待测列表](https://www.notion.so/4c7c901336434f52a5d556c77c7e7f46)

# 校验时遇到的问题

请求数据 =》校验与转换 =》bloc =转换=》 返回数据

1. 请求的字段基础校验 (json), 格式转换, 请求model 额外携带转换后的信息, 比如加密的字段拆分为解密的多个字段
2. 请求的字段依赖底层数据的校验; 如验证一个id对应的实体是否存在, 不存在会panic
3. Optional/Required校验

# 考察校验库的角度

## 自定义校验方式方不方便
## 性能
测试机配置
```
                    'c.          straydragon@MacBook-Pro.local
                 ,xNMM.          -----------------------------
               .OMMMMo           OS: macOS 11.2.3 20D91 x86_64
               OMMM0,            Host: MacBookPro15,4
     .;loddo:' loolloddol;.      Kernel: 20.3.0
   cKMMMMMMMMMMNWMMMMMMMMMM0:    Uptime: 10 days, 21 hours, 18 mins
 .KMMMMMMMMMMMMMMMMMMMMMMMWd.    Packages: 182 (brew)
 XMMMMMMMMMMMMMMMMMMMMMMMX.      Shell: zsh 5.8
;MMMMMMMMMMMMMMMMMMMMMMMM:       Resolution: 1440x900@2x, 1920x1080@2x
:MMMMMMMMMMMMMMMMMMMMMMMM:       DE: Aqua
.MMMMMMMMMMMMMMMMMMMMMMMMX.      WM: yabai
 kMMMMMMMMMMMMMMMMMMMMMMMMWd.    Terminal: iTerm2
 .XMMMMMMMMMMMMMMMMMMMMMMMMMMk   Terminal Font: JetBrainsMonoNerdFontCompleteM-Regular 14
  .XMMMMMMMMMMMMMMMMMMMMMMMMK.   CPU: Intel i7-8557U (8) @ 1.70GHz
    kMMMMMMMMMMMMMMMMMMMMMMd     GPU: Intel Iris Plus Graphics 645
     ;KMMMMMMMWXXWMMMMMMMk.      Memory: 11176MiB / 16384MiB
       .cooc,.    .,coo:.
```

### 报告
```bash
$ go test -test.bench -run=none -short -bench=. ./players/...
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/faceair/jio
BenchmarkJioQueryReqRun/OK-8               48982             23468 ns/op
BenchmarkJioQueryReqRun/Err1-8             49369             24441 ns/op
BenchmarkJioQueryReqRun/ErrAll-8           49230             23891 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/faceair/jio        4.293s
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/go_ozzo/ozzo_validation
BenchmarkOzzoQueryReqRun/OK-8              56920             20453 ns/op
BenchmarkOzzoQueryReqRun/Err1-8            55339             21600 ns/op
BenchmarkOzzoQueryReqRun/ErrAll-8          55324             21727 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/go_ozzo/ozzo_validation    4.232s
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/go_playground/validator
BenchmarkValidatorQueryReqRun/OK-8                 27964             42899 ns/op
BenchmarkValidatorQueryReqRun/Err1-8               27531             44944 ns/op
BenchmarkValidatorQueryReqRun/ErrAll-8             26418             47649 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/go_playground/validator    5.054s
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/gobuffaio/validate
BenchmarkGoBuffaioValidateQueryReqRun/OK-8                 52370             23785 ns/op
BenchmarkGoBuffaioValidateQueryReqRun/Err1-8               45301             24622 ns/op
BenchmarkGoBuffaioValidateQueryReqRun/ErrAll-8             44413             26969 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/gobuffaio/validate 4.352s
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/gookit/validate
BenchmarkGookitValidateQueryReqRun/OK-8                    23784             49825 ns/op
BenchmarkGookitValidateQueryReqRun/Err1-8                  22837             52455 ns/op
BenchmarkGookitValidateQueryReqRun/ErrAll-8                25610             47966 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/gookit/validate    5.149s
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/guiferpa/gody
BenchmarkGodyQueryReqRun/OK-8              50532             24247 ns/op
BenchmarkGodyQueryReqRun/Err1-8            50606             24020 ns/op
BenchmarkGodyQueryReqRun/ErrAll-8          51660             23547 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/guiferpa/gody      4.391s
goos: darwin
goarch: amd64
pkg: github.com/straydragon/benchmark-validate-go/players/xeipuuv/gojsonschema
BenchmarkXeipuuvGojsonschemaQueryReqRun/OK-8                8688            115771 ns/op
BenchmarkXeipuuvGojsonschemaQueryReqRun/Err1-8             10000            100532 ns/op
BenchmarkXeipuuvGojsonschemaQueryReqRun/ErrAll-8            9590            107785 ns/op
PASS
ok      github.com/straydragon/benchmark-validate-go/players/xeipuuv/gojsonschema       3.103s

```

# 参考文献

- [https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835](https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835)
