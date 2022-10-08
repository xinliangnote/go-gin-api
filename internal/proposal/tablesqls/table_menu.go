package tablesqls

//CREATE TABLE `menu` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',
//`name` varchar(32) NOT NULL DEFAULT '' COMMENT '菜单名称',
//`link` varchar(100) NOT NULL DEFAULT '' COMMENT '链接地址',
//`icon` varchar(60) NOT NULL DEFAULT '' COMMENT '图标',
//`level` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '菜单类型 1:一级菜单 2:二级菜单',
//`sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
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
	sql += "`sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',"
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

/*
CREATE TABLE menu
(
    id           integer primary key ,
    pid          integer NOT NULL DEFAULT '0' ,
    name         varchar(32)  NOT NULL DEFAULT '' ,
    link         varchar(100) NOT NULL DEFAULT '' ,
    icon         varchar(60)  NOT NULL DEFAULT '' ,
    level        smallint NOT NULL DEFAULT '1' ,
    sort         integer NOT NULL DEFAULT '0' ,
    is_used      smallint NOT NULL DEFAULT '1' ,
    is_deleted   smallint NOT NULL DEFAULT '-1' ,
    created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    created_user varchar(60)  NOT NULL DEFAULT '',
    updated_at   timestamp     ,
    updated_user varchar(60)  NOT NULL DEFAULT ''
);



CREATE
OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_at
= now();
RETURN NEW;
END;
$$
language 'plpgsql';

CREATE TRIGGER update_table_name_update_at
    BEFORE UPDATE
    ON menu
    FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

comment
on table menu is '左侧菜单栏表';
comment
on column menu.id is '主键';
comment
on column menu.pid is '父类ID';
comment
on column menu.name is '菜单名称';
comment
on column menu.link is '链接地址';
comment
on column menu.icon is '图标';
comment
on column menu.level is '菜单类型 1:一级菜单 2:二级菜单';
comment
on column menu.sort is '排序';
comment
on column menu.is_used is '是否启用 1:是 -1:否';
comment
on column menu.is_deleted is '是否删除 1:是 -1:否';
comment
on column menu.created_at is '创建时间';
comment
on column menu.created_user is '创建人';
comment
on column menu.updated_at is '更新时间';
comment
on column menu.updated_user is '更新人';

*/

func CreateMenuTablePGSql() (sql string) {
	sql = `CREATE TABLE menu
		(
			id           integer primary key ,
			pid          integer NOT NULL DEFAULT '0' ,
			name         varchar(32)  NOT NULL DEFAULT '' ,
			link         varchar(100) NOT NULL DEFAULT '' ,
			icon         varchar(60)  NOT NULL DEFAULT '' ,
			level        smallint NOT NULL DEFAULT '1' ,
			sort         integer NOT NULL DEFAULT '0' ,
			is_used      smallint NOT NULL DEFAULT '1' ,
			is_deleted   smallint NOT NULL DEFAULT '-1' ,
			created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
			created_user varchar(60)  NOT NULL DEFAULT '',
			updated_at   timestamp     ,
			updated_user varchar(60)  NOT NULL DEFAULT ''
		);



		CREATE
		OR REPLACE FUNCTION update_modified_column()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.update_at
		= now();
		RETURN NEW;
		END;
		$$
		language 'plpgsql';
		
		CREATE TRIGGER update_table_name_update_at
			BEFORE UPDATE
			ON menu
			FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

		comment
		on table menu is '左侧菜单栏表';
		comment
		on column menu.id is '主键';
		comment
		on column menu.pid is '父类ID';
		comment
		on column menu.name is '菜单名称';
		comment
		on column menu.link is '链接地址';
		comment
		on column menu.icon is '图标';
		comment
		on column menu.level is '菜单类型 1:一级菜单 2:二级菜单';
		comment
		on column menu.sort is '排序';
		comment
		on column menu.is_used is '是否启用 1:是 -1:否';
		comment
		on column menu.is_deleted is '是否删除 1:是 -1:否';
		comment
		on column menu.created_at is '创建时间';
		comment
		on column menu.created_user is '创建人';
		comment
		on column menu.updated_at is '更新时间';
		comment
		on column menu.updated_user is '更新人';`

	return
}

