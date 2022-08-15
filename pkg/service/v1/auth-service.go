package v1

import (
	"errors"

	"github.com/jinzhu/gorm"
	v1 "github.com/sajanjswl/auth/gen/go/auth/v1"

	// v1 "github.com/sajanjswl/auth/pkg/api/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
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
}

// register db wiht server
func NewAuthServiceServer(db *gorm.DB, logger *zap.Logger) v1.AuthServiceServer {
	return &authServiceServer{
		db:     db,
		logger: logger,
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

// func (s *userServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

// 	if err := checkAPI(req.GetApiVersion()); err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	db := s.db

// 	user, err := utils.FindUser(db, config.FindUserByEmail, req.GetEmailID())
// 	if user == nil {
// 		log.Error(req.GetEmailID(), " ", err)
// 		return nil, err
// 	}

// 	if user.IsBlocked {

// 		if user.TimeTillBlock.Before(time.Now()) {
// 			user.IsBlocked = false
// 			user.TimeTillBlock = time.Time{}

// 			goto CONTINUE
// 		}

// 		log.Error(user.Email, " is blocked total invalid attempts ", user.WrongPwdCount)

// 		return nil, status.Error(403, config.ForbiddenReques403)
// 	}

// CONTINUE:

// 	//authenticating password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))

// 	if err != nil {

// 		user.WrongPwdCount = user.WrongPwdCount + 1
// 		if user.WrongPwdCount >= 3 {
// 			user.IsBlocked = true

// 			blockTimePeriod, _ := strconv.Atoi(os.Getenv("USER_BLOCKED_RESET_TIME"))

// 			user.TimeTillBlock = time.Now().Add(time.Duration(blockTimePeriod) * time.Minute)
// 		}

// 		if err := utils.UpdateIsBlockedWrongPwdCountAndTimeTillBlock(db, user.Email, user.WrongPwdCount, user.IsBlocked, user.TimeTillBlock); err != nil {

// 			return nil, err

// 		}

// 		return nil, status.Errorf(codes.Unauthenticated, config.LoginFailed)
// 	}
// 	user.WrongPwdCount = 0
// 	user.AccessToken, user.RefreshToken, err = utils.GetAccessTokenAndRefreshtoken(user.Email)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "Internal error")
// 	}

// 	if err := utils.SaveLoginDetails(db, user.WrongPwdCount, user.IsBlocked, user.TimeTillBlock, user.Email, user.AccessToken, user.RefreshToken); err != nil {
// 		return nil, err
// 	}

// 	response := &v1.LoginResponse{
// 		RefreshToken: user.RefreshToken,
// 		AccessToken:  user.AccessToken,
// 		Status:       "Hello  " + user.FirstName + config.SuccessFullLoginMessage,
// 	}

// 	return response, nil

// }

// func (s *userServiceServer) Register(ctx context.Context, req *v1.RegistrationRequest) (*v1.RegistrationResponse, error) {

// 	if err := checkAPI(req.GetApiVersion()); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	db := s.db
// 	user, err := utils.FindUser(db, config.FindUserByEmail, req.GetUser().GetEmailID())

// 	if user != nil {
// 		log.Error(req.GetUser().GetEmailID(), " ", config.EmailAlreadyExists)
// 		return nil, status.Error(codes.AlreadyExists, config.EmailAlreadyExists)
// 	}

// 	// //bycrpting the plaint text password
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetUser().GetPassword()), bcrypt.MinCost)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Errorf(codes.Internal, config.InternalError)
// 	}

// 	//passing  request data into dbUser type
// 	user = &schema.DbUser{
// 		Uuid:           guuid.New(),
// 		PrimaryContact: req.GetUser().GetMobileNumber(),
// 		Email:          req.GetUser().GetEmailID(),
// 		Password:       string(passwordHash),
// 		FirstName:      req.GetUser().GetFirstName(),
// 		LastName:       req.GetUser().GetLastName(),
// 		CreatedAt:      time.Now().Format(time.RFC3339),
// 		PendingDetails: true,
// 		IsVerified:     true,
// 		IsBlocked:      false,
// 	}

// 	//registering user
// 	registeringUser := db.Table(schema.TableUsers).Create(&user)
// 	if registeringUser.Error != nil {
// 		log.Error(registeringUser.Error)
// 		err := status.Error(codes.Internal, registeringUser.Error.Error())
// 		return nil, err
// 	}

// 	return &v1.RegistrationResponse{Message: "successfully registerd  " + req.GetUser().GetFirstName() + "  " + req.GetUser().GetLastName()}, nil
// }

// func (s *userServiceServer) Home(ctx context.Context, req *v1.HomeRequest) (*v1.HomeResponse, error) {
// 	// if err := checkAPI(req.GetApiVersion()); err != nil {
// 	// 	log.Println(err)
// 	// 	return nil, err
// 	// }
// 	md, ok := metadata.FromIncomingContext(ctx)

