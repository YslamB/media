

insert into languages (name) values ('tm');
insert into languages (name) values ('ru');


insert into admins (username, password) values ('admin', '$2a$10$WQlvtdWAZJBMIasjD79tHOIoP051TOpnn9C/7g1pHFpnKCyuSppke');

insert into categories (name) 
    values ('{"ru":"Фильмы","tm":"Kinolar"}'),
           ('{"ru":"Книги","tm":"Kitaplar"}'),
           ('{"ru":"Музыка","tm":"Aÿdymlar"}');

insert into sub_categories (category_id, name)
    values (1,'{"ru":"Драма","tm":"Dramma"}'),
           (1,'{"ru":"Комедия","tm":"Komediya"}'),
           (1,'{"ru":"Триллер","tm":"Triller"}'),
           (1,'{"ru":"Фантастика","tm":"Fantastika"}'),
           (1,'{"ru":"Мелодрама","tm":"Melodramma"}'),
           (1,'{"ru":"Мистика","tm":"Mistika"}');

insert into sub_categories (category_id, name)
    values (2,'{"ru":"Драма","tm":"Dramma"}'),
           (2,'{"ru":"Комедия","tm":"Komediya"}'),
           (2,'{"ru":"Триллер","tm":"Triller"}'),
           (2,'{"ru":"Фантастика","tm":"Fantastika"}'),
           (2,'{"ru":"Мелодрама","tm":"Melodramma"}'),
           (2,'{"ru":"Мистика","tm":"Mistika"}');

insert into sub_categories (category_id, name)
    values (3,'{"ru":"Драма","tm":"Dramma"}'),
           (3,'{"ru":"Комедия","tm":"Komediya"}'),
           (3,'{"ru":"Триллер","tm":"Triller"}'),
           (3,'{"ru":"Фантастика","tm":"Fantastika"}'),
           (3,'{"ru":"Мелодрама","tm":"Melodramma"}'),
           (3,'{"ru":"Мистика","tm":"Mistika"}');

