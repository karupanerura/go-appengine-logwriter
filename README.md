# go-appengine-logwriter

```go
import (
    "io"

    "google.golang.org/appengine"
    "google.golang.org/appengine/log"
    aelogwriter "github.com/karupanerura/go-appengine-logwriter"
)

ctx := appengine.NewContext(r)
writer := aelogwriter.NewWriter(ctx, log.Debugf)
io.WriteString(writer, "foo") // "foo" is printed by log.Debugf
```
