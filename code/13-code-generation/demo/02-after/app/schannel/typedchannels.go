
// Code generated DO NOT EDIT

package schannel



type IntSmartChannel struct {
	smartChannel
}
func (sc IntSmartChannel) Send(msg int) error {
	return sc.smartChannel.Send(msg)
}
func (sc IntSmartChannel) Receive() (int, bool) {
	if msg, ok := sc.smartChannel.Receive; ok {
		result, ok := msg.(int)
		return result, ok
	} else {
		var result int
		return result, ok
	}
}


type StringSmartChannel struct {
	smartChannel
}
func (sc StringSmartChannel) Send(msg string) error {
	return sc.smartChannel.Send(msg)
}
func (sc StringSmartChannel) Receive() (string, bool) {
	if msg, ok := sc.smartChannel.Receive; ok {
		result, ok := msg.(string)
		return result, ok
	} else {
		var result string
		return result, ok
	}
}


