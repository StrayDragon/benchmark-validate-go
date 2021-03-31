# Benchmark 一下热门的validator

Date Created: Jan 7, 2021 1:15 PM Property: Mar 30, 2021 2:18 PM Status: 进行中

# 待测面板

|库名|来源|校验方式|优点|缺点|特性::自定义错误提示|特性::转换字段类型或值|
|:----|:----|:----|:----|:----|:----|:----|
|ozzo-validation|https://github.com/go-ozzo/ozzo-validation|增加ViewModel的特定方法校验|类型自动补全, 易于重构; 预置方法(规则)丰富|使用时最好重命名库名字, 要不太冗余; 好久未更新了|Yes|No|
|validate|https://github.com/gobuffalo/validate|增加ViewModel的特定方法校验|类型自动补全, 易于重构|一年未更新了|Yes|No|
|validate|https://github.com/gookit/validate|structure tag的方式校验+自定义validate loc|内置规则丰富; 更新频繁|比较复杂; 重构蛋疼|Yes|No|
|validator|https://github.com/go-playground/validator|structure tag的方式校验+自定义validate loc|内置规则丰富; |重构蛋疼|Yes|No|
|gody|https://github.com/guiferpa/gody|structure tag的方式校验+自定义validate loc|接口简单, 方便学习; |比较原始|Yes|No|
|jio|https://github.com/faceair/jio|对序列化前的JSON字符串作校验|链式调用方便写; 可以作转换|两年未更新了; 重构蛋疼|No|No|
|gojsonschema|https://github.com/xeipuuv/gojsonschema|对序列化前的JSON字符串作校验|接受schema作校验, 使用方便|好久没更新了; 重构中等难度|No|No|


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
