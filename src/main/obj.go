package main

type plurkerObj struct {
	UserInfo plurkerInfoObj `json:"user_info"`
}

type plurkerInfoObj struct {
	UserID   int     `json:"uid"`
	DispName string  `json:"display_name"`
	NickName string  `json:"nick_name"`
	FullName string  `json:"full_name"`
	Karma    float32 `json:"karma"`
	ID       int     `json:"id"`
}
