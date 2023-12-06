CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "image" bytea NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" BIGSERIAL PRIMARY KEY,
  "category_name" varchar NOT NULL
);

CREATE TABLE "posts" (
  "id" BIGSERIAL PRIMARY KEY,
  "image" bytea NOT NULL,
  "title" varchar NOT NULL,
  "post_category" bigint NOT NULL,
  "content" text NOT NULL,
  "time_for_read" int NOT NULL,
  "like_number" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "reviews" (
  "id" bigint PRIMARY KEY,
  "reviewer" varchar NOT NULL,
  "review_content" text NOT NULL,
  "post" bigint NOT NULL,
  "star_degree" float NOT NULL,
  "like_number" bigint NOT NULL DEFAULT 0,
  "unlike_number" bigint NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "categories" ("category_name");

CREATE INDEX ON "posts" ("title");

CREATE INDEX ON "reviews" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("post_category") REFERENCES "categories" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("reviewer") REFERENCES "users" ("username");

ALTER TABLE "reviews" ADD FOREIGN KEY ("post") REFERENCES "posts" ("id");
