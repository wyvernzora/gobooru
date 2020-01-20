package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wyvernzora/gobooru/types"
)

type Posts struct {
	BaseClient *BaseClient
}

type ListPostsInput struct {
	Limit  int
	Page   int
	Tags   []string
	Random bool
}

func (client *Posts) List(input ListPostsInput) (*types.Posts, error) {
	request, err := client.BaseClient.NewRequest("GET", "/posts.json")
	if err != nil {
		return nil, err
	}
	qs := request.URL.Query()
	if input.Limit > 0 {
		qs.Add("limit", strconv.Itoa(input.Limit))
	}
	if input.Page > 0 {
		qs.Add("page", strconv.Itoa(input.Page))
	}
	if len(input.Tags) > 0 {
		tags := strings.Join(input.Tags, " ")
		qs.Add("tags", tags)
	}
	if input.Random {
		qs.Add("random", "true")
	}
	request.URL.RawQuery = qs.Encode()

	var posts types.Posts
	err = client.BaseClient.ExecuteRequest(request, &posts)

	return &posts, err
}

func (client *Posts) GetById(postID types.PostID) (*types.Post, error) {
	path := fmt.Sprintf("/posts/%s.json", strconv.Itoa(int(postID)))
	request, err := client.BaseClient.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var post types.Post
	err = client.BaseClient.ExecuteRequest(request, &post)

	return &post, err
}
