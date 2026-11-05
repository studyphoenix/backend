## Authentication
To log in, open [this link](/auth/v2/login) in a new tab. That will log you in automatically

## Usage of the API in the frontend
Follow the example in `src/ts/util/api-example.ts`.
You should always use my `oidc` sdk to make requests, as it can automatically handle unauthorized requests if needed, or use a nice UI handler (which has been attached to it to handle it)

## Usage of these docs
- CRUD = Create Retrieve Update Delete
- Endpoints marked with `[WIP]` are currently not completed
- The HTTP methods are displayed with colours
- parameters *not* marked with an asterisk (`*`) are optional
- You can try out queries by using the `Try it out` button when an endpoint is expanded
- The URL you should enter in the request URL on the frontend with the SDK is the one you can get from copying the endpoint
- For endpoints with placeholders (i.e. that have `{id}` in their URL) need these substituted in the URL. That's something the SDK won't do for you
