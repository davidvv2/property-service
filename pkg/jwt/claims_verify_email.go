package jwt

// VerifyEmail :This jwt token will be used to verify a users email address.
type VerifyEmail struct {
	ID   string `bson:"_id" json:"id,omitempty" `
	UUID string `json:"uuid" validate:"required,uuid4"`
	Iss  string `json:"iss"  validate:"required,alpha"`
	Sub  string `json:"sub"  validate:"required,alpha"`

	Server string `json:"server" validate:"required,alpha"`

	Exp int64 `json:"exp"  validate:"required,numeric"`
	Nbf int64 `json:"nbf"  validate:"required,numeric"`
}
