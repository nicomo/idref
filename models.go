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

// AuthorityRole is the role an authority plays in a document
// for instance Author, Translator, etc
type AuthorityRole struct {
	UnimarcCode string
	Marc21Code  string
	RoleName    string
}

// Document is a single bibliographic reference
type Document struct {
	AuthorityRole AuthorityRole
	Citation      string
	ID            string
	Source        string
	URI           string
	URL           string
}

// Documents is a slice of Document
type Documents []Document

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
