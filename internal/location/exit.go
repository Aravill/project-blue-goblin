package location

import Item "blue-goblin/internal/item"

type Lock struct {
	Description string `json:"description"`
	KeyId       string `json:"keyId"`
	Status      bool   `json:"status"`
}
type Exit struct {
	Id          string
	Description string
	Keyword     string
	Destination string
	Lock        *Lock `json:"lock"`
}

func (l *Lock) Unlock(item *Item.Item) bool {
	if item.Id == l.KeyId {
		l.Status = false
		return true
	}
	return false
}
