CREATE TABLE "users" (
     "id" SERIAL PRIMARY KEY,
     "first_name" varchar,
     "last_name" varchar,
     "email" varchar,
     "password" varchar,
     "access_level" int,
     "created_at" timestamp,
     "updated_at" timestamp
);

CREATE TABLE "reservations" (
    "id" SERIAL PRIMARY KEY,
    "first_name" varchar,
    "last_name" varchar,
    "email" varchar,
    "phone" varchar,
    "start_date" date,
    "end_date" date,
    "room_id" int,
    "created_at" timestamp,
    "updated_at" timestamp
);

CREATE TABLE "rooms" (
     "id" SERIAL PRIMARY KEY,
     "room_name" varchar,
     "created_at" timestamp,
     "updated_at" timestamp
);

CREATE TABLE "room_restrictions" (
     "id" SERIAL PRIMARY KEY,
     "start_date" date,
     "end_date" date,
     "room_id" int,
     "reservation_id" int,
     "restriction_id" int,
     "created_at" timestamp,
     "updated_at" timestamp
);

CREATE TABLE "restrictions" (
    "id" SERIAL PRIMARY KEY,
    "restriction_name" varchar,
    "created_at" timestamp,
    "updated_at" timestamp
);

ALTER TABLE "reservations" ADD CONSTRAINT room_fk FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;

ALTER TABLE "room_restrictions" ADD CONSTRAINT room_fk FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;

ALTER TABLE "room_restrictions" ADD CONSTRAINT reservation_fk FOREIGN KEY ("reservation_id") REFERENCES "reservations" ("id") ON DELETE CASCADE;

ALTER TABLE "room_restrictions" ADD CONSTRAINT restriction_fk FOREIGN KEY ("restriction_id") REFERENCES "restrictions" ("id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "users" ("email");

CREATE INDEX ON "room_restrictions" ("start_date");
CREATE INDEX ON "room_restrictions" ("end_date");
CREATE INDEX ON "room_restrictions" ("room_id");
CREATE INDEX ON "room_restrictions" ("reservation_id");
CREATE INDEX ON "reservations" ("email");
CREATE INDEX ON "reservations" ("last_name");