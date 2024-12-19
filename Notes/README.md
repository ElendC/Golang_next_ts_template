# Backend
## Project structure
### Handlers
- Similar to controllers in other languages. 
- Handles http requests and responses.
- Each handler processes specific routes `/home, /about`, performs logic, and returns data (e.g., HTML, JSON).

Each handler has its own folder inside the handler directory. 
The home folder should contain a `HomeHandler.go` that handles handler logic, 
and other functions in separate files.

### Models
- Represent data structures ("Objects") in the application.
- Defines shape of data like: `Page`, `User` etc.
- Also used for database storage, API responses, or shared data across handlers.

### main.go
- Entry point of the application.
- Run this to start the backend server.

### go.mod
- Defines the module name and manages dependencies
- Tracks external packages and their versions
- Created with the cmd: `go mod init <module-name>`
