package options

type AsanaDriver struct {
	AccessToken string
	ProjectID   string
}

func AsanaDriverFromMap(m map[string]any) *AsanaDriver {
	return &AsanaDriver{
		AccessToken: getOrEmpty(m, "token"),
		ProjectID:   getOrEmpty(m, "project"),
	}
}

func (a *AsanaDriver) Merge(b *AsanaDriver) *AsanaDriver {
	return &AsanaDriver{
		AccessToken: replaceOnEmpty(a.AccessToken, b.AccessToken),
		ProjectID:   replaceOnEmpty(a.ProjectID, b.ProjectID),
	}
}

func (a *AsanaDriver) ToMap() map[string]any {
	return map[string]any{
		"token":   a.AccessToken,
		"project": a.ProjectID,
	}
}
