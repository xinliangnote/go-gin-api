#### go_gin_api.authorized 
已授权的调用方表

| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |
| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |
| 1 | id | 主键 | int32 | true | true |   |
| 2 | business_key | 调用方key | string | false | true | ''::character varying  |
| 3 | business_secret | 调用方secret | string | false | true | ''::character varying  |
| 4 | business_developer | 调用方对接人 | string | false | true | ''::character varying  |
| 5 | remark | 备注 | string | false | true | ''::character varying  |
| 6 | is_used | 是否启用 1:是  -1:否 | int32 | false | true | 1  |
| 7 | is_deleted | 是否删除 1:是  -1:否 | int32 | false | true | '-1'::integer  |
| 8 | created_at | 创建时间 | time.Time | false | true | CURRENT_TIMESTAMP  |
| 9 | created_user | 创建人 | string | false | true | ''::character varying  |
| 10 | updated_at | 更新时间 | time.Time | false | false |   |
| 11 | updated_user | 更新人 | string | false | true | ''::character varying  |
