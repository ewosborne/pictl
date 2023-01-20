# pictl
cli interface to pihole api


**Need to set env var PIPWHASH to WEBPASSWORD in /etc/pihole/setupVars.conf**


# TODO
* docs
* more args
* -v / -d verbose and debug flags
* cli arg to override pi.hole hostname
* error checking!
* probably a Makefile

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