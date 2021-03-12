package data

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
)

const (
	//Postgres tables
	TableUsers = "public.userbasic"

	//Postgres columns
	ColID               = "id"
	ColUUID             = "uuid"
	ColFirstName        = "first_name"
	ColMiddleName       = "middle_name"
	ColLastName         = "last_name"
	ColDateBirth        = "date_birth"
	ColEmail            = "email"
	ColPrimayContact    = "primary_contact"
	ColSecondaryContact = "secondary_contact"
	ColPincode          = "pincode"
	ColAddress          = "adress"
	ColState            = "state"
	ColPassword         = "password"
	ColRegisterdAs      = "registerd_as"
	ColCreatedAt        = "created_at"
	ColIsBlocked        = "is_blocked"
	ColIsVerified       = "is_verified"
	ColPendingDetails   = "pending_details"
	ColAlias            = "alias"
	ColGoogleID         = "google_id"
	ColFacebookID       = "facebook_id"
	ColCurrentOtp       = "current_otp"
	ColOtpExpiry        = "otp_expiry"
	ColWrongPwdCount    = "wrong_pwd_count"
	ColTimeTillBlock    = "time_till_block"
	ColUserAgent        = "user_agent"
	ColRefreshToken     = "refresh_token"
	ColAccessToken      = "access_token"

	//token keys
	KeyAuthorization = "authorization"
	PrefixBearer     = "bearer_"
)

// Table schema of user basic
type DbUser struct {
	Uuid             uuid.UUID
	FirstName        string
	MiddleName       string
	LastName         string
	DateBirth        date.Date
	Email            string
	PrimaryContact   string
	SecondaryContact string
	Pincode          string
	Adress           string
	State            string
	Password         string
	RegisterdAs      string
	CreatedAt        string
	IsBlocked        bool
	IsVerified       bool
	PendingDetails   bool
	Alias            string
	GoogleId         string
	FacebookId       string
	CurrentOtp       string
	OtpExpiry        time.Time
	WrongPwdCount    int
	TimeTillBlock    time.Time
	UserAgent        string
	RefreshToken     string
	AccessToken      string
}

type AccessTokenDetails struct {
	UserId string

	Issuer string
}

type RefreshTokenDetails struct {
	UserId string
	Issuer string
}

type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	AwsRegion       string
	SMTPUser        string
	SMTPPass        string
	Host            string
	Port            int
}
