package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/mgr1054/myaktion-go/src/myaktion/src/myaktion/db"
	"github.com/mgr1054/myaktion-go/src/myaktion/src/myaktion/model"
)

func CreateCampaign(campaign *model.Campaign) error {
	result := db.DB.Create(campaign)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Tracef("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	result := db.DB.Preload("Donations").Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaigns)
	return campaigns, nil
}

func GetCampaign(id uint) (*model.Campaign, error) {
	campaign := new(model.Campaign)
	result := db.DB.Preload("Donations").First(campaign, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaign)
	return campaign, nil
}

func UpdateCampaign(id uint, campaign *model.Campaign) (*model.Campaign, error) {
	existingCampaign, err := GetCampaign(id)
	if existingCampaign == nil || err != nil {
		return existingCampaign, err
	}
	existingCampaign.Name = campaign.Name
	existingCampaign.OrganizerName = campaign.OrganizerName
	existingCampaign.TargetAmount = campaign.TargetAmount
	existingCampaign.DonationMinimum = campaign.DonationMinimum
	result := db.DB.Save(existingCampaign)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated campaign.")
	entry.Tracef("Updated: %v", existingCampaign)
	return existingCampaign, nil
}

func DeleteCampaign(id uint) (*model.Campaign, error) {
	campaign, err := GetCampaign(id)
	if campaign == nil || err != nil {
		return campaign, err
	}
	result := db.DB.Delete(campaign)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted campaign.")
	entry.Tracef("Deleted: %v", campaign)
	return campaign, nil
}
