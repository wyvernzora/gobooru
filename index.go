package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wyvernzora/gobooru/client"
)

func main() {

	host := os.Getenv("DANBOORU_HOST")
	user := os.Getenv("DANBOORU_USER")
	key := os.Getenv("DANBOORU_KEY")

	baseClient, err := client.NewBaseClient(host, user, key)
	if err != nil {
		log.Fatal(err)
	}

	postsClient := &client.Posts{
		BaseClient: baseClient,
	}

	posts, err := postsClient.List(client.ListPostsInput{
		Limit: 10,
		Page:  1,
		Tags:  []string{"rating:s"},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range *posts {
		fmt.Println(post.PostID)
	}

	post, err := postsClient.GetById(3754446)
	if err != nil {
		log.Fatal(err)
	}
	if post != nil {
		fmt.Println(post.TagString)
	} else {
		fmt.Println("No post returned")
	}

}
