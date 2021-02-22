package domain

type Repo struct {
	ExternalID      int    `json:"external_id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Owner           *User  `json:"owner"`
	HTMLUrl         string `json:"html_url"`
	URL             string `json:"url"`
	ContributorsURL string `json:"contributors_url"`
	IssuesURL       string `json:"issues_url"`
	LanguagesURL    string `json:"languages_url"`
}
