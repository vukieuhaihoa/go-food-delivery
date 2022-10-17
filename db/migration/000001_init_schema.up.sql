CREATE TABLE "restaurants" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "address" varchar(250) NOT NULL,
  -- "city_id" bigint NOT NULL,
  -- "owner_id" bigint NOT NULL,
  -- "lat" real NOT NULL,
  -- "lng" real NOT NULL,
  -- "cover" json NOT NULL,
  -- "logo" json NOT NULL,
  "city_id" bigint,
  "owner_id" bigint,
  "lat" real,
  "lng" real,
  "cover" json,
  "logo" json,
  "shiping_fee_per_km" real DEFAULT 0,
  "status" smallint NOT NULL DEFAULT 1,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "restaurants" ("name");

CREATE INDEX ON "restaurants" ("owner_id");

CREATE INDEX ON "restaurants" ("city_id");

CREATE INDEX ON "restaurants" ("status");
