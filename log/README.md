# Log

zap logger package

### Some Tips

- the methods only accept 2 argument a `string` and a `map[string]any`
- the first input is a static text and will **not** accept placeholders like `%s`
- if you want to pass the error you **must only** pass it under the key `error`
- you don't need to call the .Error() method while passing the error
- for correlating the log with its trace just pass the fetched trace ID under the key `trace_id`
- the methods for log are: `debug`,`info`,`error`,`fatal`.
- you almost never need to use the `fatal` method.

### Config

- the minimum log level and the log format can be configured.
- the default log minimum level is `info`, can be changed by setting `LOG_LEVEL=debug`
- the default log format is `json`, can be only changed string by setting `LOG_FORMAT=string`

### Usage

```go
import (
    "context"
    log "github.com/p3ym4n/shared/log/v1"
)

// add this on the starting lines of your service boot
logger, syncer := log.NewZapFromEnv()
defer syncer()

// if you want to pass the error you must only pass it under the key error
err := errors.New("inner details")

// for correlating the logs and traces
traceID := span.Context().TraceID()

// be aware that the log message is static and will not accept placeholder like %s
c.logger.Info("the log message", map[string]any{
    "error": err,
    "trace_id": traceID, // for correlating the log with its trace
    "key3": "value3",
    "key4.nested": 12 // by adding the dot you can add nested keys
})
```
