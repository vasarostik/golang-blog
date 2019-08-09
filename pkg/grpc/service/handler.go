package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type BlogItem struct {
	UserID 	 int 				`json:"id"`
	Content  string             `json:"content"`
	Title    string             `json:"title"`
}

type PostResp struct {

}


func New( red *redis.Client) *Server {
	return &Server{redisCon: red}
}

type Server struct {
	redisCon *redis.Client
}


func (s *Server) Create(ctx context.Context, in *Post) (*Response,error) {
	var post = new(BlogItem)
	var resp = Response{Code:0}

	if err := json.Unmarshal(in.Data, &post); err != nil {
		panic(err)
	}else{
		//usrM := structs.Map(marshaled)
		//sd := s.redisCon.HMSet("post:"+strconv.Itoa(marshaled.UserID),usrM)

		_ = s.redisCon.ZAdd(strconv.Itoa(post.UserID),
			&redis.Z{Score:float64(time.Now().Unix()),Member:in.Data})
		resp = Response{Code: 200}
		//log.Println(sd)
	}

	return &resp,nil
}

func (s *Server) List(ctx context.Context, in *Request) (*PostList,error) {
	var p = new(PostList)

	posts, err := s.redisCon.ZRange(strconv.Itoa(int(in.UserID)), 0, -1).Result()

	if err != nil {
		panic(err)
	}

	p.Posts = posts

	return p,nil
}


