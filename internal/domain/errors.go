package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("Account not found")
	ErrDuplicatedApiKey   = errors.New("Api key already exists")
	ErrInvoiceNotFound    = errors.New("Invoice not found")
	ErrUnauthorizedAccess = errors.New("Unauthorized access")
	ErrInvalidAmount      = errors.New("Invalid amount")
	ErrInvalidStatus      = errors.New("Invalid status")
)
