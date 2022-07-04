package global

type EnvService interface {
	GetAvailableVersions() []string
	Install() error
}
