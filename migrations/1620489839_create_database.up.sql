CREATE TABLE "public"."users" (
    "id" uuid unique NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "display_name" text,
    "handle" text unique,
    "mail" text,
    "password" text,
    PRIMARY KEY ("handle")
);
