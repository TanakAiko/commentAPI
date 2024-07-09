package handlers

import (
	"comment/internals/tools"
	"database/sql"
	"net/http"
	"strconv"

	md "comment/models"
)

func deleteComment(w http.ResponseWriter, comment md.Comment, db *sql.DB) {
	result, err := db.Exec("DELETE FROM comments WHERE id = ?", comment.Id)
	if err != nil {
		http.Error(w, "Error while deleting comment : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No comment found with ID: "+strconv.Itoa(comment.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, "Comment well deleted", http.StatusOK)
}
