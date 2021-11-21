#### go_gin_api.menu 
左侧菜单栏表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int unsigned | PRI | NO | auto_increment |  |
| 2 | pid | 父类ID | int unsigned |  | NO |  | 0 |
| 3 | name | 菜单名称 | varchar(32) |  | NO |  |  |
| 4 | link | 链接地址 | varchar(100) |  | NO |  |  |
| 5 | icon | 图标 | varchar(60) |  | NO |  |  |
| 6 | level | 菜单类型 1:一级菜单 2:二级菜单 | tinyint unsigned |  | NO |  | 1 |
| 7 | sort | 排序 | int unsigned |  | NO |  | 0 |
| 8 | is_used | 是否启用 1:是 -1:否 | tinyint(1) |  | NO |  | 1 |
| 9 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 10 | created_at | 创建时间 | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 11 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 12 | updated_at | 更新时间 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 13 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
