"""Just a plain ol' web server, with one extra MIME type."""

# TODO(twifkak): Let the user pass the path to a package, and serve only that.
# TODO(twifkak): Generate a page that includes a prefetch and link to the pkg.

import SimpleHTTPServer
import SocketServer
import signal
import sys

SimpleHTTPServer.SimpleHTTPRequestHandler.extensions_map['.wpk'] = 'application/signed-exchange;v=b0'

handler = SimpleHTTPServer.SimpleHTTPRequestHandler
# Change "localhost" to "" to serve non-local requests.
# Change 8000 to any desired port.
server = SocketServer.TCPServer(("localhost", 8000), handler)

server.serve_forever()
