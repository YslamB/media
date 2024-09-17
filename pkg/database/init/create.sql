create table admins (
    "username" character varying(20) primary key ,
    "password" character varying(100) not null,
    "created_at" timestamp without time zone default now(),
    "updated_at" timestamp without time zone default now()
);

create table languages (
    "name" character varying(3) primary key,
    "created_at" timestamp without time zone default now()
);

create table categories (
    "id" serial primary key,
    "name" character varying(100),
    "created_at" timestamp without time zone default now()
);

create table sub_categories(
    "id" serial primary key,
    "category_id" integer,
    "name" character varying(100),
    "created_at" timestamp without time zone default now(),
    CONSTRAINT sub_categories_category_id_fk 
        FOREIGN KEY ("category_id") 
            REFERENCES categories ("id")
                on update cascade ON DELETE CASCADE
);

drop table if exists musics;
create table musics (
    "id" serial primary key not null,
    "sub_category_id" integer not null,
    "language_id" integer not null,
    "title" character varying(100) not null,
    "description" text not null,
    "path" character varying(100) not null,
    "image_path" character varying(100) not null,
    "created_at" timestamp without time zone default now(),
    CONSTRAINT musics_sub_category_id_fk 
        FOREIGN KEY ("sub_category_id") 
            REFERENCES sub_categories ("id")
                on update cascade ON DELETE CASCADE,
    CONSTRAINT musics_language_id_fk 
        FOREIGN KEY ("language_id") 
            REFERENCES languages ("id")
                on update cascade ON DELETE CASCADE
);

drop table if exists films;
create table films (
    "id" serial primary key not null,
    "sub_category_id" integer not null,
    "language_id" integer not null,
    "title" character varying(100) not null,
    "description" text not null,
    "path" character varying(100) not null,
    "image_path" character varying(100) not null,
    "created_at" timestamp without time zone default now(),
    CONSTRAINT films_sub_category_id_fk 
        FOREIGN KEY ("sub_category_id") 
            REFERENCES sub_categories ("id")
                on update cascade ON DELETE CASCADE,
    CONSTRAINT musics_language_id_fk 
        FOREIGN KEY ("language_id") 
            REFERENCES languages ("id")
                on update cascade ON DELETE CASCADE
);

drop table if exists books;
create table books (
    "id" serial primary key not null,
    "sub_category_id" integer not null,
    "language_id" integer not null,
    "title" character varying(100) not null,
    "description" text not null,
    "path" character varying(100) not null,
    "image_path" character varying(100) not null,
    "created_at" timestamp without time zone default now(),
    CONSTRAINT books_sub_category_id_fk 
        FOREIGN KEY ("sub_category_id") 
            REFERENCES sub_categories ("id")
                on update cascade ON DELETE CASCADE,
    CONSTRAINT musics_language_id_fk 
        FOREIGN KEY ("language_id") 
            REFERENCES languages ("id")
                on update cascade ON DELETE CASCADE
);