func CreateMenuTableDataSql() (sql string) {
	sql = "INSERT INTO `menu` (`id`, `pid`, `name`, `link`, `icon`, `level`, `sort`, `created_user`) VALUES"
	sql += "(1, 0, '配置信息', '', 'mdi-settings-box', 1, 10, 'init'),"
	sql += "(2, 1, '告警邮箱', '/config/email', '', 2, 101, 'init'),"
	sql += "(3, 1, '错误码', '/config/code', '', 2, 102, 'init'),"
	sql += "(4, 0, '代码生成器', '', 'mdi-code-not-equal-variant', 1, 20, 'init'),"
	sql += "(5, 4, '生成数据表 CURD', '/generator/gorm', '', 2, 201, 'init'),"
	sql += "(6, 4, '生成控制器方法', '/generator/handler', '', 2, 202, 'init'),"
	sql += "(26, 4, 'PgSQL生成数据表 CURD', '/generator/pgsql/gorm', '', 2, 203, 'init'),"
	//sql += "(27, 4, 'PgSQL生成控制器方法', '/generator/pgsql/gorm/execute', '', 2, 204, 'init'),"

	sql += "(7, 0, '授权调用方', '', 'mdi-playlist-check', 1, 30, 'init'),"
	sql += "(8, 7, '调用方', '/authorized/list', '', 2, 301, 'init'),"
	sql += "(9, 7, '使用说明', '/authorized/demo', '', 2, 302, 'init'),"
	sql += "(10, 0, '系统管理员', '', 'mdi-account', 1, 50, 'init'),"
	sql += "(11, 10, '管理员', '/admin/list', '', 2, 501, 'init'),"
	sql += "(12, 10, '菜单管理', '/admin/menu', '', 2, 502, 'init'),"
	sql += "(13, 0, '查询小助手', '', 'mdi-database-search', 1, 60, 'init'),"
	sql += "(14, 13, '查询缓存', '/tool/cache', '', 2, 601, 'init'),"
	sql += "(15, 13, '查询数据', '/tool/data', '', 2, 602, 'init'),"
	sql += "(16, 0, '实用工具箱', '', 'mdi-tools', 1, 70, 'init'),"
	sql += "(17, 16, 'Hashids', '/tool/hashids', '', 2, 702, 'init'),"
	sql += "(18, 16, '调用日志', '/tool/logs', '', 2, 703, 'init'),"
	sql += "(19, 16, '接口文档', '/swagger/index.html', '', 2, 704, 'init'),"
	sql += "(20, 16, 'GraphQL', '/graphql', '', 2, 705, 'init'),"
	sql += "(21, 16, '接口指标', '/metrics', '', 2, 706, 'init'),"
	sql += "(22, 16, '服务升级', '/upgrade', '', 2, 701, 'init'),"
	sql += "(23, 0, '后台任务', '', 'mdi-av-timer', 1, 40, 'init'),"
	sql += "(24, 23, '任务列表', '/cron/list', '', 2, 401, 'init'),"
	sql += "(25, 16, 'WebSocket', '/tool/websocket', '', 2, 707, 'init');"

	return
}

func CreateMenuTableDataPGSql() (sql string) {
	sql = "INSERT INTO menu (id, pid, name, link, icon, level, sort, created_user) VALUES"
	sql += "(1, 0, '配置信息', '', 'mdi-settings-box', 1, 10, 'init'),"
	sql += "(2, 1, '告警邮箱', '/config/email', '', 2, 101, 'init'),"
	sql += "(3, 1, '错误码', '/config/code', '', 2, 102, 'init'),"
	sql += "(4, 0, '代码生成器', '', 'mdi-code-not-equal-variant', 1, 20, 'init'),"
	sql += "(5, 4, '生成数据表 CURD', '/generator/gorm', '', 2, 201, 'init'),"
	sql += "(6, 4, '生成控制器方法', '/generator/handler', '', 2, 202, 'init'),"
	sql += "(26, 4, 'PgSQL生成数据表 CURD', '/generator/pgsql/gorm', '', 2, 203, 'init'),"
	//sql += "(27, 4, 'PgSQL生成控制器方法', '/generator/pgsql/gorm/execute', '', 2, 204, 'init'),"

	sql += "(7, 0, '授权调用方', '', 'mdi-playlist-check', 1, 30, 'init'),"
	sql += "(8, 7, '调用方', '/authorized/list', '', 2, 301, 'init'),"
	sql += "(9, 7, '使用说明', '/authorized/demo', '', 2, 302, 'init'),"
	sql += "(10, 0, '系统管理员', '', 'mdi-account', 1, 50, 'init'),"
	sql += "(11, 10, '管理员', '/admin/list', '', 2, 501, 'init'),"
	sql += "(12, 10, '菜单管理', '/admin/menu', '', 2, 502, 'init'),"
	sql += "(13, 0, '查询小助手', '', 'mdi-database-search', 1, 60, 'init'),"
	sql += "(14, 13, '查询缓存', '/tool/cache', '', 2, 601, 'init'),"
	sql += "(15, 13, '查询数据', '/tool/data', '', 2, 602, 'init'),"
	sql += "(16, 0, '实用工具箱', '', 'mdi-tools', 1, 70, 'init'),"
	sql += "(17, 16, 'Hashids', '/tool/hashids', '', 2, 702, 'init'),"
	sql += "(18, 16, '调用日志', '/tool/logs', '', 2, 703, 'init'),"
	sql += "(19, 16, '接口文档', '/swagger/index.html', '', 2, 704, 'init'),"
	sql += "(20, 16, 'GraphQL', '/graphql', '', 2, 705, 'init'),"
	sql += "(21, 16, '接口指标', '/metrics', '', 2, 706, 'init'),"
	sql += "(22, 16, '服务升级', '/upgrade', '', 2, 701, 'init'),"
	sql += "(23, 0, '后台任务', '', 'mdi-av-timer', 1, 40, 'init'),"
	sql += "(24, 23, '任务列表', '/cron/list', '', 2, 401, 'init'),"
	sql += "(25, 16, 'WebSocket', '/tool/websocket', '', 2, 707, 'init');"

	return
}
