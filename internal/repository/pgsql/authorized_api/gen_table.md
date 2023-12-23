#### go_gin_api.authorized_api 
已授权接口地址表

| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |
| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |
| 1 | id | 主键 | int32 | true | true |   |
| 2 | business_key | 调用方key | string | false | true | ''::character varying  |
| 3 | method | 请求方式 | string | false | true | ''::character varying  |
| 4 | api | 请求地址 | string | false | true | ''::character varying  |
| 5 | is_deleted | 是否删除 1:是  -1:否 | int32 | false | true | '-1'::smallint  |
| 6 | created_at | 创建时间 | time.Time | false | true | CURRENT_TIMESTAMP  |
| 7 | created_user | 创建人 | string | false | true | ''::character varying  |
| 8 | updated_at | 更新时间 | time.Time | false | false |   |
| 9 | updated_user | 更新人 | string | false | true | ''::character varying  |
