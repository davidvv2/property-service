package jwt

import "property-service/pkg/permissions/scopes"

// ResetPassword :This is the reset password claim struct.
type ResetPassword struct {
	ID string `bson:"_id" json:"id,omitempty" `

	UUID string `json:"uuid" validate:"required,uuid4"`

	Iss    string `json:"iss"  validate:"required,alpha"`
	Sub    string `json:"sub"  validate:"required,alpha"`
	Server string `json:"server" validate:"required,alpha"`

	Scopes scopes.Scopes `bson:"Scopes" validate:"dive,required"`

	Exp int64 `json:"exp"  validate:"required,numeric"`
	Nbf int64 `json:"nbf"  validate:"required,numeric"`

	IsRefresh     bool `json:"isRefresh" validate:"boolean"`
	MFA           bool `json:"mfa" validate:"boolean"`
	ResetToken    bool `json:"resetToken" validate:"boolean"`
	EmailVerified bool `bson:"EmailVerified" json:"emailVerified" validate:"boolean"`
}

func (rp *ResetPassword) IsRefreshToken() bool {
	return rp.IsRefresh
}

func (rp *ResetPassword) IsPasswordResetToken() bool {
	return rp.ResetToken
}

func (rp *ResetPassword) IsEmailVerified() bool {
	return rp.EmailVerified
}

func (rp *ResetPassword) GetScope() scopes.Scopes {
	return rp.Scopes
}
