package configs

import (
	"os"
)

// New : Creates and configures a new config object from environmental variables.
func New() Config {
	return Config{
		Backend:       createBackendConfig(),
		Database:      createDatabase(),
		Gcloud:        createGoogleCloud(),
		Caching:       createCaching(),
		SchemeVersion: createSchemeVersion(),
	}
}
func createBackendConfig() BackendStruct {
	return BackendStruct{
		Version:     os.Getenv("version"),
		URL:         os.Getenv("backend"),
		Environment: os.Getenv("environment"),
		Port:        os.Getenv("Port"),
		Ssl:         os.Getenv("ssl"),
		Debug:       os.Getenv("Debug"),
	}
}

func createDatabase() DatabaseStruct {
	return DatabaseStruct{
		URI:      os.Getenv("mongoDBURI"),
		Username: os.Getenv("mongoUser"),
		Password: os.Getenv("mongoPass"),
		CSFLE: CSFLE{
			Email:      os.Getenv("service_account_email"),
			PrivateKey: os.Getenv("service_account_private_key"),
			ProjectID:  os.Getenv("gcloud_project"),
			Location:   os.Getenv("gcloud_location"),

			KMSProvider: os.Getenv("kms_provider"),
			KeyRing:     os.Getenv("key_ring"),
			KeyName:     os.Getenv("key_name"),
		},
	}
}

func createGoogleCloud() GoogleCloudStruct {
	return GoogleCloudStruct{
		AccessID:           os.Getenv("GoogleCloudAccessEmail"),
		PrivateKey:         []byte(os.Getenv("GoogleCloudPrivateKey")),
		PubSubSubscription: os.Getenv("PubSubSubscription"),
	}
}

func createCaching() CachingStruct {
	return CachingStruct{
		Addr:     os.Getenv("redisAddress"),
		Password: os.Getenv("redisPassword"),
		DB:       0,
	}
}

func createSchemeVersion() SchemeVersionStruct {
	return SchemeVersionStruct{
		AuthenticationUser: os.Getenv("SchemeVAuthUser"),
		FactsFact:          os.Getenv("SchemeVFactsFact"),
	}
}
