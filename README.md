# Golang Technical Test

## Instructions
To start the program just run 
```
make deps
make run
```

## Usage
The API can be used with all the filters provided by the GitHub search API.

```
curl http://localhost:5000/repositories?language=go
curl http://localhost:5000/repositories?language=go,javascript
curl http://localhost:5000/repositories?language=go&min_stars=5000
```
Ordering and Size limit are also supported such as 
```
curl http://localhost:5000/repositories?language=go&min_stars=5000&limit=2
```


All the available filters can be found here :

```
type GetGithubRepositorySearchQueryParams struct {
	Languages *[]string `schema:"languages"`
	MinStars  *int      `schema:"min_stars"`
	MaxStars  *int      `schema:"max_stars"`
	Forks     *int      `schema:"forks"`
	Pushed    *string   `schema:"pushed"`
	Topic     *string   `schema:"topic"`
	License   *string   `schema:"license"`
	User      *string   `schema:"user"`
	Org       *string   `schema:"org"`
	Fork      *bool     `schema:"fork"`
	Archived  *bool     `schema:"archived"`
	Mirror    *bool     `schema:"mirror"`

	Limit *int    `schema:"limit"`
	Sort  *string `schema:"sort"`
	Order *string `schema:"order"`
}
```
## Architecture
This is a pretty simple program that I decomposed in a few parts.

```
/api contains the API request and responses Body.

/internal contains most of the application.
    /handlers contains the "http" handler for each service. It receives the structs defined by the API package and transform them to be used in the Service
    /services contains the actual business logic
```
