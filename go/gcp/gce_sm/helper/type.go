package helper

// Secret holds single secret information
type Secret struct {
	Name   string            `yaml: name`
	Value  string            `yaml: value`
	Labels map[string]string `yaml: labels`
}

// SecretFile holds the whole yaml secrets data, including Project ID
type SecretFile struct {
	Project string   `yaml: project`
	Secrets []Secret `yaml: inline`
}
