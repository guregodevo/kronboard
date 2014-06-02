package collector

import (
	//"code.google.com/p/goprotobuf/proto"	
	"bytes"
	"encoding/gob"
)

type EventEncodeDecoder struct {

}

func (d *EventEncodeDecoder) Encode(event Event) (error, []byte) {
	var network bytes.Buffer
	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	errEnc := enc.Encode(event)
	if errEnc != nil {
		return errEnc, nil
	}
	return nil, network.Bytes() 
}

func (d *EventEncodeDecoder) Decode(data []byte) (error, Event) {
	var event Event
	// Create an encoder and send a value.
	network := bytes.NewBuffer(data)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&event)
	if err != nil {
		return err, nil
	}
	return nil, event
}

