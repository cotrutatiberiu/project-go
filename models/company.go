package models

import "errors"

// Company model
type Company struct {
	CompanyID   uint64
	AccountID   uint64
	Domain      string
	Subdomain   string
	Country     string
	City        string
	Language    string
	Name        string
	Description string
	Verified    bool
	OneMark     int64
	TwoMark     int64
	ThreeMark   int64
	FourMark    int64
	FiveMark    int64
	Created     int64
	Updated     int64
}

// Validate company data
func (c Company) Validate() error {
	if c.AccountID == 0 {
		return errors.New("invalid account ID")
	}
	return nil
}
