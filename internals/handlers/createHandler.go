package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"comment/internals/tools"
	md "comment/models"
)

func createComment(w http.ResponseWriter, comment md.Comment, db *sql.DB) {
	fmt.Println("comment: ", comment)
	if err := comment.CreateComment(db); err != nil {
		fmt.Println("ERROR : ", err.Error())
		http.Error(w, "Error while creating comment : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteResponse(w, "New comment created", http.StatusCreated)
}
