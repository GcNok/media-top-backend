package auth

import (
	"context"

	"github.com/davecgh/go-spew/spew"

	dbRepo "github.com/shiji-naoki/media-top-appli-back/domain/repository/database"
)

type (
	// LoginUseCase このクラスで定義するメソッドを定義
	LoginUseCase interface {
		Do(ctx context.Context, id string) (name string)
	}

	// このクラスで使用するクラスを定義
	loginUseCase struct {
		loginRepository dbRepo.LoginRepository
	}
)

// NewLoginUseCase このクラスのインスタンスを返す
func NewLoginUseCase(lr dbRepo.LoginRepository) LoginUseCase {
	return &loginUseCase{loginRepository: lr}
}

// このクラスのメイン処理
func (l *loginUseCase) Do(ctx context.Context, id string) (name string) {
	// /registry/repostory.goで紐づけたloginRepositoryのGetUserNameを呼び出す
	admin := l.loginRepository.GetUserName(id)
	spew.Dump(admin)
	return admin[0].Name
}
