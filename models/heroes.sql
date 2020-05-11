-- Table Definition
CREATE SEQUENCE IF NOT EXISTS heroes_id_seq;
CREATE TABLE "heroes" (
    "id" int4 NOT NULL DEFAULT nextval('heroes_id_seq'::regclass),
    "name" varchar(255) NOT NULL,
    "title" varchar(255) NOT NULL,
    "faction" varchar(255) NOT NULL,
    "race" varchar(255) NOT NULL,
    "location" varchar(255) NOT NULL,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id")
);
