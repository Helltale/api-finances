package models

type Account struct { // (db) entity gorup + account
	idAccaunt int64
	tgId      int64
	name      string
	groupId   int64
}

type AccountJSON struct {
	IdAccaunt int64  `json:"id_accaunt"`
	TgId      int64  `json:"tg_id"`
	Name      string `json:"name"`
	GroupId   int64  `json:"group_id"`
}

func (a *Account) ToJSON() (*AccountJSON, error) {
	return &AccountJSON{
		IdAccaunt: a.idAccaunt,
		TgId:      a.tgId,
		Name:      a.name,
		GroupId:   a.groupId,
	}, nil
}

func (a *Account) GetIdAccaunt() int64 {
	return a.idAccaunt
}

func (a *Account) GetTgId() int64 {
	return a.tgId
}

func (a *Account) GetName() string {
	return a.name
}

func (a *Account) GetGroupId() int64 {
	return a.groupId
}

func (a *Account) SetIdAccaunt(id int64) {
	a.idAccaunt = id
}

func (a *Account) SetTgId(id int64) {
	a.tgId = id
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a *Account) SetGroupId(id int64) {
	a.groupId = id
}
