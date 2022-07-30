package widgets

import (
	"context"
	"fmt"
	"meowfetch/configs"

	"github.com/shirou/gopsutil/v3/host"
)

func W_uptime() string {
    up, _ := host.UptimeWithContext(context.Background())
    var ret string
    var h, m uint64 = 0, 0

    secs := up % (60 * 60 * 24)
    if h = secs / 60 / 60; h > 0 {
        ret += fmt.Sprintf("%d", h)
        if *configs.Uptime_short {
            ret += " Hours "
        } else {
            ret += "H "
        }

    }
    secs = up % (60 * 60)
    if m = secs / 60; m > 0 {
        ret += fmt.Sprintf("%d", m)
        if *configs.Uptime_short {
            ret += " Mins "
        } else {
            ret += "M "
        }

    }
    secs = up % 60
    ret += fmt.Sprintf("%d", secs)
    if *configs.Uptime_short {
        ret += " Secs "
    } else {
        ret += "S "
    }

    return ret
}
