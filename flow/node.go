package flow

// Node going flow node
type Node interface {
	GetID() string
	GetChannelIn() *Channel
	GetChannelOut() []*Channel
	SetChannelIn(chanIn *Channel) error
	SetChannelOut(chanOut *Channel) error
}
