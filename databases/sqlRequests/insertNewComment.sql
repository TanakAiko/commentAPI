INSERT INTO comments (
        postId,
        userId,
        nickname,
        likedBy,
        dislikedBy,
        content,
        nbrLike,
        nbrDislike,
        createdAt
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)