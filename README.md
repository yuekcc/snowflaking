Snow Flaking
=============

一个超慢的 ID 生成库，可以生成连续的数值流水号。灵感来自 [SnowFlake](https://github.com/twitter/snowflake) 的 [Golang 实现](https://github.com/zheng-ji/goSnowFlake)。

像这样：

```
	uid_test.go:33: 10002016061718222627096
	uid_test.go:33: 10002016061718222627097
	uid_test.go:33: 10002016061718222627098
	uid_test.go:33: 10002016061718222627099
	uid_test.go:33: 10002016061718222627101
	uid_test.go:33: 10002016061718222627102
	uid_test.go:33: 10002016061718222627103
	uid_test.go:33: 10002016061718222627104
	uid_test.go:33: 10002016061718222627105
	uid_test.go:33: 10002016061718222627106
	uid_test.go:33: 10002016061718222627107
```

流水号生成规则：

- 每毫秒最多产生 99（01 ~ 99）个 ID
- 前四位是 Worker ID
- 后二位是序号，不足两位，补零

## License

见 LICENSE。
