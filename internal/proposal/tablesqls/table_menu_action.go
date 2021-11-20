package tablesqls

//CREATE TABLE `menu_action` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单栏ID',
//`method` varchar(30) NOT NULL DEFAULT '' COMMENT '请求方式',
//`api` varchar(100) NOT NULL DEFAULT '' COMMENT '请求地址',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
//PRIMARY KEY (`id`),
//KEY `idx_menu_id` (`menu_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='功能权限表';

func CreateMenuActionTableSql() (sql string) {
	sql = "CREATE TABLE `menu_action` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单栏ID',"
	sql += "`method` varchar(30) NOT NULL DEFAULT '' COMMENT '请求方式',"
	sql += "`api` varchar(100) NOT NULL DEFAULT '' COMMENT '请求地址',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`),"
	sql += "KEY `idx_menu_id` (`menu_id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='功能权限表';"

	return
}

func CreateMenuActionTableDataSql() (sql string) {
	sql = "INSERT INTO `menu_action` (`id`, `menu_id`, `method`, `api`, `created_user`) VALUES"
	sql += "(1, 17, 'GET', '/api/tool/hashids/**', 'init'),"
	sql += "(2, 14, 'POST', '/api/tool/cache/search', 'init'),"
	sql += "(3, 14, 'PATCH', '/api/tool/cache/clear', 'init'),"
	sql += "(4, 15, 'GET', '/api/tool/data/dbs', 'init'),"
	sql += "(5, 15, 'POST', '/api/tool/data/mysql', 'init'),"
	sql += "(6, 15, 'POST', '/api/tool/data/tables', 'init'),"
	sql += "(7, 2, 'PATCH', '/api/config/email', 'init'),"
	sql += "(8, 5, 'POST', '/generator/gorm/execute', 'init'),"
	sql += "(9, 6, 'POST', '/generator/handler/execute', 'init'),"
	sql += "(10, 8, 'GET', '/authorized/add', 'init'),"
	sql += "(11, 8, 'GET', '/authorized/api/*', 'init'),"
	sql += "(12, 8, 'GET', '/api/authorized', 'init'),"
	sql += "(13, 8, 'PATCH', '/api/authorized/used', 'init'),"
	sql += "(14, 8, 'DELETE', '/api/authorized/*', 'init'),"
	sql += "(15, 8, 'POST', '/api/authorized', 'init'),"
	sql += "(16, 8, 'GET', '/api/authorized_api', 'init'),"
	sql += "(17, 8, 'POST', '/api/authorized_api', 'init'),"
	sql += "(18, 8, 'DELETE', '/api/authorized_api/*', 'init'),"
	sql += "(19, 11, 'GET', '/admin/add', 'init'),"
	sql += "(20, 11, 'POST', '/api/admin', 'init'),"
	sql += "(21, 11, 'GET', '/api/admin', 'init'),"
	sql += "(22, 11, 'PATCH', '/api/admin/used', 'init'),"
	sql += "(23, 11, 'PATCH', '/api/admin/reset_password/*', 'init'),"
	sql += "(24, 11, 'DELETE', '/api/admin/*', 'init'),"
	sql += "(25, 11, 'GET', '/admin/action/*', 'init'),"
	sql += "(26, 11, 'GET', '/api/admin/menu/*', 'init'),"
	sql += "(27, 11, 'POST', '/api/admin/menu', 'init'),"
	sql += "(28, 12, 'GET', '/admin/menu_action/*', 'init'),"
	sql += "(29, 12, 'GET', '/api/menu', 'init'),"
	sql += "(30, 12, 'DELETE', '/api/menu/*', 'init'),"
	sql += "(31, 12, 'GET', '/api/menu/*', 'init'),"
	sql += "(32, 12, 'PATCH', '/api/menu/used', 'init'),"
	sql += "(33, 12, 'POST', '/api/menu', 'init'),"
	sql += "(34, 12, 'GET', '/api/menu_action', 'init'),"
	sql += "(35, 12, 'POST', '/api/menu_action', 'init'),"
	sql += "(36, 12, 'DELETE', '/api/menu_action/*', 'init'),"
	sql += "(37, 22, 'POST', '/upgrade/execute', 'init'),"
	sql += "(38, 11, 'PATCH', '/api/admin/offline', 'init'),"
	sql += "(39, 12, 'PATCH', '/api/menu/sort', 'init'),"
	sql += "(40, 24, 'GET', '/cron/add', 'init'),"
	sql += "(41, 24, 'GET', '/cron/edit/*', 'init'),"
	sql += "(42, 24, 'POST', '/api/cron', 'init'),"
	sql += "(43, 24, 'POST', '/api/cron/*', 'init'),"
	sql += "(44, 24, 'GET', '/api/cron', 'init'),"
	sql += "(45, 24, 'GET', '/api/cron/*', 'init'),"
	sql += "(46, 24, 'PATCH', '/api/cron/used', 'init'),"
	sql += "(47, 24, 'PATCH', '/api/cron/exec/*', 'init'),"
	sql += "(48, 25, 'POST', '/api/tool/send_message', 'init');"

	return
}
