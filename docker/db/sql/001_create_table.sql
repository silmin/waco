drop table if exists users;
create table if not exists users (
	`id`			int	not null primary key auto_increment,
	`card_no`     	varchar(15) not null unique,
	`display_name` 	varchar(50) not null,
	`full_name`    	varchar(50) not null,
	`email`       	varchar(255) default null
) default CHARSET=utf8 COLLATE=utf8_bin;

