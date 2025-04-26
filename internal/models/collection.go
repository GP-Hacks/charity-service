package models

type Collection struct {
	ID           int
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
