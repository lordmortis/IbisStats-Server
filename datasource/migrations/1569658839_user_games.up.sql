CREATE TABLE "user_game" (
    user_id UUID NOT NULL references users(id),
    game_id UUID NOT NULL references games(id),
    PRIMARY KEY(user_id, game_id)
);