drop table if exists dives;
drop table if exists competitions;
drop table if exists alternate_email;
alter table clubs drop constraint fk_club_main_contact
drop table if exists users;
drop table if exists clubs;

create table clubs (
    id serial primary key,
    name varchar(128),
    contact_info varchar(2048),
    main_contact_id integer
);

create table users (
    id serial primary key,
    club_id integer references clubs(id),
    email varchar(128) unique,
    firstnama varchar(128),
    lastname varchar(128),
    birthday date
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

create table competitions (
    id serial primary key,
    name varchar(128),
    start_date date,
    end_date date
);

create table dives (
    id serial primary key,
    user_id integer references users(id),
    competition_id integer references competitions(id),
    scores numeric(1,1)[],
    video_reference varchar(256),
    user_comment text
);