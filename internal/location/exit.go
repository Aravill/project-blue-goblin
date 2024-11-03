package location

import Item "blue-goblin/internal/item"

type Lock struct {
	Description string `json:"description"`
	KeyId       string `json:"keyId"`
	Status      bool   `json:"status"`
}
type Exit struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Keyword     string `json:"keyword"`
	Destination string `json:"destination"`
	Lock        *Lock  `json:"lock"`
}

func (l *Lock) Unlock(item *Item.Item) bool {
	if item.Id == l.KeyId {
		l.Status = false
		return true
	}
	return false
}
