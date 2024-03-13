package githubInterface

import "github.com/Scalingo/sclng-backend-test-v1/internal/services"

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

func GetHttpInterfaceToService(queryParams GetGithubRepositorySearchQueryParams) services.GithubRepositorySearchOptions {
	return services.GithubRepositorySearchOptions{
		Languages: queryParams.Languages,
		MinStars:  queryParams.MinStars,
		MaxStars:  queryParams.MaxStars,
		Forks:     queryParams.Forks,
		Pushed:    queryParams.Pushed,
		Topic:     queryParams.Topic,
		License:   queryParams.License,
		User:      queryParams.User,
		Org:       queryParams.Org,
		Fork:      queryParams.Fork,
		Archived:  queryParams.Archived,
		Mirror:    queryParams.Mirror,
		Limit:     queryParams.Limit,
		Sort:      queryParams.Sort,
		Order:     queryParams.Order,
	}
}

type RepositoryStats struct {
	Name       string                    `json:"name"`
	Owner      string                    `json:"owner"`
	Repository string                    `json:"repository"`
	Languages  map[string]map[string]int `json:"languages"`
}

type GetGithubRepositorySearchResponse struct {
	Repositories []RepositoryStats `json:"repositories"`
}

func GetRepositoryStatsResponse(repositories []services.GithubRepository) GetGithubRepositorySearchResponse {
	response := make([]RepositoryStats, len(repositories))
	for i, _ := range repositories {
		languages := make(map[string]map[string]int, len(repositories[i].Language))
		for key, value := range repositories[i].Language {
			languages[key] = map[string]int{"bytes": value}
		}

		response[i] = RepositoryStats{
			Name:       repositories[i].Name,
			Owner:      repositories[i].Owner,
			Repository: repositories[i].Repository,
			Languages:  languages,
		}
	}

	return GetGithubRepositorySearchResponse{Repositories: response}
}
