package communication

type messageInput struct {
	title       string
	Description string `json:"description"`
}

type messageOutput struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Embeds    `json:"embeds"`
}

type Embeds struct {
	Author      `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
	Thumbnail   `json:"thumbnail"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}
