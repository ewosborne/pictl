# pictl
cli interface to pihole api


**Need to set env var PIPWHASH to WEBPASSWORD as found on your pihole in /etc/pihole/setupVars.conf or GUI Settings->API/Web Interface->Show API Token**


# TODO
* docs
* more args
* -v / -d verbose and debug flags
* cli arg to override pi.hole hostname
* error checking!
* probably a Makefile
* should probably have some tests I guess
* catch "pictl disable 42" as an incorrect use of "pictl disable -d 42"

# EXAMPLES
```
eric@air pictl % pictl status
200 OK
{"status":"disabled"}

eric@air pictl % pictl enable
200 OK
{"status":"enabled"}

eric@air pictl % pictl status
200 OK
{"status":"enabled"}

eric@air pictl % pictl disable
200 OK
{"status":"disabled"}

eric@air pictl % pictl disable -d 35
200 OK
{"status":"disabled"}
```