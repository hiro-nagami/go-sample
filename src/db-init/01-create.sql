create table todos
(
    id integer not null,
    title text not null,
    done boolean,
    user_id integer not null,
    PRIMARY KEY (id)
);

create table users
(
    id integer not null,
    name varchar not null,
    PRIMARY KEY (id)
);