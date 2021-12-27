create table todos
(
    id SERIAL not null,
    title text not null,
    done boolean,
    user_id integer not null,
    PRIMARY KEY (id)
);

create table users
(
    id SERIAL not null,
    name varchar not null,
    PRIMARY KEY (id)
);