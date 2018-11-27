package idref

// AuthorityRecord is the struct common to all
// authority types, e.g. Person or Organization, etc.
type AuthorityRecord struct {
	ID           string
	DateCreated  string
	DateUpdated  string
	Identifiers  []Identifier
	Person       Person
	Organization Organization
}

// Document is a single bibliographic reference
type Document struct {
	Citation string `json:"citation"`
	Source   string `json:"referentiel"`
	ID       int    `json:"id"`
	URI      string `json:"URI"`
	URL      string `json:"URL"`
}

// Identifier is an ID for the authority
// in a source other than IdRef itself
// e.g. the French National library or ISNI
type Identifier struct {
	ID     string
	Source string
}

// Organization regroups fields specific to an Organization authority record
type Organization struct {
	AltLabel    []string
	DateOfBirth string
	Name        string
	PrefLabel   string
}

// Person regroups fields specific to a Person authority record
type Person struct {
	DateBirth  string
	DateDeath  string
	FamilyName string
	GivenName  string
	Name       string
	PrefLabel  string
	Surname    string
}

// ReferencesResult is used to unamrshall
// the raw json result coming from the references web service
type ReferencesResult struct {
	Role []struct {
		Marc21Code  string     `json:"marc21Code"`
		RoleName    string     `json:"roleName"`
		Count       int        `json:"count"`
		Doc         []Document `json:"doc"`
		UnimarcCode string     `json:"unimarcCode"`
	} `json:"role"`
	Name       string `json:"name"`
	CountRoles int    `json:"countRoles"`
}
