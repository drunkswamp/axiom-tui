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
	KeyDockerErrorFormat  Key = "docker.error.format"
	KeyDockerLoading      Key = "docker.loading"
	KeyDockerEmpty        Key = "docker.empty"
	KeyDockerTableHeader  Key = "docker.table.header"
	KeySystemdPlaceholder Key = "systemd.placeholder"
	KeySystemdErrorFormat Key = "systemd.error.format"
	KeySystemdLoading     Key = "systemd.loading"
	KeySystemdEmpty       Key = "systemd.empty"
	KeySystemdTableHeader Key = "systemd.table.header"
	KeyMetricsPlaceholder Key = "metrics.placeholder"
	KeyMetricsErrorFormat Key = "metrics.error.format"
	KeyMetricsLoading     Key = "metrics.loading"
	KeyMetricsSummary     Key = "metrics.summary"
	KeyHelp               Key = "help"
)
