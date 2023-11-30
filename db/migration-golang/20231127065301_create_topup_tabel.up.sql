BEGIN;
CREATE TABLE IF NOT EXISTS "public"."topup" (
    "id" varchar(255) NOT NULL PRIMARY KEY,
    "user_id" integer NOT NULL,
    "amount" integer NOT NULL,
    "status" integer default 0,
    "snap_url" varchar(255) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
COMMIT;