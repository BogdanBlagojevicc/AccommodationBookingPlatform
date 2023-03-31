package model

type UserRole int

const (
	Unauthenticated UserRole = iota
	RegularUser
	Administrator
)
