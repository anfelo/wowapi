-- Table Definition
CREATE SEQUENCE IF NOT EXISTS races_id_seq;
CREATE TABLE "races" (
    "id" int4 NOT NULL DEFAULT nextval('races_id_seq'::regclass),
    "name" varchar(255) NOT NULL,
    "is_allied" boolean NOT NULL,
    "created_at" timestamp(0),
    "updated_at" timestamp(0),
    PRIMARY KEY ("id")
);
