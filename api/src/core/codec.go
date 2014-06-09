package core

import (
	//"code.google.com/p/goprotobuf/proto"	
	"bytes"
	"encoding/gob"
	"encoding/base64"
)

type EventEncodeDecoder struct {

}

func (d *EventEncodeDecoder) EncodeBase64(event Event) (string, error) {
	err, data := d.Encode(event)
	if err!=nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func (d *EventEncodeDecoder) DecodeBase64(str string) (Event, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err 
	}
	return d.Decode(data)
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

func (d *EventEncodeDecoder) Decode(data []byte) (Event, error) {
	var event Event
	// Create an encoder and send a value.
	network := bytes.NewBuffer(data)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

