package messages

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"github.com/pkg/errors"
)

type (
	msgMeta struct {
		msg       Message
		locations []Location
		tmpl      *template.Template
	}
)

var (
	msgMetas = map[reflect.Type]*msgMeta{}
)

//Add bind message with specific pattern and specific sending locations
func Add(msg Message, locations ...Location) error {
	typ := messageType(msg)
	if meta, ok := msgMetas[typ]; ok {
		return errors.Errorf("duplicate with name=%s", meta.msg.Name())
	}

	tmpl, err := template.New(fmt.Sprintf("%d", len(msgMetas))).Parse(msg.Template())
	if err != nil {
		return errors.WithMessage(err, "bad template")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, msg); err != nil {
		return errors.WithMessage(err, "failed to generate tmpl message")
	}
	msgMetas[typ] = &msgMeta{
		msg:       msg,
		locations: locations,
		tmpl:      tmpl,
	}
	return nil
}

//Send msg
func Send(msg Message) error {
	typ := messageType(msg)
	meta, ok := msgMetas[typ]
	if !ok {
		return errors.Errorf("unsupport msg name=%s", msg.Name())
	}

	var buf bytes.Buffer
	if err := meta.tmpl.Execute(&buf, msg); err != nil {
		return errors.WithMessage(err, "generate text fail")
	}

	text := buf.String()
	for _, l := range meta.locations {
		l.Send(msg, text)
	}
	return nil
}

func messageType(msg Message) reflect.Type {
	val := reflect.Indirect(reflect.ValueOf(msg))
	return val.Type()
}
