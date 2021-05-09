CREATE TABLE "public"."users" (
    "id" uuid NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "display_name" text,
    "handle" uuid,
    "mail" bool,
    PRIMARY KEY ("id")
);