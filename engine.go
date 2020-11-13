package messages

import (
	"bytes"
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
	msgMetas = map[string]*msgMeta{}
)

//Add bind message with specific pattern and specific sending locations
func Add(msg Message, tmplPattern string, locations ...Location) error {
	id := msg.ID()
	if meta, ok := msgMetas[id]; ok {
		return errors.Errorf("duplicate id=%s with name=%s", id, meta.msg.Name())
	}

	tmpl, err := template.New(id).Parse(tmplPattern)
	if err != nil {
		return errors.WithMessage(err, "bad template")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, msg); err != nil {
		return errors.WithMessage(err, "failed to generate tmpl message")
	}
	msgMetas[id] = &msgMeta{
		msg:       msg,
		locations: locations,
		tmpl:      tmpl,
	}
	return nil
}

//Send msg
func Send(msg Message) error {
	meta, ok := msgMetas[msg.ID()]
	if !ok {
		return errors.Errorf("unsupport msg id=%s name=%s", msg.ID(), msg.Name())
	}

	var buf bytes.Buffer
	if err := meta.tmpl.Execute(&buf, msg); err != nil {
		return errors.WithMessage(err, "generate text fail")
	}

	text := buf.String()
	for _, l := range meta.locations {
		l.Send(msg.ID(), msg.Name(), text)
	}

	return nil
}
