drop table if exists users;
create table if not exists users (
	`card_no`     		char(15) not null primary key unique,
	`display_name` 		varchar(50) default null,
	`full_name`    		varchar(50) default null,
	`pronunciation` 	varchar(100) default null,
	`playlist` 			varchar(100) default null,
	`email`       		varchar(255) default null
) default CHARSET=utf8 COLLATE=utf8_bin;

drop table if exists current_users;
create table if not exists current_users (
	`card_no`     	char(15) not null primary key unique,
	`date_touched` 	datetime null default current_timestamp,
	index idx_current(card_no),
	foreign key fk_current(card_no) references users(card_no)
) default CHARSET=utf8 COLLATE=utf8_bin;

