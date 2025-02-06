# Receipt Processor Solution
To install and start the receipt processor server (receipt-server), instructions are provided at the
[bottom of this README](#installing).

## Language Selection
I have over 30 years of experience with C and C++.  I have similar years worth of experience with 6502ASM, x86ASM, Java, HTML,
Perl, python, and more.  I have not yet experience Go.  Your readme allows for any language, but seems to subtly imply Go is
used primarily.  I opted to teach myself Go and have a ... go ... at solving this problem therein.

Coming from a C and C++ heavy background, I used a Makefile rather than e.g., cmake, etc.

### Tools Used
I used the following:
 * go version go1.23.5 linux/amd64
 * ogen v1.9.0
 * curl v8.11.1 

### Assumptions Made
These are some of the assumptions I made while working on this project

 * According RFC3339 section 5.6 a "time" consists of 'time-hour ":" time-minute ":" time-second', and this is 
how ogen and its json implementation handle parsing times.  The examples provided only provide time-hour and time-minute.  This 
necessitated a change to receipt/oas_json_gen.go to "manually" try to parse with the full spec and then fall back to the format
provided.

 * I chose ogen to do the heavy-lifting dealing with the API.  This was based on some cursory digs into finding an API generator for 
Go, and this seemed to be fairly well supported in the communites.

 * Use HTTP/2 

### Miscellaneous
I documented some of my discoveries and notes in the attached document notes.org

I wouldn't normally leave so much debug logging, but for this product it made sense
for a couple of reasons.  First, since this is not a production machine, a request 
taking a few ms more due to logging isn't mission critical.  Second, the logging
is light, mostly used in rules to determine points.  The logging helps break up the rules
and makes it easier to track where points are erroneous.

The receipt server will accept flags. ```-debug``` which enables logging.  The flag requires a $LEVEL: "debug", "info", "warn", "error", or "none".  
It uses log/slog and follows the time-tested pattern with those options.  ```-port``` starts the server listening on a specific port.  

``` sh

$ bin/server --?
flag provided but not defined: -?
Usage of bin/server:
  -debug string
        Specify debugging level. 'debug','info','warn','error','none' (default "none")
  -port int
        Listen to specified port (default 8080)
```

To generate the environment, these steps were taken:
``` sh
$ go mod init receipt
$ go install -v github.com/ogen-go/ogen/cmd/ogen@latest
$ cp ~/api.yml 
$ echo -e "package main\n\n//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target receipt --clean api.yml" > generate.go
$ go get .
$ go generate ./...

```

In general, I would not put heavy-lifting work (such as the points calculation) within a service request.
For this specific exercise, the rules are relatively quick; however, it is not impossible to imagine a situation
where the rules become a full-fledged complex state machine also requiring external inputs.  To have a user 
wait on that is inappropriate.  Instead, the calculation should be done either during periods of low 
activity (by an asynchronous schedule, perhaps) or just-in-time should a get-points request come early.

## Installing and Running {#installing}

### Installing

``` sh

$ git clone https://github.com/skzitz/receipt-challenge
$ cd receipt-challenge
$ go get .
$ make all install

```

### Running

The server accepts two parameters, both of which have a sensible default option.  Passing any illegal option will generate
help text.

If port 8080 is not available, you may specify -port=$PORT.  Logging is generated using levels similar to syslog: -d=$LEVEL
will enable debug logging ($LEVEL must be one of 'debug', 'info', 'warn', 'error', or 'none' (the dfefault)).

``` sh
$ bin/server 
```

### Testing from the command line

The CLI test scripts provided require curl.

Start the server:

``` sh
$ bin/server --debug=debug
Receipt Server starting on : 8080
```

In another term, run one of the test scripts:

``` sh
$ tests/test1.sh
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 06 Feb 2025 16:22:22 GMT
Content-Length: 45

{"id":"51815e37-7aba-44c2-ad38-562f7ed6d56f"}

```

Verify the points generated match expectation:

``` sh
$ tests/get_points 51815e37-7aba-44c2-ad38-562f7ed6d56f
{"points":28}
```


