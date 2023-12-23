package tablesqls

//CREATE TABLE `authorized_api` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`business_key` varchar(32) NOT NULL DEFAULT '' COMMENT '调用方key',
//`method` varchar(30) NOT NULL DEFAULT '' COMMENT '请求方式',
//`api` varchar(100) NOT NULL DEFAULT '' COMMENT '请求地址',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='已授权接口地址表';

func CreateAuthorizedAPITableSql() (sql string) {
	sql = "CREATE TABLE `authorized_api` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`business_key` varchar(32) NOT NULL DEFAULT '' COMMENT '调用方key',"
	sql += "`method` varchar(30) NOT NULL DEFAULT '' COMMENT '请求方式',"
	sql += "`api` varchar(100) NOT NULL DEFAULT '' COMMENT '请求地址',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='已授权接口地址表';"

	return
}

/*
CREATE TABLE authorized_api
(
    id           integer primary key,
    business_key varchar(32)  NOT NULL DEFAULT '',
    method       varchar(30)  NOT NULL DEFAULT '',
    api          varchar(100) NOT NULL DEFAULT '',
    is_deleted   smallint     NOT NULL DEFAULT '-1',
    created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_user varchar(60)  NOT NULL DEFAULT '',
    updated_at   timestamp    ,
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
    ON authorized_api
    FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

commnet
on table authorized_api is '已授权接口地址表';
comment
on column authorized_api.id is '主键';
comment
on column authorized_api.business_key is '调用方key';
comment
on column authorized_api.method is '请求方式';
comment
on column authorized_api.api is '请求地址';
comment
on column authorized_api.is_deleted is '是否删除 1:是  -1:否';
comment
on column authorized_api.created_at is '创建时间';
comment
on column authorized_api.created_user is '创建人';
comment
on column authorized_api.updated_at is '更新时间';
comment
on column authorized_api.updated_user is '更新人';
*/

func CreateAuthorizedAPITablePGSql() (sql string) {
	sql = `CREATE TABLE authorized_api
		(
			id           integer primary key,
			business_key varchar(32)  NOT NULL DEFAULT '',
			method       varchar(30)  NOT NULL DEFAULT '',
			api          varchar(100) NOT NULL DEFAULT '',
			is_deleted   smallint     NOT NULL DEFAULT '-1',
			created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
			created_user varchar(60)  NOT NULL DEFAULT '',
			updated_at   timestamp    ,
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
			ON authorized_api
			FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

		comment
		on table authorized_api is '已授权接口地址表';
		comment
		on column authorized_api.id is '主键';
		comment
		on column authorized_api.business_key is '调用方key';
		comment
		on column authorized_api.method is '请求方式';
		comment
		on column authorized_api.api is '请求地址';
		comment
		on column authorized_api.is_deleted is '是否删除 1:是  -1:否';
		comment
		on column authorized_api.created_at is '创建时间';
		comment
		on column authorized_api.created_user is '创建人';
		comment
		on column authorized_api.updated_at is '更新时间';
		comment
		on column authorized_api.updated_user is '更新人';`

	return
}

func CreateAuthorizedAPITableDataSql() (sql string) {
	sql = "INSERT INTO `authorized_api` (`id`, `business_key`, `method`, `api`,`created_user`) VALUES"
	sql += "(1, 'admin', 'GET', '/api/**', 'init'),"
	sql += "(2, 'admin', 'POST', '/api/**', 'init'),"
	sql += "(3, 'admin', 'PUT', '/api/**', 'init'),"
	sql += "(4, 'admin', 'DELETE', '/api/**', 'init'),"
	sql += "(5, 'admin', 'PATCH', '/api/**', 'init');"

	return
}

func CreateAuthorizedAPITableDataPGSql() (sql string) {
	sql = "INSERT INTO authorized_api (id, business_key, method, api,created_user) VALUES"
	sql += "(1, 'admin', 'GET', '/api/**', 'init'),"
	sql += "(2, 'admin', 'POST', '/api/**', 'init'),"
	sql += "(3, 'admin', 'PUT', '/api/**', 'init'),"
	sql += "(4, 'admin', 'DELETE', '/api/**', 'init'),"
	sql += "(5, 'admin', 'PATCH', '/api/**', 'init');"

	return
}
