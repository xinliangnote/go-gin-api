#### go_gin_api.admin_menu 
管理员菜单栏表

| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |
| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |
| 1 | id | 主键 | int32 | true | true |   |
| 2 | admin_id | 管理员ID | int32 | false | true | '0'::smallint  |
| 3 | menu_id | 菜单栏ID | int32 | false | true | 0  |
| 4 | created_at | 创建时间 | time.Time | false | true | CURRENT_TIMESTAMP  |
| 5 | created_user | 创建人 | string | false | true | ''::character varying  |
