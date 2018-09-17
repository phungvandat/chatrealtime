-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  "password" text NOT NULL,
  "phone_number" text DEFAULT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "updated_at" timestamptz,
  UNIQUE("name"),
  UNIQUE("phone_number"),
  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."users";