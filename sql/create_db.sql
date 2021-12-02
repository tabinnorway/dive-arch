drop table if exists alternate_email;
drop table if exists users;
top table if exists clubs;

create table clubs (
    id serial primary key,
    name varchar(128),
    contact_info varchar(2048),
    main_contact_id integer
);

create table users (
    id serial primary key,
    email varchar(128) unique,
    firstnama varchar(128),
    lastname varchar(128),
    club_id integer references clubs(id)
);

alter table clubs
    add constraint fk_club_main_contact
    foreign key (main_contact_id)
    references users(id);

create table alternate_email (
    id serial primary key,
    user_id integer references users(id),
    email varchar(128)
);
