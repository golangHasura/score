package dto

type Team struct {
	EntityId int      `json:"entity_id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Logo     string   `json:"logo,omitempty"`
	Players  []Player `json:"players,omitempty"`
	Format   Format   `json:"format,omitempty"`
}

type Player struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Format struct {
	EntityId int    `json:"entityId,omitempty"`
	Name     string `json:"name,omitempty"`
}

type RetrieveAllResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
