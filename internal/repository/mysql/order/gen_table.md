#### mydb.order 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int unsigned | PRI | NO | auto_increment |  |
| 2 | order_no |  | char(32) | UNI | NO |  |  |
| 3 | order_fee | () | int unsigned |  | NO |  | 0 |
| 4 | status |  1:  2: | tinyint unsigned |  | NO |  | 1 |
| 5 | is_deleted |  1:  -1: | tinyint(1) |  | NO |  | -1 |
| 6 | created_at |  | timestamp |  | NO | DEFAULT_GENERATED | CURRENT_TIMESTAMP |
| 7 | created_user |  | varchar(60) |  | NO |  |  |
| 8 | updated_at |  | timestamp |  | NO | DEFAULT_GENERATED on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 9 | updated_user |  | varchar(60) |  | NO |  |  |
