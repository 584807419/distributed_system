package registry

type ServiceName string

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName // 当前服务依赖的其他服务
	ServiceUpdateURL string
}

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
