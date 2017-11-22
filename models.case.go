package main

type investigation struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Users    []int    `json:"users_ID"`
	Admins   []int    `json:"admins_ID"`
	Wallets  []wallet `json:"wallet"`
	Currency string   `json:"currency"`
}

type user struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Groups []int  `json:"groups_ID"`
}

type group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
