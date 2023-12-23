package tablesqls

//CREATE TABLE `admin` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
//`password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
//`nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
//`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
//PRIMARY KEY (`id`),
//UNIQUE KEY `unique_username` (`username`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

func CreateAdminTableSql() (sql string) {
	sql = "CREATE TABLE `admin` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',"
	sql += "`password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',"
	sql += "`nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',"
	sql += "`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`),"
	sql += "UNIQUE KEY `unique_username` (`username`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';"

	return
}

/*
CREATE TABLE admin
(
    id           integer primary key ,
    username     varchar(32)  NOT NULL DEFAULT '' ,
    password     varchar(100) NOT NULL DEFAULT '' ,
    nickname     varchar(60)  NOT NULL DEFAULT '' ,
    mobile       varchar(20)  NOT NULL DEFAULT '' ,
    is_used      smallint NOT NULL DEFAULT '1' ,
    is_deleted   smallint NOT NULL DEFAULT '-1' ,
    created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    created_user varchar(60)  NOT NULL DEFAULT '' ,
    updated_at   timestamp   ,
    updated_user varchar(60)  NOT NULL DEFAULT '' ,
    UNIQUE(username)
);


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
    ON admin
    FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

commnet
on table admin is '管理员表';
comment
on column admin.id is '主键';
comment
on column admin.username is '用户名';
comment
on column admin.password is '密码';
comment
on column admin.nickname is '昵称';
comment
on column admin.mobile is '手机号';
comment
on column admin.is_used is '是否启用 1:是  -1:否';
comment
on column admin.is_deleted is '是否删除 1:是  -1:否';
comment
on column admin.created_at is '创建时间';
comment
on column admin.created_user is '创建人';
comment
on column admin.created_user is '创建人';
comment
on column admin.updated_at is '更新时间';
comment
on column admin.updated_user is '更新人';
*/

func CreateAdminTablePGSql() (sql string) {
	sql = `CREATE TABLE admin
		(
			id           integer primary key ,
			username     varchar(32)  NOT NULL DEFAULT '' ,
			password     varchar(100) NOT NULL DEFAULT '' ,
			nickname     varchar(60)  NOT NULL DEFAULT '' ,
			mobile       varchar(20)  NOT NULL DEFAULT '' ,
			is_used      smallint NOT NULL DEFAULT '1' ,
			is_deleted   smallint NOT NULL DEFAULT '-1' ,
			created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
			created_user varchar(60)  NOT NULL DEFAULT '' ,
			updated_at   timestamp   ,
			updated_user varchar(60)  NOT NULL DEFAULT '' ,
			UNIQUE(username)
		);


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
			ON admin
			FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
		
		comment
		on table admin is '管理员表';
		comment
		on column admin.id is '主键';
		comment
		on column admin.username is '用户名';
		comment
		on column admin.password is '密码';
		comment
		on column admin.nickname is '昵称';
		comment
		on column admin.mobile is '手机号';
		comment
		on column admin.is_used is '是否启用 1:是  -1:否';
		comment
		on column admin.is_deleted is '是否删除 1:是  -1:否';
		comment
		on column admin.created_at is '创建时间';
		comment
		on column admin.created_user is '创建人';
		comment
		on column admin.created_user is '创建人';
		comment
		on column admin.updated_at is '更新时间';
		comment
		on column admin.updated_user is '更新人';`

	return
}
func CreateAdminTableDataSql() (sql string) {
	sql = "INSERT INTO `admin` (`id`, `username`, `password`, `nickname`, `mobile`, `created_user`) VALUES"
	sql += "(1, 'admin', 'f78382de80cf583cf854bbac0b6e796fbde36fe2739ca4ae072637010f179cb0', '管理员', '13888888888', 'init');"
	return
}
func CreateAdminTableDataPGSql() (sql string) {
	sql = "INSERT INTO admin (id, username, password, nickname, mobile, created_user) VALUES"
	sql += "(1, 'admin', 'f78382de80cf583cf854bbac0b6e796fbde36fe2739ca4ae072637010f179cb0', '管理员', '13888888888', 'init');"
	return
}
