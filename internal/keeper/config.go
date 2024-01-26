package keeper

type Config struct {
	GeoConfig          GeoConfig
	ToMail             string
	EmailJSUserID      string
	EmailJSServiceID   string
	EmailJSTemplateID  string
	EmailJSAccessToken string
	HaloApiURL         string
}

type GeoConfig []struct {
	MinPrice         int        `json:"minPrice"`
	MaxPrice         int        `json:"maxPrice"`
	Type             RealEstate `json:"type"`
	GeoPolygonString string     `json:"geoPolygonString"`
}
