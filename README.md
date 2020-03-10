# timerange
Time toolkit for golang

## Install
```
go get -u github.com/mef13/timerange
```

## Usage
Calculating time range based on current time:
```go
package main

import (
    "fmt"
    "github.com/mef13/timerange"
)

func main() {
    daysInMonth := timerange.Now().InMonth().ToDays()
    secondsInQuarter := timerange.Now().InQuarter().ToSeconds()
    fmt.Println(daysInMonth)
    fmt.Println(secondsInQuarter)
}
```

Calculating time range based on custom time:
```go
package main

import (
    "fmt"
    "github.com/mef13/timerange"
    "time"
)

func main() {
    customTime := time.Date(2020, 2, 3, 2, 1, 10, 9, time.UTC)
    secondsInMonth := timerange.FromTime(customTime).InMonth().ToSeconds()
    fmt.Println(secondsInMonth)
}
```

##Contributing

Contributions are always welcome!

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
