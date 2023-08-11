# GoHST
## Go HTTP Server Template

This is the usual template I follow when all I need is to have a quick "just works" web server with the ability to extend it further when needed.

Your code goes to the `src`; your HTML pages go the `pages`; your JS files go to the `scripts`; your static content goes to `static`.

Code directories to consider:

- `conf` is a configuration file for the server to use. You can specify a port to work on, base directory where to look pages in and paths to the SSL keys. In larger projects it is bound to contain much more.
- `db` is where the database logic resides, base structures and helper functions to extract/add them from/to the database. 
- `encryption` is just a bunch of helper functions to deal with BASE64 or SHA
- `logger` is for as-basic-as-it-gets logger usage
- `server/api` is where I usually write API|Page-specific handlers
- `server/page` contains a helper function to merge `pages/base.html` and any other page together
- `server/server` is a glue between everything there is. The actual server stuff is happening there

With some work it's possible to strip everything unneeded and to just have a static web server.

There is a `modernc.org/sqlite` dependency which is there only for "compile, run and see" ability. Otherwise, even if you don't use a database it won't launch. It is up to you to replace it with another driver or if you don't need it, strip `src/db` completely.

# License
Do What The Fuck You Want To Public License for all files except for, obviously, bootstrap, as well as for src/logger which is under MIT 