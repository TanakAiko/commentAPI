package handlers

import (
	"comment/internals/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	md "comment/models"
)

func getAllComment(w http.ResponseWriter, db *sql.DB) {
	rows, err := db.Query("SELECT id, postId, userId, nickname, likedBy, dislikedBy, content, nbrLike, nbrDislike, createdAt FROM comments ORDER BY createdAt DESC")
	if err != nil {
		http.Error(w, "Error while getting comments : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	comments := []md.Comment{}
	for rows.Next() {
		var comment md.Comment
		var likedByJSON string
		var dislikedByJSON string
		if err := rows.Scan(&comment.Id, comment.PostId, &comment.UserId, &comment.Nickname, &likedByJSON, &dislikedByJSON, &comment.Content, &comment.NbrLike, &comment.NbrDislike, &comment.CreateAt); err != nil {
			fmt.Println("ERROR 1")
			http.Error(w, "Error while getting comments : "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal([]byte(likedByJSON), &comment.LikedBy); err != nil {
			fmt.Println("ERROR 3")
			http.Error(w, "Error while getting comments : "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal([]byte(dislikedByJSON), &comment.DisLikedBy); err != nil {
			fmt.Println("ERROR 4")
			http.Error(w, "Error while getting comments : "+err.Error(), http.StatusInternalServerError)
			return
		}

		comments = append(comments, comment)

	}

	tools.WriteResponse(w, comments, http.StatusOK)
}
