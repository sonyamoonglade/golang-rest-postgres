-- noinspection SqlNoDataSourceInspectionForFile


CREATE TABLE "users"(
    "user_id" serial not null primary key,
    "name" varchar(255) not null,
    "worth" real not null
);

CREATE TABLE "car" (
    "car_id" serial not null primary key,
    "model" varchar(255) not null,
    "year" integer not null,
    "price" real not null
);

CREATE TABLE "garage"(
    "garage_id" serial not null primary key,
    "user_id" integer not null,
    "name" varchar(255) not null,
    "capacity" integer not null,
    "space_left" integer not null,
    CONSTRAINT user_id_fk
        foreign key(user_id)
        references users(user_id)
        on delete cascade
);

CREATE TABLE "garage_car"(
    "car_id" integer not null,
    "garage_id" integer not null,
    "user_id"  integer not null,
    CONSTRAINT "car_id_fk"
        foreign key(car_id)
        references car(car_id)
        on delete cascade,
    CONSTRAINT "garage_id_fk"
        foreign key(garage_id)
        references garage(garage_id)
        on delete cascade,
    CONSTRAINT "user_id_fk"
        foreign key(user_id)
        references users(user_id)
        on delete cascade
);

CREATE INDEX ON "garage_car" ("car_id","garage_id","user_id");

CREATE INDEX ON "garage" ("user_id");
