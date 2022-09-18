package widgets

import (
	config "meowfetch/configs"
	"os"
	"os/exec"
	"path"
	"strings"
)

func W_shell() string {
    // default values
    var ret string
    var shell string = "sh"
    var shell_path string = "/usr/bin/sh"

    shell_path, status := os.LookupEnv("SHELL")
    if status {
        shell = path.Base(shell_path)
    }

    ret = shell

    if !*config.Shell_path {
        ret = shell
    } else {
        switch shell {
        case "bash":
            out, err := exec.Command(shell_path, "--version").Output()
            if err == nil {
                ret += " " + strings.Split(strings.Split(string(out), "\n")[0], " ")[3]
            }
        case "zsh":
            out, err := exec.Command(shell_path, "--version").Output()
            if err == nil {
                ret += " " + strings.Split(strings.Split(string(out), "\n")[0], " ")[1]
            }
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
