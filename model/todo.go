package todo

type ListItems struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ItemData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Item struct {
	TodoItemData ItemData `json:"todo_item_data"`
	Id           int      `json:"id"`
	Done         bool     `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
