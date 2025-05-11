package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/repositories"
	"encoding/json"
	"errors"
)

type AppSettingService interface {
	GetAll() ([]*models.AppSetting, error)
	GetByKey(key enums.AppSettingKey) (*models.AppSetting, error)
	Set(key enums.AppSettingKey, value interface{}) error
	Delete(key enums.AppSettingKey) error
}

type appSettingService struct {
	repo repositories.AppSettingRepository
}

func NewAppSettingService(
	repo repositories.AppSettingRepository,
) AppSettingService {
	return &appSettingService{
		repo: repo,
	}
}

func (s *appSettingService) GetAll() ([]*models.AppSetting, error) {
	return s.repo.GetAll()
}

func (s *appSettingService) GetByKey(key enums.AppSettingKey) (*models.AppSetting, error) {
	setting, err := s.repo.GetByKey(key)
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (s *appSettingService) Set(key enums.AppSettingKey, value interface{}) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return errors.New("failed to serialize value to JSON")
	}

	return s.repo.Set(key, string(jsonBytes))
}

func (s *appSettingService) Delete(key enums.AppSettingKey) error {
	return s.repo.Delete(key)
}

func GetSettingValue[T any](s AppSettingService, key enums.AppSettingKey) (*T, error) {
	setting, err := s.GetByKey(key)
	if err != nil || setting == nil {
		return nil, err
	}

	var value T
	if err = json.Unmarshal([]byte(setting.Value), &value); err != nil {
		return nil, err
	}

	return &value, nil
}
