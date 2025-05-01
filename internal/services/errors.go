package services

import "errors"

var (
	InternalServerError = errors.New("Internal server error")
	NotFound            = errors.New("Category not found")
	AccessDenied        = errors.New("Access denied")
	InvalidName         = errors.New("Invalid category name")
	AlreadyExists       = errors.New("Already exists")
)
