Linkchecker
===========

`linkchecker` will extract urls from all files in the working
directory. It will request each url once and check if the HTTP status
code is one of the success codes. All non-success codes will be
printed to the terminal.

Usage
-----

    $ linkchecker

Future work
-----------

Tests should be added (pkg net/http/httptest). Error handling should
improve (use get if head is not supported; add an error to the link
struct instead of just printing it). Requesting the http head for the
first links could be done parallel to extracting more links.

### Possible Commandline Options
Print all links. Print only error links. Print output as JSON/CSV
data. Only check files matching a regex. Use different path then
current working directory.
