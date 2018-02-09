package main

var auth string

type Sicd struct {
	Name string
	Birth_date string
	Personal_number string
	Branch_code string
}

type Dh struct {
	Name string
	Birth_date string
}

type Token struct {
	Access_token string
	Expires_in uint64
	Token_type string
	Scope string
}

type Nik struct {
	NIK string
}

type Controller struct {
	NIK string
	Name string
	Birth_date string
	Personal_number string
	Branch_code string
}