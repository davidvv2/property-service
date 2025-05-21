package jwt

import "property-service/pkg/permissions/scopes"

// AuthClaims :This is the Authentication claim struct.
type AuthClaims struct {
	ID string `bson:"_id" json:"id,omitempty" `

	UUID string `json:"uuid" validate:"required,uuid4"`

	Iss    string `json:"iss"  validate:"required,alpha"`
	Sub    string `json:"sub"  validate:"required,alpha"`
	Server string `json:"server" validate:"required,alpha"`

	Scopes scopes.Scopes `bson:"Scopes"`

	Exp int64 `json:"exp"  validate:"required,numeric"`
	Nbf int64 `json:"nbf"  validate:"required,numeric"`

	IsRefresh     bool `json:"isRefresh"`
	MFA           bool `json:"mfa"`
	ResetToken    bool `json:"resetToken"`
	EmailVerified bool `bson:"EmailVerified" json:"emailVerified"`
}

func (rp *AuthClaims) IsRefreshToken() bool {
	return rp.IsRefresh
}

func (rp *AuthClaims) IsPasswordResetToken() bool {
	return rp.ResetToken
}

func (rp *AuthClaims) IsEmailVerified() bool {
	return rp.EmailVerified
}

func (rp *AuthClaims) GetScope() scopes.Scopes {
	return rp.Scopes
}
