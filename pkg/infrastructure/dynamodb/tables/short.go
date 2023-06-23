package tables

const Short = "shorts"

var ShortColumns = struct {
	ID        string
	Key       string
	URL       string
	Author    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Key:       "key",
	URL:       "url",
	Author:    "author",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
