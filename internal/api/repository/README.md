## repository

#### 数据访问层。

- `./db_repo` 访问 DB 数据
- `./cache_repo` 访问 Cache 数据

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

