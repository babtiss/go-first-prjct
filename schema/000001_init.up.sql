CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_item_data
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
);

CREATE TABLE todo_items
(
    id               serial  not null unique,
    todo_item_dataID int     not null unique,
    foreign key (todo_item_dataID) references todo_item_data (id) ON DELETE CASCADE ,
    done             boolean not null default false
);


CREATE TABLE lists_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);