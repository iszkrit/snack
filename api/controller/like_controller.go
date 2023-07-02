package controller

import (
	"db/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchLike(w http.ResponseWriter) {

	likes, err := usecase.FetchLike()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Something went wrong when searching likes.\n")
		fmt.Print(err)
		return
	}

	likesJSON, err := json.MarshalIndent(likes, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Something went wrong.\n")
		fmt.Print(err)
	}

	fmt.Fprint(w, string(likesJSON))
}

func PushLike(w http.ResponseWriter, r *http.Request) {

	type NewLikeInfo struct {
		MessageId string `json:"messageId"`
		AccountId string `json:"accountId"`
	}
	var newLikeInfo NewLikeInfo

	if err := json.NewDecoder(r.Body).Decode(&newLikeInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Something wrong. \n")
		fmt.Print(err)
		return
	}

	if err := usecase.PushLike(newLikeInfo.MessageId, newLikeInfo.AccountId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Cannot register the user.\n")
		fmt.Print(err)
		return
	}

	w.Write([]byte("201 Created"))
}

func UndoLike(w http.ResponseWriter, r *http.Request) {

	type DeleteId struct {
		Id string `json:"id"`
	}
	var deleteId DeleteId

	if err := json.NewDecoder(r.Body).Decode(&deleteId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Something wrong. \n")
		fmt.Print(err)
		return
	}

	if err := usecase.UndoLike(deleteId.Id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Cannot register the user.\n")
		fmt.Print(err)
		return
	}

	w.Write([]byte("200 OK"))
}
