package configs

// Config : A config struct.
type Config struct {
	Database      DatabaseStruct
	Backend       BackendStruct
	SchemeVersion SchemeVersionStruct
	Gcloud        GoogleCloudStruct
	Caching       CachingStruct
}

type SchemeVersionStruct struct {
	AuthenticationUser string
	FactsFact          string
}
type ServicesStruct struct {
}

type BackendStruct struct {
	Version     string
	Debug       string
	URL         string
	Environment string
	Port        string
	Ssl         string
}

type DatabaseStruct struct {
	URI      string
	Username string
	Password string
	CSFLE    CSFLE
}

type CSFLE struct {
	Email       string
	PrivateKey  string
	KMSProvider string

	ProjectID string
	Location  string
	KeyRing   string
	KeyName   string
}

type GoogleCloudStruct struct {
	AccessID           string
	PubSubSubscription string
	PrivateKey         []byte
}

type CachingStruct struct {
	Addr     string
	Password string
	DB       int
}
