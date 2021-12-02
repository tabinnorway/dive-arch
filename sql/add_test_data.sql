insert into clubs (name, contact_info)
values ('Bergen Stupeklubb', 'Vi holder til på AdO')
;

insert into users (email, firstnama, lastname, club_id)
values ('terje@bergesen.info', 'Terje', 'Bergesen', 1)
;

insert into alternate_email (user_id, email)
values ( 1, 'terjeb@yahoo.com' )
;

insert into alternate_email (user_id, email)
values ( 1, 'terje.a.bergesen@gmail.com' )
;
