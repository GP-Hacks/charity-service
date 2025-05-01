package models

type Collection struct {
	ID           int64
	CategoryID   int64
	Name         string
	Description  string
	Organization string
	Phone        string
	Website      string
	Goal         int
	Current      int
	Photo        string
}

type CollectionWithCategory struct {
	ID           int64
	Category     string
	Name         string
	Description  string
	Organization string
	Phone        string
	Website      string
	Goal         int
	Current      int
	Photo        string
}

func (c *Collection) ToCollectionWithCategory(category string) *CollectionWithCategory {
	return &CollectionWithCategory{
		ID:           c.ID,
		Category:     category,
		Name:         c.Name,
		Description:  c.Description,
		Organization: c.Organization,
		Phone:        c.Phone,
		Website:      c.Website,
		Goal:         c.Goal,
		Current:      c.Current,
		Photo:        c.Photo,
	}
}
