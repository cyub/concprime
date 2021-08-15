# concprime

Benchmark the Effect of GOMAXPROCS In Docker with CPU limit.

通过获取小于n的所有素数运算，来进行Go运行在CPU限制的Docker环境下的CPU负载测试。

## 使用

### 测试

```bash
docker run --rm -it cyub/concprime:1.0
docker run --cpus 1 --rm -it cyub/concprime:1.0
docker run --cpus 1 --rm -e GOMAXPROCS=1 -it cyub/concprime:1.0
```

### 基准测试

```bash
make benchmark
```

### 构建镜像

```bash
make build-image
```

## 相关资料

- [Benchmark Cgroups Quota (i.e. Limits) and the Effect of GOMAXPROCS](https://github.com/embano1/gotutorials/tree/master/concprime)
- [Build your Go image](https://docs.docker.com/language/golang/build-images/)
- [Determine If A Number Is Prime Using The Go Programming Language](https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/)
- [How channel work in using channel to find prime number problem?](https://stackoverflow.com/questions/65287453/how-channel-work-in-using-channel-to-find-prime-number-problem)