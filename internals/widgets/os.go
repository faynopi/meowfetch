package widgets

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/host"
)

func W_os() string {
    var ret string
    h, _ := host.Info()

    ret += fmt.Sprintf("%s %s %s", strings.Title(h.OS), strings.Title(h.Platform), h.KernelVersion)

    return ret
}
