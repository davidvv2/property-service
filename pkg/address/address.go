package address

// Address represents a physical address that can be converted to GeoJSON coordinates.
type Address struct {
	FirstLine  string `bson:"FirstLine" json:"firstLine"`   // First line of the address, e.g., "123 Main St"
	Street     string `bson:"Street" json:"street"`         // Street name, e.g., "Main St"
	City       string `bson:"City" json:"city"`             // City name, e.g., "Springfield"
	County     string `bson:"County" json:"county"`         // County name, e.g., "Sangamon County"
	Country    string `bson:"Country" json:"country"`       // Country name, e.g., "USA"
	PostalCode string `bson:"PostalCode" json:"postalCode"` // Postal code, e.g., "62701"
	// GeoJSON is the GeoJSON representation of the address.
	// It will be populated after geocoding the address.
	GeoJSON *GeoJSONCoordinates `bson:"GeoJSON" json:"geoJSON"` // GeoJSON coordinates of the address
}

// GeoJSONCoordinates holds the GeoJSON representation of a point.
type GeoJSONCoordinates struct {
	Type        string     `bson:"Type" json:"type"`               // Always "Point"
	Coordinates [2]float64 `bson:"Coordinates" json:"coordinates"` // [longitude, latitude]
}

func (a Address) IsEmpty() bool {
	return a.FirstLine == "" && a.Street == "" && a.City == "" && a.County == "" && a.Country == "" && a.PostalCode == "" && a.GeoJSON == nil
}
