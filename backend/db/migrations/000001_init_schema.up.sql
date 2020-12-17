CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
      "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
      "username" varchar NOT NULL,
      "email" varchar NOT NULL,
      "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "files" (
      "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
      "filename" varchar NOT NULL,
      "filesize" int NOT NULL,
      "fileextension" varchar,
      "uploaded_at" timestamp DEFAULT (now()),
      "uploaded_by" uuid
);

CREATE TABLE "reports" (
      "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
      "file" uuid,
      "malicious" boolean NOT NULL
);

CREATE TABLE "apps" (
    "id" uuid PRIMARY KEY DEFAULT  uuid_generate_v4(),
    "name" varchar NOT NULL,
    "public_key" varchar NOT NULL UNIQUE,
    "encrypted_secret_key" varchar NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
);

ALTER TABLE "files" ADD FOREIGN KEY ("uploaded_by") REFERENCES "users" ("id");

ALTER TABLE "reports" ADD FOREIGN KEY ("file") REFERENCES "files" ("id");
