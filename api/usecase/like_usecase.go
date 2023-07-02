package usecase

import (
	"db/dao"
	"db/model"
)

func FetchLike() ([]model.Like, error) {
	likes, err := dao.FetchLike()
	return likes, err
}

func PushLike(messageId, accountId string) error {

	likeId := model.CreateId()

	newLike := model.Like{
		LikeId:    likeId,
		MessageId: messageId,
		AccountId: accountId,
	}

	err := dao.PushLike(newLike)

	return err
}

func UndoLike(id string) error {
	err := dao.UndoLike(id)
	return err
}
