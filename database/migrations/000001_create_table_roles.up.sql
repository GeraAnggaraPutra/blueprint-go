CREATE TABLE IF NOT EXISTS "roles" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

INSERT INTO roles(guid, name)
VALUES
(gen_random_uuid(), 'Super Admin'),
(gen_random_uuid(), 'Admin')
ON CONFLICT DO NOTHING;
