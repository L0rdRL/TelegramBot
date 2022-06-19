package telegram

type UpdatesRespons struct {
	Ok     bool     `json: "ok"`
	Result []Update `jason: "result"`
}

type Update struct {
	ID      int    `json: "update_id"`
	Message string `jason: "message"`
}
