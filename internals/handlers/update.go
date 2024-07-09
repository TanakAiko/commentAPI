package handlers

import (
	"comment/internals/tools"
	md "comment/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func updateLike(w http.ResponseWriter, comment md.Comment, db *sql.DB) {
	query := `
        UPDATE comments
        SET nbrLike = ?, nbrDislike = ?, likedBy = ?, dislikedBy = ?
        WHERE id = ?;
    `

	likedByJSON, err := json.Marshal(comment.LikedBy)
	if err != nil {
		fmt.Println("ERROR 0")
		http.Error(w, "Error while deleting comment : "+err.Error(), http.StatusInternalServerError)
		return
	}

	dislikedByJSON, err := json.Marshal(comment.DisLikedBy)
	if err != nil {
		fmt.Println("ERROR 0.5")
		http.Error(w, "Error while deleting comment : "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := db.Exec(query, comment.NbrLike, comment.NbrDislike, string(likedByJSON), string(dislikedByJSON), comment.Id)
	if err != nil {
		fmt.Println("ERROR 1")
		http.Error(w, "Error while deleting comment : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("ERROR 2")
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("ERROR 3")
		http.Error(w, "No comment found with ID: "+strconv.Itoa(comment.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, "Post well updated", http.StatusOK)
}
