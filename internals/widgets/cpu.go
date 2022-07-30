package widgets

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
)

func W_cpu() string {
    cp, _ := cpu.Info()
    return fmt.Sprintf("%s", strings.Split(cp[0].ModelName, "@")[0])
}
