package v1

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/sajanjswl/auth/config"
	v1 "github.com/sajanjswl/auth/gen/go/auth/v1"
	"github.com/sajanjswl/auth/models"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type authServiceServer struct {
	db     *gorm.DB
	logger *zap.Logger
	v1.UnimplementedAuthServiceServer
	config *config.Config
}

// register db wiht server
func NewAuthServiceServer(db *gorm.DB, logger *zap.Logger, cfg *config.Config) v1.AuthServiceServer {
	return &authServiceServer{
		db:     db,
		logger: logger,
		config: cfg,
	}
}

// checkAPI checks if the API version requested by client is supported by server
func checkAPI(api string) error {

	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}

		return nil

	}
	return errors.New("unsupported API version: Api version cannot be nil")
}

func (s *authServiceServer) LoginUser(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	if err := checkAPI(req.GetApiVersion()); err != nil {
		s.logger.Warn("api version conflice", zap.Error(err))
		return nil, err
	}

	user := &models.User{}
	if err := models.GetUser(s.db, user, req.GetEmailID()); err != nil {
		s.logger.Warn("error reading user from datatbase", zap.String("email", user.Email), zap.Error(err))
		return nil, status.Error(codes.NotFound, "user not found")
	}

	//authenticating password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		s.logger.Info("password incorrect", zap.String("email", user.Email), zap.Error(err))
		return nil, errors.New("Password incorrect")
	}

	response := &v1.LoginResponse{
		Status:  "200",
		Message: "Hello  " + user.Name + "Logged in Successfully!!",
	}
	return response, nil

}

func (s *authServiceServer) RegisterUser(ctx context.Context, req *v1.RegistrationRequest) (*v1.RegistrationResponse, error) {
	if err := checkAPI(req.GetApiVersion()); err != nil {
		s.logger.Warn("Api verison error", zap.Error(err))
		return nil, err
	}
	user := &models.User{}
	if err := models.GetUser(s.db, user, req.GetUser().EmailID); err == nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	//bycrpting the plaint text password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetUser().GetPassword()), bcrypt.MinCost)
	if err != nil {
		s.logger.Warn("Failed to hash password", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	user.Email = req.GetUser().EmailID
	user.Password = string(passwordHash)
	user.Name = req.GetUser().Name
	user.Mobile = req.GetUser().MobileNumber

	if err := models.CreateUser(s.db, user); err != nil {
		s.logger.Warn("Failed to register user", zap.String("email", user.Email), zap.Error(err))
	}

	s.logger.Info("register user", zap.String("email", user.Email))

	return &v1.RegistrationResponse{Message: "successfully registerd  " + req.GetUser().GetName() + "  " + req.GetUser().EmailID}, nil
}

func (s *authServiceServer) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	if err := checkAPI(req.GetApiVersion()); err != nil {
		s.logger.Warn("Api verison error", zap.Error(err))
		return nil, err
	}
	user := &models.User{}
	if err := models.GetUser(s.db, user, req.User.EmailID); err != nil {
		s.logger.Warn("error reading user from datatbase", zap.String("email", user.Email), zap.Error(err))
		return nil, status.Error(codes.NotFound, "user not found")
	}

	user.Email = req.User.EmailID
	user.Name = req.User.Name
	user.Mobile = req.User.MobileNumber

	if err := models.UpdateUser(s.db, user); err != nil {
		s.logger.Warn("Failed to update user", zap.String("email", user.Email), zap.Error(err))
	}

	s.logger.Info("update user", zap.String("email", user.Email))

	return &v1.UpdateUserResponse{Message: "successfully update  " + req.GetUser().GetName() + "  " + req.GetUser().EmailID}, nil
}

func (s *authServiceServer) Home(ctx context.Context, req *v1.HomeRequest) (*v1.HomeResponse, error) {
	// if err := checkAPI(req.GetApiVersion()); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	md, ok := metadata.FromIncomingContext(ctx)

	log.Println(ok)

	log.Println("Printing headers", md)

	return &v1.HomeResponse{
		Message: "You are in home!!",
	}, nil

}
