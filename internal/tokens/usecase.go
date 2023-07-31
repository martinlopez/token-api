package tokens

import (
	"context"
	"errors"
	"time"
	"tokens-api/pkg/handler"

	"gorm.io/gorm"
)

type usecase struct {
	//ifps client
	repo *gorm.DB
}

func NewUseCase(repo *gorm.DB) *usecase {
	return &usecase{repo}
}

//byID returns the entity with the given id
func (s *usecase) ByID(_ context.Context, id string) (any, error) {
	var ent Tokens
	result := s.repo.First(&ent, "id = ?", id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, handler.ErrNotFound
	}
	return ent, result.Error
}

//all returns all the entities based on paramters
func (s *usecase) All(_ context.Context, q map[string]any) (any, error) {
	var ent []Tokens
	result := s.repo.Where(q).Find(&ent)
	return ent, result.Error
}

func (s *usecase) Create(_ context.Context, data any) (any, error) {
	now := time.Now()
	ent := data.(Tokens)
	ent.CreatedAt = &now
	ent.UpdatedAt = &now
	result := s.repo.Create(&ent)
	return ent, result.Error
}

//update updates the entity with the given id
func (s *usecase) Update(_ context.Context, id string, ent any) (any, error) {
	return nil, errors.New("not implemented")
}

//delete deletes the entity with the given id
func (s *usecase) Delete(_ context.Context, id string) error {
	err := s.repo.Delete(&Tokens{}, "id = ?", id)
	return err.Error
}
