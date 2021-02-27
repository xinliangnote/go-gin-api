## repository

#### 数据访问层。

- `./db_repo` 访问 DB 数据
- `./cache_repo` 访问 Cache 数据
- `./third_party_request` 访问外部 HTTP 接口数据。

#### SQL 建议：
- 建议每张表需包含字段：主键(id)、标记删除(is_deteled)、创建时间(created_at)、更新时间(updated_at) 

```mysql
`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
```

#### 命名规范：

- 包名应以 `_repo` 结尾；
- `./db_repo` 目录下的包名以 `数据表名`+ `_repo` 命名；

#### 脚本生成 MySQL CURD

1. 定义生成的表，设置 config 中 cmd.genTables，可以自定义设置多张表，为空表示生成库中所有的表，如果设置多个表可用','分割；
1. 在根目录下执行脚本文件：`./scripts/gormgen.sh`；

以用户表（user_demo）为例：
- 结构体文件：user_demo_repo/gen_model.go；
- CURD 方法文件：user_demo_repo/gen_user_demo.go；
- 表结构 MD 文件：user_demo_repo/gen_table.md；
