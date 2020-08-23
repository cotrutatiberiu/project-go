package models
import "time"

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
	AdminLevel int
	Created    time.Time
	Updated    time.Time
}