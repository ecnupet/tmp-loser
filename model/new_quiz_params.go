package model

type NewQuizParams struct {
	UserName string		`json:"userName"`
	Types 	 []uint32	`json:"types"`
}