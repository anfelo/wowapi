-- Table Definition
CREATE SEQUENCE IF NOT EXISTS factions_id_seq;
CREATE TABLE "factions" (
    "id" int4 NOT NULL DEFAULT nextval('factions_id_seq'::regclass),
    "name" varchar(255) NOT NULL,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id")
);

INSERT INTO factions(name, updated_at, created_at) values('Horde', NOW(), NOW());
INSERT INTO factions(name, updated_at, created_at) values('Alliance', NOW(), NOW());
INSERT INTO factions(name, updated_at, created_at) values('Burning Legion', NOW(), NOW());
INSERT INTO factions(name, updated_at, created_at) values('Illidari', NOW(), NOW());