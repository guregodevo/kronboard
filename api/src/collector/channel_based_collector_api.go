package collector

import (
	"net/url"
	"net/http"
	"sync"
)

type ChannelCollectorResource struct {
    in  chan Event
    Out chan Event
    rb *RingBuffer
	mu  *sync.Mutex    
}

func NewChannelCollectorResource() (*ChannelCollectorResource){
	in := make(chan Event)
	Out := make(chan Event, 50)
	rb := NewRingBuffer(in, Out)
    return &ChannelCollectorResource{in, Out, rb, &sync.Mutex{}}
}

func (api *ChannelCollectorResource) Run() {
    api.rb.Run()
}


func (api *ChannelCollectorResource) Post(values url.Values, event Event) (int, interface{}) {
    api.mu.Lock()
    api.in <- event
    defer api.mu.Unlock()
	return http.StatusOK, nil
}
