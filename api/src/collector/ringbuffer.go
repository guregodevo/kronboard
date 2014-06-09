package collector

import "core"

// See http://pivotallabs.com/a-concurrent-ring-buffer-for-go/

type RingBuffer struct {
    inputChannel  <-chan core.Event
    outputChannel chan core.Event
}

func NewRingBuffer(inputChannel <-chan core.Event, outputChannel chan core.Event) *RingBuffer {
    return &RingBuffer{inputChannel, outputChannel}
}

func (r *RingBuffer) Run() {
    for v := range r.inputChannel {
        select {
	        case r.outputChannel <- v:
	        default:
	            <-r.outputChannel
	            r.outputChannel <- v
	    }
    }
    close(r.outputChannel)
}

