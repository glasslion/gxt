# gxt
Simple process supervisor that retrys failed commands, written in Go.

**demo**:
![ttygif](https://github.com/glasslion/gxt/raw/master/gxt.gif)

## USAGE
```bash
# optional configuration environment variables:
export GXT_MAX_RETRY=10 # max retry times
export GXT_RETRY_WAIT=3 # seconds to wait before each retry
#-------------------------------

gxt <command_name>
```

## About Name (´・ω・`)
```go
"go", "刑天" := <-gxt
```
