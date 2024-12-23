package sdk

// PointOfInitiation struct contains the logic for dynamic and static QR code settings
type PointOfInitiation struct {
	DynamicQR string
	StaticQR  string
}

// NewPointOfInitiation initializes and returns a new PointOfInitiation instance
func NewPointOfInitiation(emv *EMV) *PointOfInitiation {
	return &PointOfInitiation{
		DynamicQR: emv.DefaultDynamicQR,
		StaticQR:  emv.DefaultStaticQR,
	}
}

// Dynamic retrieves the dynamic QR code setting
func (p *PointOfInitiation) Dynamic() string {
	return p.DynamicQR
}

// Static retrieves the static QR code setting
func (p *PointOfInitiation) Static() string {
	return p.StaticQR
}
