## repository

数据访问层。

- `./db_repo` 访问 DB 数据
- `./cache_repo` 访问 Cache 数据
- `./third_party_request` 访问外部 HTTP 接口数据。

SQL 建议：
- 禁止使用 SQL k v 拼接，好处是避免 SQL 注入；
- 禁止使用连表查询，好处是易扩展，比如分库分表；
- 禁止使用万能方法，好处是便于后期维护，比如字段调整；
- 禁止使用删除方法，好处是避免数据丢失；
- 建议每张表需包含字段：主键(id)、标记删除(is_deteled)、创建时间(created_at)、更新时间(updated_at) 

```mysql
`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
```
什么是万能方法？

指的是特别灵活的查询，比如通过非固定的参数返回全部字段，建议做到需要什么返回什么，不要返回大而全的数据，更新时也不能传递什么参数更新什么参数，更新字段要提前约定好。

命名规范：

- 包名应以 `_repo` 结尾；
- `./db_repo` 目录下的包名以 `数据表名`+ `_repo` 命名；
