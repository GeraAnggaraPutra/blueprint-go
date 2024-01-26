CREATE TABLE IF NOT EXISTS "sessions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL,
  "user_guid" varchar NOT NULL,
  "ip_address" varchar NOT NULL,
  "token" varchar,
  "user_agent" varchar NOT NULL,
  "expired_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now())
);
