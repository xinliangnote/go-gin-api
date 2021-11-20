package tablesqls

//CREATE TABLE `cron_task` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
//`spec` varchar(64) NOT NULL DEFAULT '' COMMENT 'crontab 表达式',
//`command` varchar(255) NOT NULL DEFAULT '' COMMENT '执行命令',
//`protocol` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '执行方式 1:shell 2:http',
//`http_method` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT 'http 请求方式 1:get 2:post',
//`timeout` int(11) unsigned NOT NULL DEFAULT '60' COMMENT '超时时间(单位:秒)',
//`retry_times` tinyint(1) NOT NULL DEFAULT '3' COMMENT '重试次数',
//`retry_interval` int(11) NOT NULL DEFAULT '60' COMMENT '重试间隔(单位:秒)',
//`notify_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知',
//`notify_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '通知类型 1:邮件 2:webhook',
//`notify_receiver_email` varchar(255) NOT NULL DEFAULT '' COMMENT '通知者邮箱地址(多个用,分割)',
//`notify_keyword` varchar(255) NOT NULL DEFAULT '' COMMENT '通知匹配关键字(多个用,分割)',
//`remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
//PRIMARY KEY (`id`),
//KEY `idx_name` (`name`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台任务表';

func CreateCronTaskTableSql() (sql string) {
	sql = "CREATE TABLE `cron_task` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',"
	sql += "`spec` varchar(64) NOT NULL DEFAULT '' COMMENT 'crontab 表达式',"
	sql += "`command` varchar(255) NOT NULL DEFAULT '' COMMENT '执行命令',"
	sql += "`protocol` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '执行方式 1:shell 2:http',"
	sql += "`http_method` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT 'http 请求方式 1:get 2:post',"
	sql += "`timeout` int(11) unsigned NOT NULL DEFAULT '60' COMMENT '超时时间(单位:秒)',"
	sql += "`retry_times` tinyint(1) NOT NULL DEFAULT '3' COMMENT '重试次数',"
	sql += "`retry_interval` int(11) NOT NULL DEFAULT '60' COMMENT '重试间隔(单位:秒)',"
	sql += "`notify_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知',"
	sql += "`notify_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '通知类型 1:邮件 2:webhook',"
	sql += "`notify_receiver_email` varchar(255) NOT NULL DEFAULT '' COMMENT '通知者邮箱地址(多个用,分割)',"
	sql += "`notify_keyword` varchar(255) NOT NULL DEFAULT '' COMMENT '通知匹配关键字(多个用,分割)',"
	sql += "`remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`),"
	sql += "KEY `idx_name` (`name`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台任务表';"

	return
}
