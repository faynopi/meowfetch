package widgets

import (
	"fmt"
	"meowfetch/configs"

	"github.com/shirou/gopsutil/v3/mem"
)

func W_mem() string {
    var ret string
    v, _ := mem.VirtualMemory()
    ret += fmt.Sprintf("%d / %d MiB",(v.Used / 1048576), (v.Total / 1048576))

    if *configs.Memory_precent {
        ret += fmt.Sprintf(" (%.0f%%)", v.UsedPercent)
    }

    return ret
}
