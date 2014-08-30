
# go-debug

 Conditional debug logging for Go libraries.

 The basic premise is that every library should have some form of debug logging, ideally enabled without touching code. Go-debug supports enabling and filtering these logs in real-time without reloading the program which is very useful for inspecting runtime behaviour of a production application.

 View the [docs](http://godoc.org/github.com/visionmedia/go-debug).

## Installation

```
$ go get github.com/visionmedia/go-debug
```

## Example

```go
package main

import . "github.com/visionmedia/go-debug"
import "time"

var debug = Debug("single")

func main() {
  for {
    debug("sending mail")
    debug("send email to %s", "tobi@segment.io")
    debug("send email to %s", "loki@segment.io")
    debug("send email to %s", "jane@segment.io")
    time.Sleep(500 * time.Millisecond)
  }
}
```

If you ran the program with the `DEBUG=*` environment variable you would see:

```
15:58:15.115 34us   33us   single - sending mail
15:58:15.116 3us    3us    single - send email to tobi@segment.io
15:58:15.116 1us    1us    single - send email to loki@segment.io
15:58:15.116 1us    1us    single - send email to jane@segment.io
15:58:15.620 504ms  504ms  single - sending mail
15:58:15.620 6us    6us    single - send email to tobi@segment.io
15:58:15.620 4us    4us    single - send email to loki@segment.io
15:58:15.620 4us    4us    single - send email to jane@segment.io
15:58:16.123 503ms  503ms  single - sending mail
15:58:16.123 7us    7us    single - send email to tobi@segment.io
15:58:16.123 4us    4us    single - send email to loki@segment.io
15:58:16.123 4us    4us    single - send email to jane@segment.io
15:58:16.625 501ms  501ms  single - sending mail
15:58:16.625 4us    4us    single - send email to tobi@segment.io
15:58:16.625 4us    4us    single - send email to loki@segment.io
15:58:16.625 5us    5us    single - send email to jane@segment.io
```

A timestamp and two deltas are displayed. The timestamp consists of hour, minute, second and microseconds. The left-most delta is relative to the previous debug call of any name, followed by a delta specific to that debug function. These may be useful to identify timing issues and potential bottlenecks.

## Live debugging

 A unix domain socket is created at `/tmp/debug-<pid>.sock` allowing you to
 enable, view, and disable debug output in realtime. If part of your program
 is acting up, you can `telnet` in and enable one or more `debug()` functions.

 You can enable just by typing the pattern, or disable all with "disable" or "d",
 and quit with "quit" or "q".

 Here's an example session where everything is enabled via "*", then
 all disabled with "d", followed by enabling two specific functions,
 and finally quitting with "q".

```
$ telnet /tmp/debug-15324.sock
*
16:56:54.693 71s    71s    multiple:c - doing stuff
16:56:54.786 71s    71s    multiple:b - doing stuff
16:56:54.794 101ms  101ms  multiple:c - doing stuff
16:56:54.899 104ms  104ms  multiple:c - doing stuff
16:56:55.003 104ms  104ms  multiple:c - doing stuff
16:56:55.038 252ms  252ms  multiple:b - doing stuff
16:56:55.108 104ms  104ms  multiple:c - doing stuff
16:56:55.212 104ms  104ms  multiple:c - doing stuff
16:56:55.293 254ms  254ms  multiple:b - doing stuff
16:56:55.317 104ms  104ms  multiple:c - doing stuff
16:56:55.421 104ms  104ms  multiple:c - doing stuff
16:56:55.491 72s    72s    multiple:a - doing stuff
16:56:55.526 104ms  104ms  multiple:c - doing stuff
16:56:55.548 254ms  254ms  multiple:b - doing stuff
16:56:55.630 104ms  104ms  multiple:c - doing stuff
d

multiple:a
16:57:27.580 32s    32s    multiple:a - doing stuff
16:57:28.585 1s     1s     multiple:a - doing stuff
16:57:29.586 1s     1s     multiple:a - doing stuff
16:57:30.587 1s     1s     multiple:a - doing stuff
d

multiple:b
16:57:34.953 39s    39s    multiple:b - doing stuff
16:57:35.208 254ms  254ms  multiple:b - doing stuff
16:57:35.461 252ms  252ms  multiple:b - doing stuff
16:57:35.711 250ms  250ms  multiple:b - doing stuff
16:57:35.963 251ms  251ms  multiple:b - doing stuff
16:57:36.216 253ms  253ms  multiple:b - doing stuff
16:57:36.471 255ms  255ms  multiple:b - doing stuff
16:57:36.724 252ms  252ms  multiple:b - doing stuff
16:57:36.979 255ms  255ms  multiple:b - doing stuff
16:57:37.229 250ms  250ms  multiple:b - doing stuff
16:57:37.483 254ms  254ms  multiple:b - doing stuff

q
```

 If your `telnet` doesn't support unix domain sockets you can try socat or netcat:

```
$ socat - UNIX-CONNECT:/tmp/debug-$pid.sock
$ nc -U /tmp/debug-$pid.sock
```

## The DEBUG environment variable

 Executables often support `--verbose` flags for conditional logging, however
 libraries typically either require altering your code to enable logging,
 or simply omit logging all together. go-debug allows conditional logging
 to be enabled via the __DEBUG__ environment variable, where one or more
 patterns may be specified.

 For example suppose your application has several models and you want
 to output logs for users only, you might use `DEBUG=models:user`. In contrast
 if you wanted to see what all database activity was you might use `DEBUG=models:*`,
 or if you're love being swamped with logs: `DEBUG=*`. You may also specify a list of names delimited by a comma, for example `DEBUG=mongo,redis:*`.

 The name given _should_ be the package name, however you can use whatever you like.

# License

MIT