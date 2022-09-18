package main

type TokenResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TokenRequest struct {
	Token string `json:"token"`
	Type  int    `json:"type"`
	Time  int    `json:"time"`
}
