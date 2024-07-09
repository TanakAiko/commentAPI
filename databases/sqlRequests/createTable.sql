CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    postId INTEGER NOT NULL,
    userId INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    likedBy TEXT NOT NULL,
    dislikedBy TEXT NOT NULL,
    content TEXT NOT NULL,
    nbrLike INTEGER,
    nbrDislike INTEGER,
    createdAt DATETIME NOT NULL
);