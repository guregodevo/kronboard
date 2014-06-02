package collector

import (
	"net/url"
	"net/http"
)

type ChannelCollectorResource struct {
    in  chan Event
    Out chan Event
    rb *RingBuffer
}

func NewChannelCollectorResource() (*ChannelCollectorResource){
	in := make(chan Event)
	Out := make(chan Event, 5)
	rb := NewRingBuffer(in, Out)
    return &ChannelCollectorResource{in, Out, rb}
}

func (api *ChannelCollectorResource) Run() {
    api.rb.Run()
}



	
func (api *ChannelCollectorResource) Post(values url.Values, event Event) (int, interface{}) {
    api.in <- event
	return http.StatusOK, nil
}
