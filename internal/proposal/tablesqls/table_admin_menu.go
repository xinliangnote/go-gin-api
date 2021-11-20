package tablesqls

//CREATE TABLE `admin_menu` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
//`menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单栏ID',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//PRIMARY KEY (`id`),
//KEY `idx_admin_id` (`admin_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员菜单栏表';

func CreateAdminMenuTableSql() (sql string) {
	sql = "CREATE TABLE `admin_menu` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',"
	sql += "`menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单栏ID',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "PRIMARY KEY (`id`),"
	sql += "KEY `idx_admin_id` (`admin_id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员菜单栏表';"

	return
}

func CreateAdminMenuTableDataSql() (sql string) {
	sql = "INSERT INTO `admin_menu` (`id`, `admin_id`, `menu_id`, `created_user`) VALUES"
	sql += "(1, 1, 16, 'init'),"
	sql += "(2, 1, 21, 'init'),"
	sql += "(3, 1, 20, 'init'),"
	sql += "(4, 1, 19, 'init'),"
	sql += "(5, 1, 18, 'init'),"
	sql += "(6, 1, 17, 'init'),"
	sql += "(7, 1, 13, 'init'),"
	sql += "(8, 1, 15, 'init'),"
	sql += "(9, 1, 14, 'init'),"
	sql += "(10, 1, 10, 'init'),"
	sql += "(11, 1, 12, 'init'),"
	sql += "(12, 1, 11, 'init'),"
	sql += "(13, 1, 7, 'init'),"
	sql += "(14, 1, 9, 'init'),"
	sql += "(15, 1, 8, 'init'),"
	sql += "(16, 1, 4, 'init'),"
	sql += "(17, 1, 6, 'init'),"
	sql += "(18, 1, 5, 'init'),"
	sql += "(19, 1, 1, 'init'),"
	sql += "(20, 1, 3, 'init'),"
	sql += "(21, 1, 2, 'init'),"
	sql += "(22, 1, 22, 'init'),"
	sql += "(23, 1, 23, 'init'),"
	sql += "(24, 1, 24, 'init'),"
	sql += "(25, 1, 25, 'init');"

	return
}
