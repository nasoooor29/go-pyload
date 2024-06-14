package models


type ConfigItemListModel struct {
	Desc  string   `json:"desc"`
	List  []string `json:"list"`
	Type  string   `json:"type"`
	Value string   `json:"value"`
}