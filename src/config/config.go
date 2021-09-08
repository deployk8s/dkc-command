package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg"
)

var (
	//download
	Kconfig struct {
		Mirror      string
		DownloadDir string
		OnlyOne     string
		CheckMD5    bool
		Cache       bool

		InventoryFile string
		//Docker        bool
		//Ansible       bool
		Show bool

		//Mongo     bool
		//K8s       bool
		UseDocker bool
		//Ops       bool
		ExtraArgs string

		Port int

		TemplateDir    string
		Template       string
		DockerUser     string
		DockerPassword string
		Fool           bool
		ChartRepo      string
		ChartVersion   string
		GenerateOnly   bool
		OpsappsOnly bool

		Hostname            string
		Ip                  string
		OnlyUpdateInventory bool

		RemoveData bool
	}
)

func InitDownload() {
	initGlobal()
	Kconfig.Mirror = strings.TrimSpace(viper.GetString("mirror"))
	Kconfig.OnlyOne = viper.GetString("only-one")
	Kconfig.CheckMD5 = viper.GetBool("check-md5")
	Kconfig.Cache = viper.GetBool("use-cache")
}

func InitDownloadSurpass() {
	initGlobal()

	Kconfig.DockerUser = viper.GetString("docker-user")
	Kconfig.DockerPassword = viper.GetString("docker-password")
	Kconfig.OnlyOne = viper.GetString("only-one")
	Kconfig.Fool = viper.GetBool("fool")
	Kconfig.ChartRepo = strings.TrimSpace(viper.GetString("chart-repo"))
	Kconfig.GenerateOnly = viper.GetBool("generate-only")
	Kconfig.OpsappsOnly = viper.GetBool("opsapps")
	Kconfig.TemplateDir = viper.GetString("template-dir")
	//Kconfig.ChartVersion = viper.GetString("chart-version")
}

func initGlobal() {
	Kconfig.InventoryFile = viper.GetString("inventory-file")
	Kconfig.DownloadDir = "."
	Kconfig.Template = ""
	Kconfig.TemplateDir = "./"
	if viper.GetBool("debug") {
		log.Log.SetLevel(logrus.DebugLevel)
		log.Log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	}
}
func InitPrepare() {
	initGlobal()
}
func InitWeb() {
	initGlobal()
	Kconfig.Mirror = viper.GetString("mirror")
	Kconfig.Port = viper.GetInt("port")
}
func InitInstall() {
	initGlobal()
	Kconfig.UseDocker = false
	Kconfig.ExtraArgs = viper.GetString("extra-args")
}
func InitStatus() {
	initGlobal()
}
func InitUninstall() {
	initGlobal()
}
func InitAddNode() {
	initGlobal()
	Kconfig.Hostname = strings.TrimSpace(viper.GetString("hostname"))
	Kconfig.Ip = strings.TrimSpace(viper.GetString("ip"))
	Kconfig.OnlyUpdateInventory = viper.GetBool("only-update-inventory")
	if Kconfig.Hostname == "" || Kconfig.Ip == "" {
		pkg.Help("node", "add")
		os.Exit(0)
	}
}
func InitRemoveNode() {
	initGlobal()
	Kconfig.Hostname = strings.TrimSpace(viper.GetString("hostname"))
	Kconfig.RemoveData = viper.GetBool("remove-data")
	Kconfig.OnlyUpdateInventory = viper.GetBool("only-update-inventory")

	if Kconfig.Hostname == "" {
		pkg.Help("node", "del")
		os.Exit(0)
	}
}
