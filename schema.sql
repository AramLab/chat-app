create table users (
    id integer primary key,
    username text not null,
    password text not null
);

create table channels (
    id integer primary key,
    name text not null
);

create table messages (
    id integer primary key,
    channel_id integer not null,
    user_id integer not null,
    message text not null,
    created_at timestamp default CURRENT_TIMESTAMP
);