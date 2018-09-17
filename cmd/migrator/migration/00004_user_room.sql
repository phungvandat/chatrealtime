-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."user_rooms" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL REFERENCES users,
  "room_id" uuid NOT NULL REFERENCES rooms,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "updated_at" timestamptz,
  CONSTRAINT "user_rooms_pkey" PRIMARY KEY ("id")
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."user_rooms";