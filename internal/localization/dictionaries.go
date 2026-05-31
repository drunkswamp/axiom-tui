package localization

var dictionaries = map[Language]map[Key]string{
	Russian: {
		KeyAppTitle:           "axiomtui — локальная панель управления",
		KeyTabDocker:          "Docker",
		KeyTabSystemd:         "Systemd",
		KeyTabMetrics:         "Metrics",
		KeyDockerPlaceholder:  "Контейнеры Docker будут отображаться здесь",
		KeySystemdPlaceholder: "Сервисы systemd будут отображаться здесь",
		KeyMetricsPlaceholder: "Метрики CPU/RAM/Disk будут отображаться здесь",
		KeyHelp:               "Tab/←/→: переключение вкладок • Mouse: выбор вкладки • q / Ctrl+C: выход",
	},
	English: {
		KeyAppTitle:           "axiomtui — local control plane",
		KeyTabDocker:          "Docker",
		KeyTabSystemd:         "Systemd",
		KeyTabMetrics:         "Metrics",
		KeyDockerPlaceholder:  "Docker containers will be displayed here",
		KeySystemdPlaceholder: "systemd services will be displayed here",
		KeyMetricsPlaceholder: "CPU/RAM/Disk metrics will be displayed here",
		KeyHelp:               "Tab/←/→: switch tabs • Mouse: select tab • q / Ctrl+C: quit",
	},
}
