-- Table Definition
CREATE SEQUENCE IF NOT EXISTS classes_id_seq;
CREATE TABLE "classes" (
    "id" int4 NOT NULL DEFAULT nextval('classes_id_seq'::regclass),
    "name" varchar(255) NOT NULL,
    "roles" jsonb NOT NULL DEFAULT '{}'::jsonb,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id")
);

INSERT INTO classes(name, roles, updated_at, created_at) values('Warrior', '{"data": ["Tank", "Damage"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Paladin', '{"data": ["Tank", "Damage", "Healer"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Hunter', '{"data": ["Damage"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Rogue', '{"data": ["Damage"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Priest', '{"data": ["Damage", "Healer"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Shaman', '{"data": ["Damage", "Healer"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Mage', '{"data": ["Damage"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Warlock', '{"data": ["Damage"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Monk', '{"data": ["Tank", "Damage", "Healer"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Druid', '{"data": ["Tank", "Damage", "Healer"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Demon Hunter', '{"data": ["Tank", "Damage"]}', NOW(), NOW());
INSERT INTO classes(name, roles, updated_at, created_at) values('Death Knight', '{"data": ["Tank", "Damage"]}', NOW(), NOW());