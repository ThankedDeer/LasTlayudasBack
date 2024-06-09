CREATE TABLE
  "users" (
    "user_id" SERIAL PRIMARY KEY,
    "firstname" VARCHAR NOT NULL,
    "lastname" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "email" VARCHAR UNIQUE NOT NULL,
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

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("permission_id");