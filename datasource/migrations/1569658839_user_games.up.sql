CREATE TYPE user_game_type AS ENUM ('owner', 'reader');
CREATE TABLE "user_game" (
    user_id UUID NOT NULL references users(id),
    game_id UUID NOT NULL references games(id),
    PRIMARY KEY(user_id, game_id),
    type user_game_type NOT NULL,
    created_at timestamp,
    updated_at timestamp
);