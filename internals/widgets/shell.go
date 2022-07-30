package widgets

import (
	config "meowfetch/configs"
	"os"
	"os/exec"
	"path"
	"strings"
)

func W_shell() string {
    var ret string

    shell_path, status := os.LookupEnv("SHELL")

    if !status {
        return ""
    }

    shell := path.Base(ret)

    ret = shell_path
    if !*config.Shell_path {
        ret = shell
    }

    ret = "sh"
    shell = "sh"
    shell_path = "/usr/bin/sh"
    if *config.Shell_version {
        switch shell {
            case "bash":
                ret += " " + os.Getenv("BASH_VERSION")
            case "zsh":
                ret += " " + os.Getenv("ZSH_VERSION")
            case "yash":
                out, err := exec.Command(shell_path, "--version").Output()
                if err == nil {
                    ret += " " + strings.Split(strings.Split(string(out), "\n")[0], " ")[4]
                }
            case "sh":
                out, err := exec.Command(shell_path, "--version").Output()
                if err == nil {
                    ret += " " + strings.Split(strings.Split(string(out), "\n")[0], " ")[3]
                }
            case "osh":
                ret += " " + os.Getenv("OIL_VERSION")
        }
    }


    return ret
}
