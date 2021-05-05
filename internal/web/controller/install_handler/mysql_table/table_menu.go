package mysql_table

//CREATE TABLE `menu` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',
//`name` varchar(32) NOT NULL DEFAULT '' COMMENT '菜单名称',
//`link` varchar(100) NOT NULL DEFAULT '' COMMENT '链接地址',
//`icon` varchar(60) NOT NULL DEFAULT '' COMMENT '图标',
//`level` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '菜单类型 1:一级菜单 2:二级菜单',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是 -1:否',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是 -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='左侧菜单栏表';

func CreateMenuTableSql() (sql string) {
	sql = "CREATE TABLE `menu` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',"
	sql += "`name` varchar(32) NOT NULL DEFAULT '' COMMENT '菜单名称',"
	sql += "`link` varchar(100) NOT NULL DEFAULT '' COMMENT '链接地址',"
	sql += "`icon` varchar(60) NOT NULL DEFAULT '' COMMENT '图标',"
	sql += "`level` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '菜单类型 1:一级菜单 2:二级菜单',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是 -1:否',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='左侧菜单栏表';"

	return
}
