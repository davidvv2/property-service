package address

import (
	"context"
	"errors"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
)

// Geocoder defines an interface for geocoding operations.
type Geocoder interface {
	// Geocode converts a real address to GeoJSON coordinates.
	Geocode(ctx context.Context, address string) (*GeoJSONCoordinates, error)
	// ReverseGeocode converts geolocation coordinates to a human-readable address.
	ReverseGeocode(ctx context.Context, lat, lng float64) (Address, error)
}

// googleGeocoderClient is our implementation of Geocoder using Google Maps API.
type googleGeocoderClient struct {
	apiKey   string
	geocoder geo.Geocoder
}

// NewGoogleGeocoderClient creates a new Geocoder client using the provided API key.
func NewGoogleGeocoderClient(apiKey string) Geocoder {
	return &googleGeocoderClient{
		apiKey:   apiKey,
		geocoder: google.Geocoder(apiKey),
	}
}

// Geocode converts a real address to GeoJSON coordinates using the Google geocoding service.
func (c *googleGeocoderClient) Geocode(ctx context.Context, address string) (*GeoJSONCoordinates, error) {
	if address == "" {
		return nil, errors.New("address cannot be empty")
	}

	location, err := c.geocoder.Geocode(address)
	if err != nil {
		return nil, err
	}

	return &GeoJSONCoordinates{
		Type:        "Point",
		Coordinates: [2]float64{location.Lng, location.Lat},
	}, nil
}

// ReverseGeocode converts latitude and longitude values to a human-readable address using the Google geocoding service.
func (c *googleGeocoderClient) ReverseGeocode(ctx context.Context, lat, lng float64) (Address, error) {
	location, err := c.geocoder.ReverseGeocode(lat, lng)
	if err != nil {
		return Address{}, err
	}

	if location == nil || location.FormattedAddress == "" {
		return Address{}, errors.New("no address found for the given location")
	}
	new_address := Address{
		FirstLine:  location.HouseNumber,
		Street:     location.Street,
		City:       location.City,
		County:     location.County,
		Country:    location.Country,
		PostalCode: location.Postcode,
		GeoJSON: &GeoJSONCoordinates{
			Type:        "Point",
			Coordinates: [2]float64{lng, lat},
		},
	}
	return new_address, nil
}
