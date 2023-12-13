package options

type AsanaDriver struct {
	AccessToken string
	ProjectID   string
}

func (a *AsanaDriver) Merge(b *AsanaDriver) {
	if a.AccessToken == "" {
		a.AccessToken = b.AccessToken
	}

	if a.ProjectID == "" {
		a.ProjectID = b.ProjectID
	}
}

func (a *AsanaDriver) ToMap() map[string]any {
	return map[string]any{
		"token":   a.AccessToken,
		"project": a.ProjectID,
	}
}
