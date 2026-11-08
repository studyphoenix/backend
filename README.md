# backend

## Environment Variables
```env
DATA_DIR=<dir> # Directory to store data (sqlite db and uploads)
FRONTEND_URL=<url to frontend>
BACKEND_URL=<url of the application>
DEFAULT_REDIRECT=<url to redirect to upon login>
TEST_MODE=true # used for running the tests via bruno
```

## Development
> [!NOTE]
> This is a MUCH more technical description than for the frontend

### API docs
API documentation is available at `<url>/swagger/index.html`, for local development typically at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Starting for backend development
```bash
./launch.sh
```
This will run it in development mode with a sqlite backend. Setting env vars will override the config file settings

### Running tests
Run the provided script in the `testing` directory after spinning up the backend using
```bash
./launch.sh testing
```
This will start the backend in testing mode (i.e. with some extra routes to facilitate the tests, such as allowing userID switching for Access Control tests)
