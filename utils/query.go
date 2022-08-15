package utils

import (
	"time"

	"github.com/sajanjswl/auth/config"
	schema "github.com/sajanjswl/auth/data"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// updates IsBlocke WrongPwdCount And TimeTillBlock
func UpdateIsBlockedWrongPwdCountAndTimeTillBlock(db *gorm.DB, email string, wrongPwdCount int, isBlocked bool, timeTillBlocked time.Time) error {

	save := db.Table(schema.TableUsers).Where(config.FindUserByEmail, email).Updates(map[string]interface{}{
		schema.ColIsBlocked:     isBlocked,
		schema.ColTimeTillBlock: timeTillBlocked,
		schema.ColWrongPwdCount: wrongPwdCount,
	})
	if save.Error != nil {
		log.Error(save.Error)
		return status.Errorf(codes.Internal, "internal error")
	}

	return nil
}

// saves accessToken and refreshToken
func SaveLoginDetails(db *gorm.DB, wrongPwdCount int, isBlocked bool, timeTilBlock time.Time, email, accessToken, refreshToken string) error {

	save := db.Table(schema.TableUsers).Where(config.FindUserByEmail, email).Updates(map[string]interface{}{
		schema.ColAccessToken:   accessToken,
		schema.ColRefreshToken:  refreshToken,
		schema.ColWrongPwdCount: wrongPwdCount,
		schema.ColIsBlocked:     isBlocked,
		schema.ColTimeTillBlock: timeTilBlock,
	})
	if save.Error != nil {
		log.Error(save.Error)
		return status.Errorf(codes.Internal, "internal error")
	}

	return nil
}

// finds user
func FindUser(db *gorm.DB, conditions interface{}, args ...interface{}) (*schema.DbUser, error) {
	var user schema.DbUser
	findUser := db.Table(schema.TableUsers).
		Where(conditions, args...).
		First(&user)
	if findUser.Error != nil {
		log.Error(findUser.Error)
		switch findUser.Error {
		case gorm.ErrRecordNotFound:
			return nil, status.Errorf(codes.NotFound, config.UserNotFound)

		}

	}

	return &user, nil
}

func SaveOTPAndOTPExpiryTime(db *gorm.DB, email string, token string, otpExpiryTime time.Time) error {

	save := db.Table(schema.TableUsers).Where(config.FindUserByEmail, email).Updates(map[string]interface{}{
		schema.ColCurrentOtp: token,
		schema.ColOtpExpiry:  otpExpiryTime,
	})
	if save.Error != nil {
		log.Error(save.Error)
		return status.Errorf(codes.Internal, "internal error")
	}

	return nil

}

func SaveAccessTokenAndRefreshToken(db *gorm.DB, email string, accessToken string, refreshToken string) error {

	save := db.Table(schema.TableUsers).Where(config.FindUserByEmail, email).Updates(map[string]interface{}{
		schema.ColAccessToken:  accessToken,
		schema.ColRefreshToken: refreshToken,
	})
	if save.Error != nil {
		log.Error(save.Error)
		return status.Errorf(codes.Internal, "internal error")
	}

	return nil

}

func SavePasswordAccessTokenAndRefreshToken(db *gorm.DB, email, password, accessToken, refreshToken string) error {

	save := db.Table(schema.TableUsers).Where(config.FindUserByEmail, email).Updates(map[string]interface{}{
		schema.ColAccessToken:   accessToken,
		schema.ColRefreshToken:  refreshToken,
		schema.ColPassword:      password,
		schema.ColWrongPwdCount: 0,
		schema.ColIsBlocked:     false,
		schema.ColTimeTillBlock: time.Time{},
	})
	if save.Error != nil {
		log.Error(save.Error)
		return status.Errorf(codes.Internal, "internal error")
	}

	return nil

}
