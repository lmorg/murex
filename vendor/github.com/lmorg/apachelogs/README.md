# apachelogs
Go package for parsing Apache logs.

## Change Log

### Version 3.0.0

_STABLE!_

This is now a stable release. Unfortunately it's also quite a rewrite to made the code more idiomatic Go. A lot of that is in the form of variables and constants being renamed - which will break compatibility. I have also changed the naming of the AccessLog / AccessLogs (for slices) to be more meaningful: `AccessLine` is a single entity, `AccessLog` is a slice of AccessLine's (ie `[]AccessLine`).

_New in version 3:_

Support for error logs. Currently untested but the API follows the same format as for access log parsing so I cannot envisage any breaking changes as that code matures.  

_Usage:_

In time, more useful guides will be included in this readme. However for now see Godocs[1] for code documentation and Firesword[2] for a working implementation.

[1] https://godoc.org/github.com/lmorg/apachelogs
[2] https://github.com/lmorg/firesword