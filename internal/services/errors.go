package services

import "errors"

var (
	InternalServerError = errors.New("Internal server error")
	CategoryNotFound    = errors.New("Category not found")
	AccessDenied        = errors.New("Access denied")
)
