## Goa 微服务脚手架

### 依赖

- mysql 5.7+
- redis 4.0+
- golang 1.15

### 配置

参考 `config.sample.yml`

### 编译

```
make build
```

### 启动服务

```shell
boot --config config.yml runserver
```

### 接口文档

#### 错误码

- 60001 参数错误
