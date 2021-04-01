# Benchmark 一下热门的validator

Date Created: Jan 7, 2021 1:15 PM Property: Mar 30, 2021 2:18 PM Status: 进行中

# 待测面板

|库名|来源|校验方式|优点|缺点|
|:----|:----|:----|:----|:----|
|ozzo-validation|https://github.com/go-ozzo/ozzo-validation|增加ViewModel的特定方法校验|类型自动补全, 易于重构; 预置方法(规则)丰富|使用时最好重命名库名字, 要不太冗余; 好久未更新了
|validate|https://github.com/gobuffalo/validate|增加ViewModel的特定方法校验|类型自动补全, 易于重构|一年未更新了
|validate|https://github.com/gookit/validate|structure tag的方式校验+自定义validate loc|内置规则丰富; 更新频繁|比较复杂; 重构蛋疼
|validator|https://github.com/go-playground/validator|structure tag的方式校验+自定义validate loc|内置规则丰富; |重构蛋疼
|gody|https://github.com/guiferpa/gody|structure tag的方式校验+自定义validate loc|接口简单, 方便学习; |比较原始
|jio|https://github.com/faceair/jio|对序列化前的JSON字符串作校验|链式调用方便写; 可以作转换|两年未更新了; 重构蛋疼
|gojsonschema|https://github.com/xeipuuv/gojsonschema|对序列化前的JSON字符串作校验|接受schema作校验, 使用方便|好久没更新了; 重构中等难度

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

> go test -v -run=none -short -benchmem -bench=. ./players/...

#### QueryReq

```bash
BenchmarkJioQueryReqRun
BenchmarkJioQueryReqRun/OK
BenchmarkJioQueryReqRun/OK-8               48537             23831 ns/op            3605 B/op         78 allocs/op
BenchmarkJioQueryReqRun/Err1
BenchmarkJioQueryReqRun/Err1-8             49759             24026 ns/op            3621 B/op         78 allocs/op
BenchmarkJioQueryReqRun/ErrAll
BenchmarkJioQueryReqRun/ErrAll-8           50278             23995 ns/op            3565 B/op         71 allocs/op

BenchmarkOzzoQueryReqRun
BenchmarkOzzoQueryReqRun/OK
BenchmarkOzzoQueryReqRun/OK-8      56701             20999 ns/op            2977 B/op         49 allocs/op
BenchmarkOzzoQueryReqRun/Err1
BenchmarkOzzoQueryReqRun/Err1-8                    52014             21446 ns/op            3305 B/op         52 allocs/op
BenchmarkOzzoQueryReqRun/ErrAll
BenchmarkOzzoQueryReqRun/ErrAll-8                  56056             21275 ns/op            3305 B/op         50 allocs/op

BenchmarkGoPlayGroundValidatorQueryReqRun
BenchmarkGoPlayGroundValidatorQueryReqRun/OK
BenchmarkGoPlayGroundValidatorQueryReqRun/OK-8                     28544             41374 ns/op           16490 B/op        188 allocs/op
BenchmarkGoPlayGroundValidatorQueryReqRun/Err1
BenchmarkGoPlayGroundValidatorQueryReqRun/Err1-8                   28656             41856 ns/op           16713 B/op        193 allocs/op
BenchmarkGoPlayGroundValidatorQueryReqRun/ErrAll
BenchmarkGoPlayGroundValidatorQueryReqRun/ErrAll-8                 28036             42862 ns/op           17294 B/op        199 allocs/op

BenchmarkGoBuffaioValidateQueryReqRun
BenchmarkGoBuffaioValidateQueryReqRun/OK
BenchmarkGoBuffaioValidateQueryReqRun/OK-8                 53008             22707 ns/op            1744 B/op         25 allocs/op
BenchmarkGoBuffaioValidateQueryReqRun/Err1
BenchmarkGoBuffaioValidateQueryReqRun/Err1-8               51734             23302 ns/op            2121 B/op         28 allocs/op
BenchmarkGoBuffaioValidateQueryReqRun/ErrAll
BenchmarkGoBuffaioValidateQueryReqRun/ErrAll-8             45303             26335 ns/op            2344 B/op         38 allocs/op

BenchmarkGookitValidateQueryReqRun
BenchmarkGookitValidateQueryReqRun/OK
BenchmarkGookitValidateQueryReqRun/OK-8                    24292             49747 ns/op           22248 B/op        143 allocs/op
BenchmarkGookitValidateQueryReqRun/Err1
BenchmarkGookitValidateQueryReqRun/Err1-8                  24139             49942 ns/op           22851 B/op        147 allocs/op
BenchmarkGookitValidateQueryReqRun/ErrAll
BenchmarkGookitValidateQueryReqRun/ErrAll-8                25896             45801 ns/op           21919 B/op        133 allocs/op

BenchmarkGodyQueryReqRun
BenchmarkGodyQueryReqRun/OK
BenchmarkGodyQueryReqRun/OK-8      51745             22535 ns/op            4073 B/op         58 allocs/op
BenchmarkGodyQueryReqRun/Err1
BenchmarkGodyQueryReqRun/Err1-8                    52340             23563 ns/op            4129 B/op         60 allocs/op
BenchmarkGodyQueryReqRun/ErrAll
BenchmarkGodyQueryReqRun/ErrAll-8                  48054             23286 ns/op            4113 B/op         56 allocs/op

BenchmarkXeipuuvGojsonschemaQueryReqRun
BenchmarkXeipuuvGojsonschemaQueryReqRun/OK
BenchmarkXeipuuvGojsonschemaQueryReqRun/OK-8                9535            117713 ns/op           32676 B/op        414 allocs/op
BenchmarkXeipuuvGojsonschemaQueryReqRun/Err1
BenchmarkXeipuuvGojsonschemaQueryReqRun/Err1-8             10000            101142 ns/op           32155 B/op        420 allocs/op
BenchmarkXeipuuvGojsonschemaQueryReqRun/ErrAll
BenchmarkXeipuuvGojsonschemaQueryReqRun/ErrAll-8           10000            106280 ns/op           34285 B/op        450 allocs/op
```

