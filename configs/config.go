package configs

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	_ "embed"

	config "github.com/stvp/go-toml-config"
)


type Config struct {
}


var (
    //go:embed meow.conf
    defaultConfig []byte

    xdg_confdir, xdg_confdir_status = os.LookupEnv("XDG_CONFIG_HOME")
    homedir, _          = os.UserHomeDir()

    // Config values
    Widgets             = config.String("core.widgets", "os cpu mem de color")
    Mode                = config.String("core.mode", "color")
    Color               = config.Int("core.color", 5)
    BorderStyle         = config.String("core.border_style", "round")
    BorderMode          = config.String("core.border_mode", "widgets")
    Separator           = config.String("core.separator", ">")
    AlignSeparator      = config.Bool("core.align_separator", true)
    Uppercase_name      = config.Bool("core.uppercase_name", true)
    Bold_name           = config.Bool("core.bold_name", true)
    Bold_art            = config.Bool("core.bold_art", true)

    Colorblock_mode     = config.String("core.blockmode", "diamond")

    // Formatting
    Uptime_short        = config.Bool("opts.uptime_short", false)
    Memory_precent      = config.Bool("opts.memory_precent", false)
    Shell_path          = config.Bool("opts.shell_path", false)
    Shell_version       = config.Bool("opts.shell_version", true)
)

func ParseConfig() {
    default_conf := false
    conf_path := GetTarget()
    if _, err := os.Stat(conf_path); (!(len(conf_path) > 0) || errors.Is(err, os.ErrNotExist)) {
        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("Config doesn't exist. would you like to create new config? [y/n]: ")
        resp, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        resp= strings.ToLower(strings.TrimSpace(resp))

        switch resp {
            case "y", "yes", "oh yeah":
                CreateConfig()
                conf_path = GetTarget()
            default:
                fmt.Println("Using default config.")
                default_conf = true
        }
    }

    if !default_conf {
        err := config.Parse(conf_path)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func GetTarget() string {
    if xdg_confdir_status {
        xdg_conf_path := path.Join(xdg_confdir, "meowfetch", "meow.conf")
        if _, err := os.Stat(xdg_conf_path); err == nil {
            return xdg_conf_path
        }
    }

    conf_path := path.Join(homedir, ".meow.conf")
    if _, err := os.Stat(conf_path); !errors.Is(err, os.ErrNotExist) {
        return conf_path
    } else {
        return "/etc/meow.conf"
    }
}

func CreateConfig() {
    if xdg_confdir_status {
        xdg_conf_path := path.Join(xdg_confdir, "meowfetch")
        if _, err := os.Stat(xdg_conf_path); errors.Is(err, os.ErrNotExist) {
            check(os.MkdirAll(xdg_conf_path, 0744))
        }
        check(os.WriteFile(path.Join(xdg_conf_path, "meow.conf"), defaultConfig, 0644))
        return
    }

    check(os.WriteFile(path.Join(homedir, ".meow.conf"), defaultConfig, 0644))
    return
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
