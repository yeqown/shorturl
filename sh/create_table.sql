use shorten_url;

create table shorted_url
(
    id      int auto_increment
        primary key,
    source  varchar(512) not null,
    hash    int          not null,
    shorted varchar(128) null,
    constraint shorted_url_hash_uindex
        unique (hash)
);