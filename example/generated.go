package example

import (
	"context"

	"github.com/Khan/genql/graphql"
)

type getViewerResponse = struct {
	Viewer struct {
		MyName *string
	} `json:"viewer"`
}

// TODO
func getViewer(ctx context.Context, client *graphql.Client) (*getViewerResponse, error) {
	var retval getViewerResponse
	err := client.MakeRequest(ctx, `
query getViewer {
	viewer {
		MyName: name
	}
}
`, &retval, nil)
	return &retval, err
}

type getUserResponse = struct {
	User *struct {
		TheirName *string `json:"theirName"`
	} `json:"user"`
}

// TODO
func getUser(ctx context.Context, client *graphql.Client, login string) (*getUserResponse, error) {
	variables := map[string]interface{}{
		"login": login,
	}

	var retval getUserResponse
	err := client.MakeRequest(ctx, `
query getUser ($login: String!) {
	user(login: $login) {
		theirName: name
	}
}
`, &retval, variables)
	return &retval, err
}
