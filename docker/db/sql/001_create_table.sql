drop table if exists users;
create table IF not exists users (
	`id`           	int(20) auto_increment primary key,
	`display_name` 	varchar(50) not null,
	`full_name`    	varchar(50) not null,
	`email`       	varchar(255) default null,
	`card_no`     	varchar(15) default null
) default CHARSET=utf8 COLLATE=utf8_bin;

