CREATE TABLE "public"."permissions" (
      "id" uuid unique NOT NULL,
      "created_at" timestamptz,
      "updated_at" timestamptz,
      "deleted_at" timestamptz,
      "users_id" uuid,
      "permission" text,
      "flag" int,
      PRIMARY KEY ("id")
);

ALTER TABLE "public"."permissions" ADD FOREIGN KEY ("users_id") REFERENCES "public"."users"("id");
