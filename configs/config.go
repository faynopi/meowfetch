package configs

import (
	"os"
	"path"
	"runtime"

	_ "embed"

	config "github.com/stvp/go-toml-config"
)

var (
    xdg_confdir, XdgConfHomeExist = os.LookupEnv("XDG_CONFIG_HOME")
    homedir, _          = os.UserHomeDir()

    // Config values
    Widgets             = config.String("core.widgets", "os cpu mem de colors")
    Mode                = config.String("core.mode", "color")
    Color               = config.Int("core.color", 5)
    BorderStyle         = config.String("core.border_style", "round")
    BorderMode          = config.String("core.border_mode", "widgets")
    Separator           = config.String("core.separator", ">")
    AlignSeparator      = config.Bool("core.align_separator", true)
    Uppercase_name      = config.Bool("core.uppercase_name", true)
    Bold_name           = config.Bool("core.bold_name", true)
    Bold_art            = config.Bool("core.bold_art", true)

    Colorblock_mode     = config.String("core.blockmode", "circle")

    // Formatting
    Uptime_short        = config.Bool("opts.uptime_short", false)
    Memory_precent      = config.Bool("opts.memory_precent", false)
    Shell_path          = config.Bool("opts.shell_path", false)
    Shell_version       = config.Bool("opts.shell_version", true)
)

func ParseConfig() {
    ConfPath := GetTarget()

    if !(len(ConfPath) == 0) {
        err := config.Parse(ConfPath)
        if err != nil {
            panic(err)
        }
    }

}

func GetTarget() string {
    if XdgConfHomeExist {
        XdgConfPath := path.Join(xdg_confdir, "meowfetch", "meow.conf")
        if checkFileExist(XdgConfPath) {
            return XdgConfPath
        }
    }

    HomeConfigPath := path.Join(homedir, ".meow.conf")
    if checkFileExist(HomeConfigPath) {
        return HomeConfigPath
    }

    LinuxFallbackConf := "/etc/meow.conf"
    if runtime.GOOS == "linux" && checkFileExist(LinuxFallbackConf) {
        return HomeConfigPath
    }

    return ""
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkFileExist(p string) bool {
    if _, err := os.Stat(p); err != nil { return false }
    return true
}
