# Small utils and snippets

## Cchannel 
fixed sized non-blocking buffer implemented using channels. Only last n values will be kept
```
import cchannel "github.com/a50397/utils/cchannel"

var cchan cchannel.Cchannel

cchan.NewChan(10)
cchan.Add(item)

if val, valid := cchan.Get(); valid {
    // val is interface{} type
    i := val.(itemType)
    ...
}
```

## dotenv
simple .env reader with optional defaults for values not included in the .env file. All keys will be created as uppercase. There are no checks on invalid chracters, etc.
```
import (
    "log"

    dotnev "github.com/a50397/utils/go_dotenv"
)

err := dotnev.LoadDotenv(dotnev.DefaultConf{Key: "PASSWORD", Value: "secret"}, dotnev.DefaultConf{Key: "login", Value: "admin"
)

if err {
    log.Println(".env file not found")
}
```


