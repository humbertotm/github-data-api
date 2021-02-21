package domain

type User struct {
	Username     string `json:"username"`
	ExternalID   int    `json:"external_id"`
	UserURL      string `json:"user_url"`
	FollowersURL string `json:"followers_url"`
	FollowingURL string `json:"following_url"`
	ReposURL     string `json:"repos_url"`
	Type         string `json:"type"`
	SiteAdmin    bool   `json:"site_admin"`
}
