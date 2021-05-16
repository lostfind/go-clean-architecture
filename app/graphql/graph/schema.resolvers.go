package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"
	"zipcode/app/graphql/graph/generated"
	graphmodel "zipcode/app/graphql/graph/model"
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/infrastructure/database"
)

func (r *queryResolver) Prefectures(ctx context.Context) ([]*graphmodel.Prefecture, error) {
	db := database.NewMySqlDB()
	repository := repositories.NewDbRepository(db)
	useCase := usecase.NewAddressFinder(repository)

	prefs, _ := useCase.GetPrefectures()

	var prefectures []*graphmodel.Prefecture
	for _, pref := range prefs {
		prefectures = append(prefectures, &graphmodel.Prefecture{
			ID:        strconv.FormatInt(pref.ID, 10),
			Name:      pref.Name,
			Nameroman: pref.NameRoman,
		})
	}

	return prefectures, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
