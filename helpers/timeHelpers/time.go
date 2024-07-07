package timehelpers

import "time"


func TimeIn(name string) time.Time {
    loc, err := time.LoadLocation(name)
    if err != nil {
        panic(err)
    }
    return time.Now().In(loc)
}
