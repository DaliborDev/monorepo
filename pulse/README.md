# Pulse
_Taken from [here](https://ieftimov.com/posts/four-steps-daemonize-your-golang-programs/)_


## Usage:
Run daemon with command:
```shell
bazel run //go-daemonize:daemonize -- --interval 10s
```

### Arguments

- `interval`
  - Configures ticking interval; controls how often code is ran
  - Values expressed in format `10s`, `10m`, `10h`, and so on...