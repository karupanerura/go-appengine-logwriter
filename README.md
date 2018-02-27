# go-appengine-logwriter

```go
import (
    "google.golang.org/appengine"
    "google.golang.org/appengine/log"
)

ctx := appengine.NewContext(r)
writer := aelogwriter.NewWriter(ctx, log.Debugf)
writer.Write([]byte("foo")) // "foo" is printed by log.Debugf
```
