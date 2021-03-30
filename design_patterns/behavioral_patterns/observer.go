package behavioral_patterns

import "fmt"

/* Define a one-to-many dependency between objects
so that when one object changes state,
all its dependents are notified and updated automatically
*/

type Event struct {
	info string
}

type Observer interface {
	OnNotify(Event)
}

type User struct {
	userId int
}

func (user User) OnNotify(event Event) {
	fmt.Printf("User %d received: %s\n", user.userId, event.info)
}

type Notifier interface {
	Subscribe(Observer)
	Unsubscribe(Observer)
	Alert(Event)
}

type YoutubeNotifier struct {
	subscribers map[Observer]bool
}

func (notifier *YoutubeNotifier) Subscribe(observer Observer) {
	notifier.subscribers[observer] = true
}

func (notifier *YoutubeNotifier) Unsubscribe(observer Observer) {
	delete(notifier.subscribers, observer)
}

func (notifier *YoutubeNotifier) Alert(event Event) {
	for observer := range notifier.subscribers {
		observer.OnNotify(event)
	}
}

func ObserverExample() {
	notifier := YoutubeNotifier{map[Observer]bool{}}
	user1, user2, user3 := User{1}, User{2}, User{3}
	notifier.Subscribe(user1)
	notifier.Subscribe(user3)
	notifier.Alert(Event{"New video updated - 1"})
	notifier.Alert(Event{"New video updated - 2"})
	notifier.Subscribe(user2)
	notifier.Alert(Event{"New video updated - 3"})
	notifier.Unsubscribe(user1)
	notifier.Alert(Event{"New video updated - 4"})
}
