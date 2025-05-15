package headers

// AuthHeaders : Header metadata struct.
type AuthHeaders struct {
	Auth      string `mod:"trim" validate:"required"`
	Lang      string `mod:"trim" validate:"required,bcp47_language_tag"`
	Agent     string `mod:"trim" validate:"required,min=10,max=1000"`
	Device    string `mod:"trim" validate:"required,min=10,max=1000"`
	IPAddress string `mod:"trim" validate:"required,hostname_port"`
}
