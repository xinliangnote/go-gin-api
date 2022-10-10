#### go_gin_api.admin 
管理员表

| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |
| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |
| 1 | id | 主键 | int32 | true | true |   |
| 2 | username | 用户名 | string | false | true | ''::character varying  |
| 3 | password | 密码 | string | false | true | ''::character varying  |
| 4 | nickname | 昵称 | string | false | true | ''::character varying  |
| 5 | mobile | 手机号 | string | false | true | ''::character varying  |
| 6 | is_used | 是否启用 1:是  -1:否 | int32 | false | true | '1'::smallint  |
| 7 | is_deleted | 是否删除 1:是  -1:否 | int32 | false | true | '-1'::smallint  |
| 8 | created_at | 创建时间 | time.Time | false | true | CURRENT_TIMESTAMP  |
| 9 | created_user | 创建人 | string | false | true | ''::character varying  |
| 10 | updated_at | 更新时间 | time.Time | false | false |   |
| 11 | updated_user | 更新人 | string | false | true | ''::character varying  |
