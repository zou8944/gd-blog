create table blog
(
    id            integer not null primary key autoincrement,
    title         text    not null,
    summary       text    not null,
    content       text    not null,
    like_count    integer not null default 0,
    collect_count integer not null default 0,
    scores        text    not null default '',
    created_at    integer not null default CURRENT_TIMESTAMP,
    updated_at    integer not null default CURRENT_TIMESTAMP,
    deleted_at    integer
);
create table category
(
    id          integer not null primary key autoincrement,
    name        text    not null,
    description text,
    created_at  integer not null default CURRENT_TIMESTAMP,
    updated_at  integer not null default CURRENT_TIMESTAMP,
    deleted_at  integer

);
create table blog_category
(
    blog_id     integer not null,
    category_id integer not null,
    created_at  integer not null default CURRENT_TIMESTAMP
);
create table label
(
    id         integer not null primary key autoincrement,
    name       text    not null,
    created_at integer not null default CURRENT_TIMESTAMP,
    updated_at integer not null default CURRENT_TIMESTAMP,
    deleted_at integer
);
create table blog_label
(
    blog_id    integer not null,
    label_id   integer not null,
    created_at integer not null default CURRENT_TIMESTAMP
);
create table visitor
(
    id         integer not null primary key autoincrement,
    name       text    not null,
    email      text    not null,
    created_at integer not null default CURRENT_TIMESTAMP,
    updated_at integer not null default CURRENT_TIMESTAMP,
    deleted_at integer
);
create table comment
(
    id         integer not null primary key autoincrement,
    blog_id    integer not null,
    visitor_id integer not null,
    reply_id   integer,
    g          integer not null,
    content    text    not null,
    approved   integer not null default 0,
    created_at integer not null default CURRENT_TIMESTAMP,
    updated_at integer not null default CURRENT_TIMESTAMP,
    deleted_at integer
);