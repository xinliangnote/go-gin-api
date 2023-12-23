#### go_gin_api.cron_task 
后台任务表

| 序号 | 名称 | 描述 | 类型 | 是否主键 | 是否为空  | 默认值 |
| :--: | :--: | :--: | :--: | :--:  | :--: | :--: |
| 1 | id | 主键 | int32 | true | true |   |
| 2 | name | 任务名称 | string | false | true | ''::character varying  |
| 3 | spec | crontab 表达式 | string | false | true | ''::character varying  |
| 4 | command | 执行命令 | string | false | true | ''::character varying  |
| 5 | protocol | 执行方式 1:shell 2:http | int32 | false | true | '1'::smallint  |
| 6 | http_method | http 请求方式 1:get 2:post | int32 | false | true | '1'::smallint  |
| 7 | timeout | 超时时间(单位:秒) | int32 | false | true | 60  |
| 8 | retry_times | 重试次数 | int32 | false | true | '3'::smallint  |
| 9 | retry_interval | 重试间隔(单位:秒) | int32 | false | true | 60  |
| 10 | notify_status | 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知 | int32 | false | true | '0'::smallint  |
| 11 | notify_type | 通知类型 1:邮件 2:webhook | int32 | false | true | '0'::smallint  |
| 12 | notify_receiver_email | 通知者邮箱地址(多个用,分割) | string | false | true | ''::character varying  |
| 13 | notify_keyword | 通知匹配关键字(多个用,分割) | string | false | true | ''::character varying  |
| 14 | remark | 备注 | string | false | true | ''::character varying  |
| 15 | is_used | 是否启用 1:是  -1:否 | int32 | false | true | '1'::smallint  |
| 16 | created_at | 创建时间 | time.Time | false | true | CURRENT_TIMESTAMP  |
| 17 | created_user | 创建人 | string | false | true | ''::character varying  |
| 18 | updated_at | 更新时间 | time.Time | false | true |   |
| 19 | updated_user | 更新人 | string | false | true | ''::character varying  |
