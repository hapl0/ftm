package main

type investigation struct {
	id      int      `json:"id"`
	name    string   `json:"name"`
	users   []int    `json:"users_ID"`
	admins  []int    `json:"admins_ID"`
	wallets []wallet `json:"wallet"`
}

type user struct {
	id     int    `json:"id"`
	name   string `json:"name"`
	groups []int  `json:"groups_ID"`
}

type group struct {
	id   int    `json:"id"`
	name string `json:"name"`
}
