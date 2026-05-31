package localization

type Language string

const (
	Russian Language = "ru"
	English Language = "en"
)

type Key string

const (
	KeyAppTitle           Key = "app.title"
	KeyTabDocker          Key = "tab.docker"
	KeyTabSystemd         Key = "tab.systemd"
	KeyTabMetrics         Key = "tab.metrics"
	KeyDockerPlaceholder  Key = "docker.placeholder"
	KeySystemdPlaceholder Key = "systemd.placeholder"
	KeyMetricsPlaceholder Key = "metrics.placeholder"
	KeyHelp               Key = "help"
)
