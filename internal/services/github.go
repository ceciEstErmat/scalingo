package services

import (
	"context"
	"fmt"
	"strings"
	"sync"

	sdk "github.com/google/go-github/v60/github"
)

type GithubRepositorySearchOptions struct {
	Languages *[]string
	MinStars  *int
	MaxStars  *int
	Forks     *int
	Pushed    *string
	Size      *int
	Topic     *string
	License   *string
	User      *string
	Org       *string
	Fork      *bool
	Archived  *bool
	Mirror    *bool

	Limit *int
	Sort  *string
	Order *string
}

type GithubRepository struct {
	Name       string
	Owner      string
	Repository string
	Language   map[string]int
}

type Github struct {
	client *sdk.Client
}

func InitGithub(client *sdk.Client) Github {
	return Github{client: client}
}

func (g *Github) GetRepository(filters GithubRepositorySearchOptions) ([]GithubRepository, error) {
	query := generateSearchQuery(filters)
	if query == "" {
		return nil, fmt.Errorf("no search criteria provided")
	}

	searchOption := generateSearchOption(filters)
	repositoriesSearchResult, _, err := g.client.Search.Repositories(context.Background(), query, &searchOption)
	if err != nil {
		return nil, err
	}
	cctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	allLanguageStats := make(map[int64]map[string]int, len(repositoriesSearchResult.Repositories))
	var statsErr error

	var wg sync.WaitGroup
	for _, repository := range repositoriesSearchResult.Repositories {
		repo := repository
		wg.Add(1)

		go func() {
			defer wg.Done()
			languageStats, _, err := g.client.Repositories.ListLanguages(cctx, *repo.Owner.Login, repo.GetName())
			if err != nil {
				statsErr = err
				cancel()
				return
			}
			allLanguageStats[repo.GetID()] = languageStats
		}()
	}
	wg.Wait()

	if statsErr != nil {
		return nil, statsErr
	}

	responses := make([]GithubRepository, len(repositoriesSearchResult.Repositories))
	for i, repository := range repositoriesSearchResult.Repositories {
		responses[i] = GithubRepository{
			Name:       *repository.FullName,
			Owner:      *repository.Owner.Login,
			Repository: *repository.Name,
			Language:   allLanguageStats[repository.GetID()],
		}
	}
	return responses, nil
}

func generateSearchOption(opts GithubRepositorySearchOptions) sdk.SearchOptions {
	limit := 100
	if opts.Limit != nil {
		limit = *opts.Limit
	}

	sort := "created"
	if opts.Sort != nil {
		sort = *opts.Sort
	}

	order := "desc"
	if opts.Order != nil {
		order = *opts.Order
	}

	return sdk.SearchOptions{
		Sort:  sort,
		Order: order,
		ListOptions: sdk.ListOptions{
			Page:    1,
			PerPage: limit,
		},
	}
}

func generateSearchQuery(opts GithubRepositorySearchOptions) string {
	var filters []string

	if opts.Languages != nil && len(*opts.Languages) > 0 {
		filters = append(filters, "language:"+strings.Join(*opts.Languages, ","))
	}

	if opts.MinStars != nil {
		filters = append(filters, "stars:>="+fmt.Sprint(*opts.MinStars))
	}

	if opts.MaxStars != nil {
		filters = append(filters, "stars:<="+fmt.Sprint(*opts.MaxStars))
	}

	if opts.Forks != nil {
		filters = append(filters, "forks:>="+fmt.Sprint(*opts.Forks))
	}

	if opts.Pushed != nil {
		filters = append(filters, "pushed:"+*opts.Pushed)
	}

	if opts.Size != nil {
		filters = append(filters, "size:>="+fmt.Sprint(*opts.Size))
	}

	if opts.Topic != nil {
		filters = append(filters, "topic:"+*opts.Topic)
	}

	if opts.License != nil {
		filters = append(filters, "license:"+*opts.License)
	}

	if opts.User != nil {
		filters = append(filters, "user:"+*opts.User)
	}

	if opts.Org != nil {
		filters = append(filters, "org:"+*opts.Org)
	}

	if opts.Fork != nil && *opts.Fork {
		filters = append(filters, "fork:true")
	}

	if opts.Archived != nil && *opts.Archived {
		filters = append(filters, "archived:true")
	}

	if opts.Mirror != nil && *opts.Mirror {
		filters = append(filters, "mirror:true")
	}

	query := strings.Join(filters, " ")

	return query
}
