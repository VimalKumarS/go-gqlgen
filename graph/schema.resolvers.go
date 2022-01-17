package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gqlgen-todos/auth"
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/model"
	"gqlgen-todos/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	_ = auth.ForContext(ctx)
	var link model.Link
	var user model.User
	link.Address = input.Address
	link.Title = input.Title
	user.Name = "test"
	link.User = &user
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

//func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
//	var user users.User
//	user.Username = input.Username
//	user.Password = input.Password
//	user.Create()
//	token, err := jwt.GenerateToken(user.Username)
//	if err != nil{
//		return "", err
//	}
//	return token, nil
//}

//func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
//	panic(fmt.Errorf("not implemented"))
//}
//
//func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
//	panic(fmt.Errorf("not implemented"))
//}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	//panic(fmt.Errorf("not implemented"))
	var links []*model.Link
	dummyLink := model.Link{
		ID: "1",
		Title: "our dummy link",
		Address: "https://address.org",
		User: &model.User{Name: "admin"},
	}
	links = append(links, &dummyLink)
	return links,nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }




//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user model.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &model.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil{
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}


//https://dev.to/ebiken/create-a-graphql-server-with-go-3mpd