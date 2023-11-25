BEGIN;


ALTER TABLE "public"."users"
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at,
DROP COLUMN IF EXISTS deleted_at;

COMMIT;