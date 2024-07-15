package handlers

import (
	"comment/internals/tools"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	md "comment/models"
)

func getLastComment(w http.ResponseWriter, db *sql.DB) {
	var comment md.Comment
	var likedByJSON string
	var dislikedByJSON string
	query := "SELECT id, postId, userId, nickname, likedBy, dislikedBy, content, nbrLike, nbrDislike, createdAt FROM comments ORDER BY id DESC LIMIT 1"
	row := db.QueryRow(query)
	err := row.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Nickname, &likedByJSON, &dislikedByJSON, &comment.Content, &comment.NbrLike, &comment.NbrDislike, &comment.CreateAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal([]byte(likedByJSON), &comment.LikedBy); err != nil {
		fmt.Println("Error while unmarshaling likedBy: " + err.Error())
	}

	if err := json.Unmarshal([]byte(dislikedByJSON), &comment.DisLikedBy); err != nil {
		fmt.Println("Error while unmarshaling dislikedBy: " + err.Error())
	}

	if err := row.Err(); err != nil {
		http.Error(w, "Error while iterating comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, comment, http.StatusOK)
}

func getAllPostComment(w http.ResponseWriter, commentR md.Comment, db *sql.DB) {

	// Préparer la requête avec context pour un meilleur contrôle
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
        SELECT id, postId, userId, nickname, likedBy, dislikedBy, content, nbrLike, nbrDislike, createdAt
        FROM comments
        WHERE postId = ?
        ORDER BY createdAt DESC
    `
	rows, err := db.QueryContext(ctx, query, commentR.PostId)
	if err != nil {
		http.Error(w, "Error while getting comments: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			http.Error(w, "Error while closing rows: "+err.Error(), http.StatusInternalServerError)
		}
	}()

	var comments []md.Comment
	for rows.Next() {
		var comment md.Comment
		var likedByJSON, dislikedByJSON string
		if err := rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Nickname, &likedByJSON, &dislikedByJSON, &comment.Content, &comment.NbrLike, &comment.NbrDislike, &comment.CreateAt); err != nil {
			fmt.Println("ERROR 1")
			http.Error(w, "Error while scanning comments: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("likedByJSON: ", likedByJSON)
		if err := json.Unmarshal([]byte(likedByJSON), &comment.LikedBy); err != nil {
			fmt.Println("ERROR 2")
			fmt.Println("Error while unmarshaling likedBy: " + err.Error())
		}

		fmt.Println("dislikedByJSON: ", dislikedByJSON)
		if err := json.Unmarshal([]byte(dislikedByJSON), &comment.DisLikedBy); err != nil {
			fmt.Println("ERROR 3")
			fmt.Println("Error while unmarshaling dislikedBy: " + err.Error())
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("ERROR 4")
		http.Error(w, "Error while iterating comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Utiliser un logger approprié pour des logs détaillés
	fmt.Printf("Fetched comments: %+v\n", comments)

	// Utiliser des outils appropriés pour l'écriture de la réponse
	tools.WriteResponse(w, comments, http.StatusOK)
}
