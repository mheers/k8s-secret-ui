package settings

type Settings struct {
	AllowedNamespaces []string `json:"allowedNamespaces"`
}

func GetSettings() Settings {
	return Settings{
		AllowedNamespaces: []string{"default"},
	}
}
