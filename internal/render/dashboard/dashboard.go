package dashboard

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/pkg/env"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
	db     mysql.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (h *handler) View() core.HandlerFunc {
	type mysqlVersion struct {
		Ver string
	}

	mysqlVer := new(mysqlVersion)
	if h.db != nil {
		h.db.GetDbR().Raw("SELECT version() as ver").Scan(mysqlVer)
	}

	redisVer := ""
	if h.cache != nil {
		redisVer = h.cache.Version()
	}

	type viewResponse struct {
		MemTotal       string
		MemUsed        string
		MemUsedPercent float64

		DiskTotal       string
		DiskUsed        string
		DiskUsedPercent float64

		HostOS   string
		HostName string

		CpuName        string
		CpuCores       int32
		CpuUsedPercent float64

		GoPath      string
		GoVersion   string
		Goroutine   int
		ProjectPath string
		Env         string
		Host        string
		GoOS        string
		GoArch      string

		ProjectVersion string
		MySQLVersion   string
		RedisVersion   string
	}

	return func(ctx core.Context) {
		memInfo, _ := mem.VirtualMemory()
		diskInfo, _ := disk.Usage("/")
		hostInfo, _ := host.Info()
		cpuInfo, _ := cpu.Info()
		cpuPercent, _ := cpu.Percent(time.Second, false)

		obj := new(viewResponse)
		obj.MemTotal = fmt.Sprintf("%d GB", memInfo.Total/GB)
		obj.MemUsed = fmt.Sprintf("%d GB", memInfo.Used/GB)
		obj.MemUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", memInfo.UsedPercent), 64)

		obj.DiskTotal = fmt.Sprintf("%d GB", diskInfo.Total/GB)
		obj.DiskUsed = fmt.Sprintf("%d GB", diskInfo.Used/GB)
		obj.DiskUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskInfo.UsedPercent), 64)

		obj.HostOS = fmt.Sprintf("%s(%s) %s", hostInfo.Platform, hostInfo.PlatformFamily, hostInfo.PlatformVersion)
		obj.HostName = hostInfo.Hostname

		if len(cpuInfo) > 0 {
			obj.CpuName = cpuInfo[0].ModelName
			obj.CpuCores = cpuInfo[0].Cores
		}

		if len(cpuPercent) > 0 {
			obj.CpuUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuPercent[0]), 64)
		}

		obj.GoPath = runtime.GOROOT()
		obj.GoVersion = runtime.Version()
		obj.Goroutine = runtime.NumGoroutine()
		dir, _ := os.Getwd()
		obj.ProjectPath = strings.Replace(dir, "\\", "/", -1)
		obj.Host = ctx.Host()
		obj.Env = env.Active().Value()
		obj.GoOS = runtime.GOOS
		obj.GoArch = runtime.GOARCH
		obj.ProjectVersion = configs.ProjectVersion
		obj.MySQLVersion = mysqlVer.Ver
		obj.RedisVersion = redisVer

		ctx.HTML("dashboard", obj)
	}
}
