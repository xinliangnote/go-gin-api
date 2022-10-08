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

/*
CREATE TABLE admin_menu
(
    id           integer  primary key ,
    admin_id     smallint NOT NULL DEFAULT '0' ,
    menu_id      integer  NOT NULL DEFAULT '0' ,
    created_at   timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    created_user varchar(60) NOT NULL DEFAULT ''
) ;

CREATE INDEX idx_admin_id ON admin_menu (admin_id);


CREATE
OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_at
= now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_table_name_update_at
    BEFORE UPDATE
    ON admin_menu
    FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

comment
on table admin_menu is '管理员菜单栏表';
comment
on column admin_menu.id is '主键';
comment
on column admin_menu.admin_id is '管理员ID';
comment
on column admin_menu.menu_id is '菜单栏ID';
comment
on column admin_menu.created_at is '创建时间';
comment
on column admin_menu.created_user is '创建人';




*/

func CreateAdminMenuTablePGSql() (sql string) {
	sql = `CREATE TABLE admin_menu
		(
			id           integer  primary key ,
			admin_id     smallint NOT NULL DEFAULT '0' ,
			menu_id      integer  NOT NULL DEFAULT '0' ,
			created_at   timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ,
			created_user varchar(60) NOT NULL DEFAULT ''
		) ;

		CREATE INDEX idx_admin_id ON admin_menu(admin_id);


		CREATE
		OR REPLACE FUNCTION update_modified_column()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.update_at
		= now();
		RETURN NEW;
		END;
		$$ language 'plpgsql';
		
		CREATE TRIGGER update_table_name_update_at
			BEFORE UPDATE
			ON admin_menu
			FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
		
		comment
		on table admin_menu is '管理员菜单栏表';
		comment
		on column admin_menu.admin_id is '管理员ID';
		comment
		on column admin_menu.menu_id is '菜单栏ID';
		comment
		on column admin_menu.created_at is '创建时间';
		comment
		on column admin_menu.created_user is '创建人';`

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
	sql += "(26, 1, 26, 'init'),"
	sql += "(27, 1, 27, 'init'),"
	sql += "(25, 1, 25, 'init');"

	return
}

func CreateAdminMenuTableDataPGSql() (sql string) {
	sql = "INSERT INTO admin_menu (id, admin_id, menu_id, created_user) VALUES"
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
	sql += "(26, 1, 26, 'init'),"
	sql += "(27, 1, 27, 'init'),"
	sql += "(25, 1, 25, 'init');"

	return
}
