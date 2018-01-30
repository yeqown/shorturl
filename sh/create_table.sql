use shorturl;

create table longurl(
	id int(64) not null primary key auto_increment,
	long_url varchar(100) not null,
	short_url varchar(40)
);