// 	log.Println(ok)

// 	log.Println("Printing headers", md)

// 	return &v1.HomeResponse{
// 		Message: "you are in home your hardwork paid off",
// 	}, nil

// }

// func (s *userServiceServer) OTP(ctx context.Context, req *v1.LoginWithOTPRequest) (*v1.LoginWithOTPResponse, error) {

// 	if err := checkAPI(req.GetApiVersion()); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	db := s.db

// 	user, err := utils.FindUser(db, config.FindUserByEmail, req.GetEmailID())

// 	if user == nil {
// 		return nil, err
// 	}

// 	otplength, err := strconv.Atoi(os.Getenv("OTP_LENGTH"))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	token := mail.GenerateOTP(otplength)

// 	otpExpirePeriod, err := strconv.Atoi(os.Getenv("OTP_EXPIRE_TIME"))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	if err := utils.SaveOTPAndOTPExpiryTime(db, user.Email, token, time.Now().Add(time.Duration(otpExpirePeriod)*time.Minute)); err != nil {

// 		return nil, err

// 	}

// 	if err := mail.SendOTP(token, user.Email, user.PrimaryContact); err != nil {
// 		return nil, err
// 	}

// 	return &v1.LoginWithOTPResponse{
// 		Message: "OTP sent Successfully",
// 	}, nil
// }

// func (s *userServiceServer) VerifyOTP(ctx context.Context, req *v1.VerifyOTPRequest) (*v1.VerifyOTPResponse, error) {

// 	if err := checkAPI(req.GetApiVersion()); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	db := s.db

// 	user, err := utils.FindUser(db, config.FindUserByEmail, req.GetEmailID())
// 	if user == nil {
// 		log.Error(req.GetEmailID(), " ", err)
// 		return nil, err
// 	}

// 	if user.OtpExpiry.Before(time.Now()) {
// 		log.Println("OTP Has expired for ", req.GetEmailID())
// 		return nil, errors.New("OTP Expired")
// 	}
// 	flag := strings.Compare(user.CurrentOtp, req.GetOTP())

// 	if flag != 0 {

// 		log.Println("current OTP requested OTP ", user.CurrentOtp, req.GetOTP())

// 		return nil, errors.New("wrong credential")

// 	}

// 	accessToken, refreshToken, err := utils.GetAccessTokenAndRefreshtoken(user.Email)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "Internal error")
// 	}

// 	return &v1.VerifyOTPResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}, nil

// }

// func (s *userServiceServer) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (*v1.ResetPasswordResponse, error) {
// 	if err := checkAPI(req.GetApiVersion()); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	db := s.db

// 	user, err := utils.FindUser(db, config.FindUserByEmail, req.GetEmailID())
// 	if user == nil {
// 		log.Error(req.GetEmailID(), " ", err)
// 		return nil, err
// 	}

// 	if user.OtpExpiry.Before(time.Now()) {
// 		log.Println("OTP Has expired for ", req.GetEmailID())
// 		return nil, errors.New("OTP Expired")
// 	}
// 	flag := strings.Compare(user.CurrentOtp, req.GetOTP())

// 	if flag != 0 {

// 		log.Println("current OTP %s requested OTP %s", user.CurrentOtp, req.GetOTP())

// 		return nil, errors.New("wrong credential")

// 	}

// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.MinCost)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Errorf(codes.Internal, config.InternalError)
// 	}

// 	user.AccessToken, user.RefreshToken, err = utils.GetAccessTokenAndRefreshtoken(user.Email)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "Internal error")
// 	}

// 	utils.SavePasswordAccessTokenAndRefreshToken(db, user.Email, string(passwordHash), user.AccessToken, user.RefreshToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &v1.ResetPasswordResponse{
// 		AccessToken:  user.AccessToken,
// 		RefreshToken: user.RefreshToken,
// 	}, nil
// }

// func (s *userServiceServer) RequestTokens(ctx context.Context, req *v1.AccessTokenAndRefreshTokenRequest) (*v1.AccessTokenAndRefreshTokenResponse, error) {

// 	if err := checkAPI(req.GetApiVersion()); err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	claims, err := auth.VerifyRefreshToken(req.GetRefreshToken())
// 	if err != nil {
// 		log.Println(err)
// 		return nil, errors.New("failed to verify token")
// 	}

// 	accessToken, refreshToken, err := utils.GetAccessTokenAndRefreshtoken(claims.UserID)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "Internal error")
// 	}

// 	err = utils.SaveAccessTokenAndRefreshToken(s.db, claims.UserID, accessToken, refreshToken)
// 	if err != nil {

// 		log.Error(err)
// 		return nil, status.Error(codes.Internal, "Internal error")
// 	}

// 	return &v1.AccessTokenAndRefreshTokenResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}, nil
// }
