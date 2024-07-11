package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Comment struct {
	Id         int       `json:"commentID"`
	PostId     int       `json:"postID"`
	UserId     int       `json:"userID"`
	Nickname   string    `json:"nickname"`
	LikedBy    []string  `json:"likedBy"`
	DisLikedBy []string  `json:"dislikedBy"`
	Content    string    `json:"content"`
	NbrLike    int       `json:"nbrLike"`
	NbrDislike int       `json:"nbrDislike"`
	CreateAt   time.Time `json:"createAt"`
}

func (comment *Comment) CreateComment(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewComment.sql")
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(string(content))
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		comment.PostId,
		comment.UserId,
		comment.Nickname,
		"",
		"",
		comment.Content,
		comment.NbrLike,
		comment.NbrDislike,
		time.Now().Format(time.RFC3339),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}

	return err
}
