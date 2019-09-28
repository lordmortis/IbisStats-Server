CREATE TABLE "players" (
 id UUID NOT NULL PRIMARY KEY,
 emailhash bytea UNIQUE NOT NULL,
 name text
);