package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/enums"
)

type AppSettingRepository interface {
	GetAll() ([]*models.AppSetting, error)
	GetByKey(key enums.AppSettingKey) (*models.AppSetting, error)
	Set(key enums.AppSettingKey, value string) error
	Delete(key enums.AppSettingKey) error
}

type appSettingRepository struct {
	db *config.DB
}

func NewAppSettingRepository(
	db *config.DB,
) AppSettingRepository {
	return &appSettingRepository{
		db: db,
	}
}

func (r *appSettingRepository) GetAll() ([]*models.AppSetting, error) {
	var appSettings []*models.AppSetting
	if err := r.db.Find(&appSettings).Error; err != nil {
		return nil, err
	}

	return appSettings, nil
}

func (r *appSettingRepository) GetByKey(key enums.AppSettingKey) (*models.AppSetting, error) {
	var appSetting models.AppSetting
	if err := r.db.Where("key = ?", key).First(&appSetting).Error; err != nil {
		return nil, err
	}

	return &appSetting, nil
}

func (r *appSettingRepository) Set(key enums.AppSettingKey, value string) error {
	appSetting := models.AppSetting{Key: string(key), Value: value}
	if err := r.db.Save(&appSetting).Error; err != nil {
		return err
	}

	return nil
}

func (r *appSettingRepository) Delete(key enums.AppSettingKey) error {
	if err := r.db.Where("key = ?", key).Delete(&models.AppSetting{}).Error; err != nil {
		return err
	}

	return nil
}
