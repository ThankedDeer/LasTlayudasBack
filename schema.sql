CREATE TABLE "users" (
  "user_id" integer PRIMARY KEY NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
  "role_id" integer PRIMARY KEY NOT NULL,
  "name" varchar UNIQUE NOT NULL,
  "description" text
);

CREATE TABLE "permissions" (
  "permission_id" integer PRIMARY KEY NOT NULL,
  "name" varchar UNIQUE NOT NULL,
  "description" text
);

CREATE TABLE "user_roles" (
  "user_id" integer NOT NULL,
  "role_id" integer NOT NULL,
  "assigned_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "role_permissions" (
  "role_id" integer NOT NULL,
  "permission_id" integer NOT NULL,
  "assigned_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("permission_id");
