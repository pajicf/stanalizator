package keeper

type RealEstate int

const (
	RentableApartment RealEstate = iota
	PurchasableApartment
	PurchasableHouse
	PurchasableLand
	PurchasablePremises
	PurchasableOfficeSpace
)

func (re RealEstate) ApiValue() string {
	// Map of our custom enum to Halo Oglasi API values
	return [...]string{"13", "12", "24", "26", "32", "38"}[re]
}
