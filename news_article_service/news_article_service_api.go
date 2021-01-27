package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "sajjan.com/sajjan_jyothi/news_article_service/user_service_grpc"
)

type userquery struct {
	UserID string   `json:"UserId"`
	Tags   []string `json:"Tags,omitempty"`
}

func getUserTags(userID string) ([]string, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(time.Second*5)) //5 second timeout
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //Adding one second context timeout
	defer cancel()
	r, err := c.GetUserDetails(ctx, &pb.UserRequest{UserId: userID})
	if err != nil {
		log.Printf("could not user details: %v", err)
		return nil, err
	}
	return r.GetTags(), nil
}

func newsArticleHandler(w http.ResponseWriter, r *http.Request) {
	var user userquery
	newslist := make([]database, 0)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Json parsing failed ", err)
		return
	}
	log.Println("Getting user details for ", user.UserID)
	tags, err := getUserTags(user.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("User details returned error ", err)
	}

	// TODO
	// There are optimized way to search in database using regexp etc.
	// If we use backend database (SQL/NoSQL) there searches and sanity tests can be done in much better way

	if len(user.Tags) == 0 {
		//Return all matching users subsribed tags
		for _, news := range _GetDataBase() {
			for _, usertag := range tags {
				for _, newstag := range news.Tags {
					if usertag == newstag {
						newslist = append(newslist, news)
					}
				}
			}
		}
	} else { //Something in request tag
		for _, news := range _GetDataBase() {
			for _, usertag := range user.Tags {
				for _, newstag := range news.Tags {

					if usertag == newstag {
						for _, subscribedtag := range tags {
							if subscribedtag == usertag { //Check whether use is actually sunscribed to the tag
								newslist = append(newslist, news)
							}
						}
					}
				}
			}
		}

	}

	if err := json.NewEncoder(w).Encode(&newslist); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Json encoding failed ", err)
	}

}
