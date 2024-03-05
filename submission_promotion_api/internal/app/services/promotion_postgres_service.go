package services

import (
	"errors"
)

type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

func NewPromotionService(PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
	return s.PromotionRepo.CreatePromotion(promo)
}

func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.PromotionRepo.GetAllPromotions()
}

func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	// Implementasi kamu taruh disini
	promo, err := s.PromotionRepo.GetPromotionbyPromotionID(promotionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.PromotionIDNotFoundError{
				Message:     "Promotion Not Found",
				PromotionID: promotionID,
			}
		}
		return models.Promotion{}, err
	}
	return promo, nil
}

func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
	updatePromo, err := s.PromotionRepo.UpdatePromotionbyPromotionID(promo)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.PromotionIDNotFoundError{
				Message:     "Duplicate Promotion Found",
				PromotionID: promo.PromotionID,
			}
		}
		return models.Promotion{}, err
	}
	return updatePromo, nil
}

func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	// Implementasi kamu taruh disini
	return s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
}
