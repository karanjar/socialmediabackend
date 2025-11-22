package notifications

import (
	"context"
	"fmt"
	"socialmediabackend/services"
	"sync"

	"github.com/google/uuid"
)

//Notification store

var Store map[uuid.UUID]chan string

//Mutex

var mux sync.Mutex

func InitNotifications() {
	Store = make(map[uuid.UUID]chan string)
}

func Register(userID uuid.UUID) {
	mux.Lock()
	defer mux.Unlock()
	if _, ok := Store[userID]; !ok {
		Store[userID] = make(chan string)

	}
}
func ListenForNotifications(ctx context.Context, userID uuid.UUID) {
	mux.Lock()
	defer mux.Unlock()
	channel, ok := Store[userID]
	if !ok {
		fmt.Printf("No notifications channel registered for user : %v", userID)
	}

	for {
		select {
		case massage := <-channel:
			fmt.Printf("New Notification received : %v", massage)

		case <-ctx.Done():
			fmt.Printf("Notification listener stopped")
			return
		}
	}

}

func NotifyUser(ctx context.Context, userId uuid.UUID, massage string) {
	fs := services.NewFriendshipService()
	friend, _ := fs.Getfriends(ctx, userId)

	mux.Lock()
	defer mux.Unlock()

	for friend = range Store {

	}
}
