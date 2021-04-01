package main

import (
	"context"
	"database/database"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"proto/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//this server will implement server interface generate by service.proto file
type server struct {
}

func (s *server) GetUsers(ctx context.Context, e interface{}) (*proto.UserListResponse, error) {
	db, err := database.GetDatabase()
	if err != nil {
		panic(err.Error())
	}

	query, err := db.Query("select * from user")
	if err != nil {
		panic(err.Error())
	}

	res := []*proto.UserRequest{}
	for query.Next() {
		var u *proto.UserRequest
		err := query.Scan(&u.Id, &u.Fname, &u.City, &u.Phone, &u.Height, &u.Married)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, u)
	}
	return &proto.UserListResponse
}

func (s *server) GetUser(ctx context.Context, request *proto.UserRequest) (response *proto.UserResponse, err error) {
	db, err := database.GetDatabase()
	user_id := request.GetId()
	fmt.Println(user_id)
	if err != nil {
		panic(err.Error())
	}
	query, err := db.Query("SELECT * FROM user WHERE id=?", user_id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Query", query)
	res := User{}
	for query.Next() {
		var user User
		err := query.Scan(&user.Id, &user.Fname, &user.City, &user.Phone, &user.Height, User.Married)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Id, user.Fname, user.City, user.Phone, user.Height, user.Married)
		res = user
	}
	fmt.Println(query.Next())
	return &proto.UserResponse
}

func (s *server) AddUser(ctx context.Context, request *proto.UserRequest) (*proto.UserStringResponse, error) {
	db, err := database.GetDatabase()
	user_id := request.GetId()
	user_fname := request.GetFname()
	user_city := request.GetCity()
	user_phone := request.GetPhone()
	user_height := request.GetHeight()
	user_married := request.GetMarried()
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("insert into user values(?,?,?,?,?,?)", user_id, user_fname, user_city, user_phone, user_height, user_married)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Result", result.Next())

}

func (s *server) DeleteUser(ctx context.Context, request *proto.UserRequest) (*proto.UserStringResponse, error) {
	db, err := database.GetDatabase()
	user_id := request.GetId()
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("DELETE FROM user WHERE id=?", user_id)
	if err != nil {
		panic(err.Error())
	}

}

func (s *server) UpdateUser(context.Context, *proto.UserRequest) (*proto.UserStringResponse, error) {
	panic("implement me")
}

type User struct {
	Id      int    `json:"id"`
	Fname   string `json:"fname"`
	City    string `json:"city"`
	Phone   int    `json:"phone"`
	Height  int    `json:"height`
	Married string `json:"married"`
}

func index_handle(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDatabase()
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("Select * from user where id=200")
	if err != nil {
		panic(err.Error())
	}
	res := []User{}
	for result.Next() {
		var user User
		err := result.Scan(&user.Id, &user.Fname, &user.City, &user.Phone, &user.Height, &user.Married)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Fname)
		res = append(res, user)
	}
	ress, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}
	w.Write(ress)
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
