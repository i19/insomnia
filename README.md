### 目的
验证用户上传的 `spec` 是否符合自定义的规范，该规范使用 `yaml` 结构进行描述，支持 `formatter` 扩展。同时利用已有的权限数据做业务隔离
.
目前添加了 version 和 describe 两个自定义 `formatter`
### 设计
#### 存储
使用 `etcd` ，主要使用其通知机制，如果是单节点运行也可以考虑使用 MySQL 或其他重新存储


#### 接口
见 `swagger`
```bash
make build
make doc
```

### 测试
目前针对 `rule_engine` 做了单元测试
api 部分做了上传规则，以及规则验证的结果测试以及不同权限用户的测试
```bash
make test

```