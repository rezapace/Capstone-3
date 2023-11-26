BEGIN;

ALTER TABLE "public"."users"
ADD COLUMN created_at timestamptz (6),
ADD COLUMN updated_at timestamptz (6),
ADD COLUMN deleted_at timestamptz (6);

COMMIT;