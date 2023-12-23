#### go_gin_api.menu 
左侧菜单栏表

| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |
| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |
| 1 | id | 主键 | int32 | true | true |   |
| 2 | pid | 父类ID | int32 | false | true | 0  |
| 3 | name | 菜单名称 | string | false | true | ''::character varying  |
| 4 | link | 链接地址 | string | false | true | ''::character varying  |
| 5 | icon | 图标 | string | false | true | ''::character varying  |
| 6 | level | 菜单类型 1:一级菜单 2:二级菜单 | int32 | false | true | '1'::smallint  |
| 7 | sort | 排序 | int32 | false | true | 0  |
| 8 | is_used | 是否启用 1:是 -1:否 | int32 | false | true | '1'::smallint  |
| 9 | is_deleted | 是否删除 1:是 -1:否 | int32 | false | true | '-1'::smallint  |
| 10 | created_at | 创建时间 | time.Time | false | true | CURRENT_TIMESTAMP  |
| 11 | created_user | 创建人 | string | false | true | ''::character varying  |
| 12 | updated_at | 更新时间 | time.Time | false | false |   |
| 13 | updated_user | 更新人 | string | false | true | ''::character varying  |
