CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT  timestamp '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "files_owned" varchar[] 
);

CREATE TABLE "files" (
    "id" varchar PRIMARY KEY,
    "owner" varchar NOT NULL,
    "s3_url" varchar NOT NULL ,
    "file_name" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())

);
CREATE INDEX ON "files" ("owner");

ALTER TABLE "files" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

