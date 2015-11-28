package app

// Msg is the struct that holds
// app messages
type Msg struct {
	Message string
	IsErr   bool
}

func (m Msg) String() string {
	return m.Message
}

// NewErrMsg returns an app Error Message
func NewErrMsg(message string) *Msg {
	return &Msg{IsErr: true, Message: message}
}

// NewMsg returns an app Error Message
func NewMsg(message string) *Msg {
	return &Msg{IsErr: false, Message: message}
}
