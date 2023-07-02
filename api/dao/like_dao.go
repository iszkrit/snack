package dao

import (
	"db/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func FetchLike() ([]model.Like, error) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM likes")

	var like model.Like
	var likes []model.Like

	for rows.Next() {
		err := rows.Scan(&like.LikeId, &like.MessageId, &like.AccountId)
		if err != nil {
			fmt.Print(err)
			rows.Close()
			break
		}
		likes = append(likes, like)
	}
	defer rows.Close()
	return likes, err
}

func PushLike(like model.Like) error {
	db := Connect()
	defer db.Close()
	_, err := db.Exec("INSERT INTO likes (likeId, messageId, accountId) VALUES (?,?,?)", like.LikeId, like.MessageId, like.AccountId)
	return err
}

func UndoLike(id string) error {
	db := Connect()
	defer db.Close()
	_, err := db.Exec("DELETE FROM likes WHERE likeId = ?", id)
	return err
}