### NestedReq

```bash
BenchmarkJioNestedReqRun
BenchmarkJioNestedReqRun/OK
BenchmarkJioNestedReqRun/OK-8              25958             49426 ns/op            9054 B/op        221 allocs/op
BenchmarkJioNestedReqRun/Err1
BenchmarkJioNestedReqRun/Err1-8            28188             40314 ns/op            6966 B/op        182 allocs/op

BenchmarkOzzoNestedReqRun
--- SKIP: BenchmarkOzzoNestedReqRun

BenchmarkGoPlayGroundValidatorNestedReqRun
BenchmarkGoPlayGroundValidatorNestedReqRun/OK
BenchmarkGoPlayGroundValidatorNestedReqRun/OK-8                    21518             54955 ns/op           19243 B/op        247 allocs/op
BenchmarkGoPlayGroundValidatorNestedReqRun/Err1
BenchmarkGoPlayGroundValidatorNestedReqRun/Err1-8                  18774             62964 ns/op           19468 B/op        251 allocs/op

BenchmarkGoBuffaioValidateNestedReqRun
BenchmarkGoBuffaioValidateNestedReqRun/OK
BenchmarkGoBuffaioValidateNestedReqRun/OK-8                43204             27484 ns/op            2120 B/op         31 allocs/op
BenchmarkGoBuffaioValidateNestedReqRun/Err1
BenchmarkGoBuffaioValidateNestedReqRun/Err1-8              38456             36390 ns/op            2520 B/op         35 allocs/op

BenchmarkGookitValidateNestedReqRun
BenchmarkGookitValidateNestedReqRun/OK
BenchmarkGookitValidateNestedReqRun/OK-8                   16126             74133 ns/op           28734 B/op        255 allocs/op
BenchmarkGookitValidateNestedReqRun/Err1
BenchmarkGookitValidateNestedReqRun/Err1-8                 14742             70768 ns/op           29494 B/op        270 allocs/op

BenchmarkGodyNestedReqRun
--- SKIP: BenchmarkGodyNestedReqRun

BenchmarkXeipuuvGojsonschemaNestedReqRun
BenchmarkXeipuuvGojsonschemaNestedReqRun/OK
BenchmarkXeipuuvGojsonschemaNestedReqRun/OK-8               6416            166384 ns/op           57153 B/op        654 allocs/op
BenchmarkXeipuuvGojsonschemaNestedReqRun/Err1
BenchmarkXeipuuvGojsonschemaNestedReqRun/Err1-8             6184            182663 ns/op           59605 B/op        716 allocs/op
```

# 参考文献

- [https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835](https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835)
