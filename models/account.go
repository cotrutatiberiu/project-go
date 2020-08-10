package models

// Account model
type Account struct {
	AccountID  uint64
	FirstName  string
	LastName   string
	UserName   string
	Email      string
	Language   string
	Password   string
	Active     bool
	AdminLevel int64
	Created    int64
	Updated    int64
}
