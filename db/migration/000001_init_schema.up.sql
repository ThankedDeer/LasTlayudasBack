CREATE TABLE
  "users" (
    "user_id" SERIAL PRIMARY KEY,
    "firstname" VARCHAR NOT NULL,
    "lastname" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "email" VARCHAR UNIQUE NOT NULL,
    "password_changed_at" TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" TIMESTAMP NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  "roles" (
    "role_id" SERIAL PRIMARY KEY,
    "name" VARCHAR UNIQUE NOT NULL,
    "description" TEXT
  );

CREATE TABLE
  "permissions" (
    "permission_id" SERIAL PRIMARY KEY,
    "name" VARCHAR UNIQUE NOT NULL,
    "description" TEXT
  );

CREATE TABLE
  "user_roles" (
    "user_id" INTEGER NOT NULL,
    "role_id" INTEGER NOT NULL,
    "assigned_at" TIMESTAMP NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  "role_permissions" (
    "role_id" INTEGER NOT NULL,
    "permission_id" INTEGER NOT NULL,
    "assigned_at" TIMESTAMP NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  "testimonials" (
    "testimonial_id" SERIAL PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "testimonial" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now ()),
    "updated_at" TIMESTAMP NOT NULL DEFAULT (now ()),
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
  );

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("permission_id");