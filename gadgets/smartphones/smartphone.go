package smartphones

// Smartphone model structure for smartphones
type Smartphone struct {
	Id            int64 // similar to a long in Java
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

// CreateSmartphoneCMD
 CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	OS            string `json:"os"`
}
