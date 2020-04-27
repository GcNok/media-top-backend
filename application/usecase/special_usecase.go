package usecase

import (
	"context"

	"github.com/davecgh/go-spew/spew"

	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type (
	// SpecialUseCase このクラスで定義するメソッドを定義
	SpecialUseCase interface {
		Do(ctx context.Context) (articles []dbEntity.Special)
	}

	// このクラスで使用するクラスを定義
	specialUseCase struct {
		specialRepository dbRepo.SpecialRepository
	}
)

// NewSpecialUseCase このクラスのインスタンスを返す
func NewSpecialUseCase(lr dbRepo.SpecialRepository) SpecialUseCase {
	return &specialUseCase{specialRepository: lr}
}

// このクラスのメイン処理
func (l *specialUseCase) Do(ctx context.Context) (articles []dbEntity.Special) {
	specials := l.specialRepository.GetArticle()
	spew.Dump(specials)
	return specials
}
