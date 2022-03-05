package entities

type TelegramResponseResult struct {
	FileId   string `json:"file_id"`
	FilePath string `json:"file_path"`
}
type TelegramResponse struct {
	Ok     bool                   `json:"ok"`
	Result TelegramResponseResult `json:"result"`
}
