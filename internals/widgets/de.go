package widgets

import "os"

func W_de() string {
    var env string
    var status bool

    if env, status = os.LookupEnv("DESKTOP_SESSION"); status {
        return env
    } else if env, status = os.LookupEnv("XDG_SESSION_DESKTOP"); status {
        return env
    }

    return "None"
}
