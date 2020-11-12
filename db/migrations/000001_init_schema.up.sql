CREATE TABLE "users" (
      "id" bigserial PRIMARY KEY,
      "username" varchar NOT NULL,
      "email" varchar NOT NULL,
      "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "files" (
      "id" bigserial PRIMARY KEY,
      "filename" varchar NOT NULL,
      "filesize" int NOT NULL,
      "fileextension" varchar,
      "uploaded_at" timestamp DEFAULT (now()),
      "uploaded_by" bigint
);

CREATE TABLE "Reports" (
      "id" bigserial PRIMARY KEY,
      "file" bigint,
      "malicious" boolean NOT NULL
);

ALTER TABLE "files" ADD FOREIGN KEY ("uploaded_by") REFERENCES "users" ("id");

ALTER TABLE "Reports" ADD FOREIGN KEY ("file") REFERENCES "files" ("id");
