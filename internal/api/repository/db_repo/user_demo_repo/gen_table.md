#### xin_ceshi.user_demo 
用户Demo表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | user_name | 用户名 | varchar(32) |  | NO |  |  |
| 3 | nick_name | 昵称 | varchar(100) |  | NO |  |  |
| 4 | mobile | 手机号 | varchar(20) |  | NO |  |  |
| 5 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 6 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 7 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
