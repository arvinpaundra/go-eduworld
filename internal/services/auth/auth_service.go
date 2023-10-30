package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/adapters/response"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, input *request.Register) error
	Login(ctx context.Context, input *request.Login) (*response.Login, error)
	ChangePassword(ctx context.Context) error
	CheckSession(ctx context.Context, username string) (*response.CheckSession, error)
	Logout(ctx context.Context, userId string) error
}

type authService struct {
	userRepository     contracts.UserRepository
	deviceRepository   contracts.DeviceRepository
	sessionRepository  contracts.SessionRepository
	interestRepository contracts.InterestRepository
	cacheRepository    contracts.CacheRepository
}

func NewAuthService(
	userRepository contracts.UserRepository,
	deviceRepository contracts.DeviceRepository,
	sessionRepository contracts.SessionRepository,
	interestRepository contracts.InterestRepository,
	cacheRepository contracts.CacheRepository,
) AuthService {
	return &authService{
		userRepository:     userRepository,
		deviceRepository:   deviceRepository,
		sessionRepository:  sessionRepository,
		interestRepository: interestRepository,
		cacheRepository:    cacheRepository,
	}
}

func (a *authService) Register(ctx context.Context, input *request.Register) error {
	// begin transaction
	tx, err := a.userRepository.Begin(ctx)
	if err != nil {
		return err
	}

	// check username availability
	user, err := a.userRepository.FindOne(ctx, "username", "username=?", input.Username)
	if err != nil && err.Error() != constants.ErrBunNotNotFound.Error() {
		return err
	}

	// return error if username already taken
	if user != nil {
		return constants.ErrUsernameAlreadyTaken
	}

	// check interest availability
	if _, err := a.interestRepository.FindOne(ctx, "id", "id=?", input.InterestId); err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			return constants.ErrInterestNotFound
		}

		return err
	}

	// setup new user
	newUser := entities.User{
		ID:             utils.GetID(),
		InterestId:     input.InterestId,
		Username:       input.Username,
		Email:          nil,
		Password:       utils.HashPassword(input.Password),
		Fullname:       input.Fullname,
		Status:         "active",
		Role:           input.Role,
		Bio:            nil,
		Phone:          nil,
		BirthDate:      nil,
		ProfilePicture: nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// do insert new user
	if err := a.userRepository.Save(ctx, tx, &newUser); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			return errorRollback
		}
		return err
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (a *authService) Login(ctx context.Context, input *request.Login) (*response.Login, error) {
	// begin transaction
	tx, err := a.userRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	// check user by username
	user, err := a.userRepository.FindOne(ctx, "id,username,password,status", "username=?", input.Username)
	if err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			return nil, constants.ErrUserNotFound
		}

		return nil, err
	}

	// compare password
	if ok := utils.ComparePassword(user.Password, input.Password); !ok {
		return nil, constants.ErrPasswordIncorrect
	}

	// check is there is device existing
	device, err := a.deviceRepository.FindOne(ctx, "id", "user_id=?", user.ID)
	if err != nil && err.Error() != constants.ErrBunNotNotFound.Error() {
		return nil, err
	}

	// if there is a device connected, remove device and the session
	if device != nil {
		// remove device
		if err := a.deviceRepository.Remove(ctx, tx, "id=?", device.ID); err != nil {
			if errorRollback := tx.Rollback(); errorRollback != nil {
				return nil, errorRollback
			}
			return nil, err
		}

		// remove session
		if err := a.sessionRepository.Remove(ctx, tx, "user_id=?", user.ID); err != nil {
			if errorRollback := tx.Rollback(); errorRollback != nil {
				return nil, errorRollback
			}
			return nil, err
		}

		// remove session from cache
		key := fmt.Sprintf("%s::%s", constants.KeySession, user.Username)

		if err := a.cacheRepository.Del(ctx, key); err != nil {
			return nil, err
		}
	}

	// setup new device
	newDevice := entities.Device{
		ID:        utils.GetID(),
		UserId:    user.ID,
		Name:      input.Device,
		Platform:  input.Platform,
		IPAddress: input.IPAddress,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// do insert new device
	if err := a.deviceRepository.Save(ctx, tx, &newDevice); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			return nil, errorRollback
		}
		return nil, err
	}

	// set time issued and expired token
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(time.Hour * 24) // 24 hours

	// generate new token
	token, err := utils.GenerateToken(&utils.JWTCustomOption{
		ID:        user.ID,
		Role:      user.Role,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	})

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	// setup new session
	newSession := entities.Session{
		ID:               utils.GetID(),
		UserId:           user.ID,
		Token:            token,
		RefreshToken:     nil,
		FCMToken:         input.FCMToken,
		GoogleOAuthToken: input.GoogleOAuthToken,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// do insert new session
	if err := a.sessionRepository.Save(ctx, tx, &newSession); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			return nil, errorRollback
		}

		return nil, err
	}

	// save creds into redis
	key := fmt.Sprintf("%s::%s", constants.KeySession, user.Username)
	ttl := time.Minute * 15

	if err := a.cacheRepository.Set(ctx, key, newSession, ttl); err != nil {
		return nil, err
	}

	// commit transaction
	if errorCommit := tx.Commit(); errorCommit != nil {
		return nil, errorCommit
	}

	return &response.Login{Token: token, IssuedAt: issuedAt.Format("2006-02-02 15:04:05"), ExpiredAt: expiredAt.Format("2006-02-02 15:04:05")}, nil
}

func (a *authService) ChangePassword(ctx context.Context) error {
	panic("not implemented")
}

func (a *authService) CheckSession(ctx context.Context, username string) (*response.CheckSession, error) {
	panic("not implemented")
}

func (a *authService) Logout(ctx context.Context, userId string) error {
	tx, err := a.userRepository.Begin(ctx)

	if err != nil {
		return err
	}

	user, err := a.userRepository.FindOne(ctx, "id,username", "id=?", userId)

	if err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			return constants.ErrUserNotFound
		}

		return err
	}

	if err := a.deviceRepository.Remove(ctx, tx, "user_id=?", user.ID); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			return err
		}

		return err
	}

	if err := a.sessionRepository.Remove(ctx, tx, "user_id=?", user.ID); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			return err
		}

		return err
	}

	// remove from redis
	key := fmt.Sprintf("%s::%s", constants.KeySession, user.Username)

	if err := a.cacheRepository.Del(ctx, key); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			return err
		}

		return err
	}

	// commit transaction
	if errorCommit := tx.Commit(); errorCommit != nil {
		return errorCommit
	}

	return nil
}
