-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."message_rooms" (
  "id" uuid NOT NULL,
  "user_name_send" text NOT NULL,
  "user_id_send" uuid NOT NULL REFERENCES users,
  "room_name_receive" text NOT NULL,
  "room_id_receive" uuid NOT NULL REFERENCES rooms,
  "body" text DEFAULT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "updated_at" timestamptz,
  CONSTRAINT "message_rooms_pkey" PRIMARY KEY ("id")
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."message_rooms";