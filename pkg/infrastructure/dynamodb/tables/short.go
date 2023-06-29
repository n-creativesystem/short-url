package tables

const Short = "shorts"

var ShortColumns = struct {
	Key       string
	URL       string
	Author    string
	CreatedAt string
	UpdatedAt string
}{
	Key:       "key",
	URL:       "url",
	Author:    "author",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
