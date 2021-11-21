#### go_gin_api.menu_action 
功能权限表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int unsigned | PRI | NO | auto_increment |  |
| 2 | menu_id | 菜单栏ID | int unsigned | MUL | NO |  | 0 |
| 3 | method | 请求方式 | varchar(30) |  | NO |  |  |
| 4 | api | 请求地址 | varchar(100) |  | NO |  |  |
| 5 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 6 | created_at | 创建时间 | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 7 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 8 | updated_at | 更新时间 | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 9 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
