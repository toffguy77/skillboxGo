package storage

import (
	"context"
	"fmt"
	"reflect"
	"skillbox/internal/models"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testRepo                         *UserRepo
	userDima, userKristina, userVova *models.User
)

func createTestDB() {
	testRepo = NewUserRepo("testdata")
	ctx := context.Background()
	if err := testRepo.collection.Drop(ctx); err != nil {
		panic(err)
	}
	userDima, _ = testRepo.Save(&models.User{
		Name:    "Dima",
		Age:     27,
		Friends: nil,
	})
	userKristina, _ = testRepo.Save(&models.User{
		Name:    "Kristina",
		Age:     22,
		Friends: nil,
	})
	fmt.Println("created users:", userDima, userKristina)
	return
}

func TestUserRepo_Get(t *testing.T) {
	createTestDB()
	type args struct {
		key primitive.ObjectID
	}
	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "get Dima",
			args: args{
				key: userDima.ID,
			},
			want: userDima,
		},
		{
			name: "get Kristina",
			args: args{
				key: userKristina.ID,
			},
			want: userKristina,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testRepo.Get(tt.args.key)
			if !reflect.DeepEqual(got.Name, tt.want.Name) && !reflect.DeepEqual(got.Age, tt.want.Age) && !reflect.DeepEqual(got.Friends, tt.want.Friends) && !got.ID.IsZero() {
				t.Errorf("UserRepo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_Save(t *testing.T) {
	type args struct {
		u *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "add Kristina twice",
			args: args{
				u: &models.User{
					Name:    " Kristina",
					Age:     18,
					Friends: nil,
				},
			},
			want:    userKristina,
			wantErr: false,
		},
		{
			name: "add Vova",
			args: args{
				u: &models.User{
					Name:    "Vova",
					Age:     14,
					Friends: nil,
				},
			},
			want: &models.User{
				Name:    "Vova",
				Age:     14,
				Friends: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			userVova, err = testRepo.Save(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(userVova.Name, tt.want.Name) && !reflect.DeepEqual(userVova.Age, tt.want.Age) && !reflect.DeepEqual(userVova.Friends, tt.want.Friends) && !userVova.ID.IsZero() {
				t.Errorf("UserRepo.Save() = %v, want %v", userVova, tt.want)
			}
		})
	}
}

func TestUserRepo_AllUsers(t *testing.T) {
	tests := []struct {
		name string
		want []models.User
	}{
		{
			name: "success",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, got := range testRepo.AllUsers() {
				if reflect.TypeOf(got) != reflect.TypeOf(models.User{}) {
					t.Errorf("UserRepo.AllUsers() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestUserRepo_Delete(t *testing.T) {
	type args struct {
		id primitive.ObjectID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete Vova",
			args: args{
				id: userVova.ID,
			},
			wantErr: false,
		},
		{
			name: "delete not found",
			args: args{
				id: primitive.NewObjectID(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testRepo.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepo_Update(t *testing.T) {
	type args struct {
		u *models.User
	}
	userDima.Age = 34
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "update Dima",
			args: args{
				u: &models.User{
					ID:  userDima.ID,
					Age: 34,
				},
			},
			want:    userDima,
			wantErr: false,
		},
		{
			name: "update not found",
			args: args{
				u: &models.User{
					ID:  primitive.NewObjectID(),
					Age: 34,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testRepo.Update(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_MakeFriend(t *testing.T) {
	type args struct {
		source *models.User
		target *models.User
	}
	userKristina.Friends = append(userKristina.Friends, userDima)
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Dima + Kristina",
			args: args{
				source: userDima,
				target: userKristina,
			},
			want:    userKristina,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testRepo.MakeFriend(tt.args.source, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.MakeFriend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.MakeFriend() = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Println(testRepo.AllUsers())
}

func TestUserRepo_GetFriends(t *testing.T) {
	type args struct {
		id primitive.ObjectID
	}
	tests := []struct {
		name    string
		args    args
		want    []models.User
		wantErr bool
	}{
		{
			name: "no friends for Dima",
			args: args{
				id: userDima.ID,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Kristina's friends",
			args: args{
				id: userKristina.ID,
			},
			want:    []models.User{*userDima},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testRepo.GetFriends(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.GetFriends() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.GetFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_DeleteFriend(t *testing.T) {
	type args struct {
		source *models.User
		target *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "no more friends for Kristina",
			args: args{
				source: userKristina,
				target: userDima,
			},
			want: &models.User{
				ID:      userKristina.ID,
				Name:    userKristina.Name,
				Age:     userKristina.Age,
				Friends: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testRepo.DeleteFriend(tt.args.source, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.DeleteFriend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) && !reflect.DeepEqual(got.Age, tt.want.Age) && !reflect.DeepEqual(got.Friends, tt.want.Friends) && !got.ID.IsZero() {
				t.Errorf("UserRepo.DeleteFriend() = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Println(testRepo.AllUsers())
}
