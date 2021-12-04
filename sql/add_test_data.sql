insert into clubs (name, contact_info)
values ('Bergen Stupeklubb', 'Vi holder til på AdO')
;

insert into users (email, firstnama, lastname, club_id, birthday)
values ('terje@bergesen.info', 'Terje Anthon', 'Bergesen', 1, '1966-10-29')
;

insert into users (email, firstnama, lastname, club_id, birthday)
values ('andrea@bergesen.info', 'Andrea Færøyvik', 'Bergesen', 1, '2011-04-11')
;

insert into alternate_email (user_id, email)
values ( 1, 'terjeb@yahoo.com' )
;

insert into alternate_email (user_id, email)
values ( 1, 'terje.a.bergesen@gmail.com' )
;
