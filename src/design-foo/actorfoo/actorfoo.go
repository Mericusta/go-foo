package actorfoo

// actor
// actor needs to know who is communicating with me

type Actor interface {
	Start()                // on actor start
	Do()                   // handle
	Stop()                 // on actor stop
	RecvMailBox() *MailBox // mail box
}

type actor struct {
	recvMailBox *MailBox
}

// implement interface Actor.Start
func (a *actor) Start() {}

// implement interface Actor.Do
func (a *actor) Do() {}

// implement interface Actor.Stop
func (a *actor) Stop() {}

// implement interface Actor.Start
func (a *actor) RecvMailBox() *MailBox {
	return a.recvMailBox
}

// constructor
func newActor() Actor {
	return &actor{}
}

// methods
func (a *actor) Bind(otherActor Actor) {

}

// mailbox
// - unlimited size
// - non-block for sender

type MailBox struct {
	list *ListNode // TODO: use lock free queue
}

type ListNode struct {
	next *ListNode
}

// ----------------------------------------------------------------

func actorfoo() {
	// query actor will produce SQL query periodically
	// query actor has more than one instance
	queryActor := newActor()

	// db actor will consume SQL query
	dbActor := newActor()

	// queryActor.
}
