package handlers

import (
	"database/sql"
	"net/http"

	"comment/internals/tools"
	md "comment/models"
)

func createComment(w http.ResponseWriter, post md.Comment, db *sql.DB) {
	if err := post.CreateComment(db); err != nil {
		http.Error(w, "Error while creating post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteResponse(w, "New post created", http.StatusCreated)
}
