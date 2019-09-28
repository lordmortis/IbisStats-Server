CREATE TYPE game_stat_type AS ENUM ('score', 'time');
CREATE TABLE "game_stats" (
 id UUID NOT NULL PRIMARY KEY,
 game_id UUID NOT NULL references games(id),
 name text NOT NULL,
 type game_stat_type NOT NULL
);