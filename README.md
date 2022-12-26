# boot-plugin-gorm-mysql-gorm-mysql

[![Build Status](https://github.com/jsmzr/boot-plugin-gorm-mysql/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-plugin-gorm-mysql/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-plugin-gorm-mysql/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-plugin-gorm-mysql)

基于 boot 的 gorm mysql 插件

## 使用方式

1. 拉取依赖，`go get -u github.com/jsmzr/boot-plugin-gorm-mysql`
2. 在 `main.go` 中导入，`import _ "github.com/jsmzr/boot-plugin-gorm-mysql`
3. 通过 `import github.com/jsmzr/boot-plugin-gorm-myql/db`, 获取创建好的连接

### 配置

```yaml
boot:
  gorm:
    enabled: true
    order: 4
    database: demo
    username: mysql_test
    password: demo123321
    host: 127.0.0.1
    port: 3306
    singularTable: true
    noLowerCase: false
```