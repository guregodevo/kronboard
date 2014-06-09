package collector

import (
	"net/url"
	"net/http"
	"sync"
	"core"
)

type ChannelCollectorResource struct {
    in  chan core.Event
    Out chan core.Event
    rb *RingBuffer
	mu  *sync.Mutex    
}

func NewChannelCollectorResource() (*ChannelCollectorResource){
	in := make(chan core.Event)
	Out := make(chan core.Event, 50)
	rb := NewRingBuffer(in, Out)
    return &ChannelCollectorResource{in, Out, rb, &sync.Mutex{}}
}

func (api *ChannelCollectorResource) Run() {
    api.rb.Run()
}


func (api *ChannelCollectorResource) Post(values url.Values, event core.Event) (int, interface{}) {
    api.mu.Lock()
    api.in <- event
    defer api.mu.Unlock()
	return http.StatusOK, nil
}
