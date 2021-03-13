## 执行命令
在根目录下执行脚本：`./scripts/gormgen.sh addr user pass name tables`；
- addr：数据库地址，例如：127.0.0.1:3306
- user：账号，例如：root
- pass：密码，例如：root
- name：数据库名称，例如：go_gin_api
- tables：表名，默认为 *，多个表名可用“,”分割，例如：user_demo

例如：
```
./scripts/gormgen.sh 127.0.0.1:3306 root root go_gin_api user_demo
```

## 参考
- https://github.com/MohamedBassem/gormgen