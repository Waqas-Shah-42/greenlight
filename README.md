# greenlight

Start
- ```go run ./cmd/api```
## Endpoints
___
| Method | URL Pattern     | handler            | Action                                 | Status               |
| :---   | :---            | :---               | :---                                   | :---                 |
| GET    | /v1/healthcheck | listMovieshandler  | Show application information           | - [x] Done             |
| GET    | /v1/movies      | listMoviesHandler  | Show the details of all movies         | [-] Work in progress |
| POST   | /v1/movies	   | createMovieHandler | Create a new movie                     | [-] Work in progress |
| GET    | /v1/movies/:id  | showMovieHandler   | Show the details of a specific movie   | [-] Work in progress |
| PUT    | /v1/movies/:id  | editMovieHandler   | Update the details of a specific movie | [-] Work in progress |
| DELETE | /v1/movies/:id  | deleteMovieHandler | Delete a specific movie                | [-] Work in progress |
