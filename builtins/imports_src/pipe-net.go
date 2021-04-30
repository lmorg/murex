package imports

// This is an optional builtin because there is no strict reason why you might
// want to pipe data over a TCP or UDP socket. However it's a cool feature to
// have available should you want it.
//
// For those unaware, this feature is similar to Bash the following in Bash:
//
//     echo "GET /" > /dev/tcp/google.com/80
//
// ...except murex's `net` pipe does allow for better interactivity over
// network sockets.

import _ "github.com/lmorg/murex/builtins/pipes/net" // piping data via TCP and UDP sockets
