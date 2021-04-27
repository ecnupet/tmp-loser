package model

type Auth struct {
	State  int64  `json:"state"`
	Detail string `json:"detail"`
	Data   Data   `json:"data"`
}

type Data struct {
	Authorization string `json:"authorization"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Message bool `json:"message"`
}
