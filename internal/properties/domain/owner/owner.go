package owner

func (o *Owner) ID() string {
	return o.id
}
func (o *Owner) Name() string {
	return o.name
}
func (o *Owner) Email() string {
	return o.email
}
func (o *Owner) Telephone() string {
	return o.telephone
}
func (o *Owner) Metadata() Metadata {
	return o.metadata
}
