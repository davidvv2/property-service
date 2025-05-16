package property

import "time"

func (p *Property) ID() string {
	return p.id
}
func (p *Property) OwnerID() string {
	return p.ownerID
}
func (p *Property) Category() string {
	return p.category
}
func (p *Property) Description() string {
	return p.description
}
func (p *Property) Title() string {
	return p.title
}
func (p *Property) Metadata() Metadata {
	return p.metadata
}
func (p *Property) Available() bool {
	return p.available
}
func (p *Property) AvailableDate() time.Time {
	return p.availableDate
}
func (p *Property) Address() string {
	return p.address
}
func (p *Property) SaleType() uint8 {
	return p.saleType
}
