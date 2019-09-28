CREATE TABLE "games" (
     id UUID NOT NULL PRIMARY KEY,
     name text NOT NULL,
     secret bytea NOT NULL,
     created_at timestamp,
     updated_at timestamp
);