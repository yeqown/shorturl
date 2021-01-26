use shorten-url;

create table shorted_url (
	id int(64) not null primary key auto_increment,
	source varchar(512) not null,
	shorted varchar(128)
);