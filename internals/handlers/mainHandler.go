package handlers

import (
	dbManager "comment/internals/dbManager"
	md "comment/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbManager.InitDB()
	if err != nil {
		log.Println("db not opening !", err)
		http.Error(w, "database can't be opened", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var req md.Request
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("req.Action : ", req.Action)

	switch req.Action {
	case "createComment":
		createComment(w, req.Body, db)
	case "delete":
		deleteComment(w, req.Body, db)
	case "getAllPostComment":
		getAllPostComment(w, req.Body, db)
	case "updateLike":
		updateLike(w, req.Body, db)
	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

}
