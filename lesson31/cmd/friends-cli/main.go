package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"os"
	"skillbox/internal/models"
)

func main() {
	app := &cli.App{
		Name:    "http-client",
		Usage:   "Send CRUD http requests to the http server",
		Version: "1.0.0",
		Suggest: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server",
				Aliases: []string{"s"},
				Value:   "127.0.0.1",
				Usage:   "server IP address",
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   "8080",
				Usage:   "server's port number",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "create new person",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Required: true,
						Usage:    "person's name",
					},
					&cli.IntFlag{
						Name:     "age",
						Required: true,
						Usage:    "person's age",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := AddUserToServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get person's info",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Required: true,
						Usage:    "person's id",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := GetUserFromServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "get-all",
				Aliases: []string{"ga"},
				Usage:   "get all person records",
				Action: func(cCtx *cli.Context) error {
					err := GetAllUserFromServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d", "del"},
				Usage:   "delete person",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Required: true,
						Usage:    "person's id",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := DeleteUserFromServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update existing person",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Required: true,
						Usage:    "person's id",
					},
					&cli.StringFlag{
						Name:  "name",
						Usage: "person's name",
					},
					&cli.IntFlag{
						Name:  "age",
						Usage: "person's age",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := UpdateUserAtServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "make-friend",
				Aliases: []string{"mf"},
				Usage:   "create friendship relationship for a person",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "s",
						Required: true,
						Usage:    "source person's id",
					},
					&cli.StringFlag{
						Name:  "t",
						Usage: "target person's id",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := MakeFriendsOnServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},

			{
				Name:    "friends",
				Aliases: []string{"f"},
				Usage:   "get friends list for a person",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Required: true,
						Usage:    "person's id",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := GetFriendsFromServer(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func AddUserToServer(cCtx *cli.Context) error {
	name := cCtx.String("name")
	if name == "" {
		return errors.New("name is not provided")
	}
	age := cCtx.Int("age")
	if age == 0 {
		return errors.New("correct person's age should be provided")
	}

	user := models.UserInRequest{Name: name, Age: age}

	json_data, err := json.Marshal(user)
	if err != nil {
		explain := fmt.Sprintf("can't convert json data: %v", err)
		return errors.New(explain)
	}

	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s:%s/users/create", server, port), "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	printResponse(resp)

	return nil
}

func GetUserFromServer(cCtx *cli.Context) error {
	id := cCtx.String("id")

	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/users/%s", server, port, id))
	if err != nil {
		return errors.New(fmt.Sprintf("can't get user %d: %v", id, err))
	}
	printResponse(resp)
	return nil
}

func GetAllUserFromServer(cCtx *cli.Context) error {
	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/users/", server, port))
	if err != nil {
		return errors.New(fmt.Sprintf("can't get users: %v", err))
	}
	printResponse(resp)
	return nil
}

func DeleteUserFromServer(cCtx *cli.Context) error {
	id := cCtx.String("id")
	client := &http.Client{}

	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%s/users/%s", server, port, id), nil)
	if err != nil {
		return errors.New(fmt.Sprintf("can't get response: %v", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("can't get response: %v", err))
	}
	printResponse(resp)
	return nil
}

func UpdateUserAtServer(cCtx *cli.Context) error {
	id := cCtx.String("id")
	name := cCtx.String("name")
	age := cCtx.Int("age")
	objectID, _ := primitive.ObjectIDFromHex(id)

	u := models.User{
		ID:      objectID,
		Name:    name,
		Age:     age,
		Friends: nil,
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		return errors.New(fmt.Sprintf("can't parse user id %d: %v", id, err))
	}

	client := &http.Client{}
	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s:%s/users/%s", server, port, id), bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New(fmt.Sprintf("can't get response: %v", err))
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("can't get response: %v", err))
	}
	printResponse(resp)
	return nil
}

func MakeFriendsOnServer(cCtx *cli.Context) error {
	sourceID, _ := primitive.ObjectIDFromHex(cCtx.String("s"))
	targetID, _ := primitive.ObjectIDFromHex(cCtx.String("t"))

	request := models.FriendRequest{
		Source_id: sourceID,
		Target_id: targetID,
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		return errors.New(fmt.Sprintf("can't convert json data: %v", err))
	}

	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s:%s/friends/%s/make_friend", server, port, cCtx.String("s")), "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	printResponse(resp)
	return nil
}

func GetFriendsFromServer(cCtx *cli.Context) error {
	id := cCtx.String("id")
	server, port, err := getConnectionOpts(cCtx)
	if err != nil {
		return err
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/friends/%s", server, port, id))
	if err != nil {
		return errors.New(fmt.Sprintf("can't get user %d: %v", id, err))
	}
	printResponse(resp)
	return nil
}

func getConnectionOpts(cCtx *cli.Context) (string, string, error) {
	server := cCtx.String("server")
	port := cCtx.String("port")
	if server == "" || port == "" {
		return server, port, errors.New(fmt.Sprintf("connection options are incomplete:server %v, port %v", server, port))
	}
	return server, port, nil
}

func printResponse(resp *http.Response) {
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
