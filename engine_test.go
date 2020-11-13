package messages

import (
	"bytes"
	"testing"
	"text/template"
)

type (
	testMsg struct {
		Str string
		Age int
	}

	testLoc struct {
		expect string
	}
)

func (testMsg) Name() string {
	return "testMsg"
}
func (testMsg) Template() string {
	return `str={{.Str}} age={{.Age}}`
}

func (tl *testLoc) Send(msg Message, text string) error {
	tmpl, _ := template.New("xx").Parse(msg.Template())
	var buf bytes.Buffer
	tmpl.Execute(&buf, msg)
	if buf.String() != text {
		panic("bad value")
	}
	return nil
}

func TestEngine(t *testing.T) {
	msg := testMsg{
		Str: "eheh",
		Age: 12,
	}
	tl := &testLoc{}
	Add(&msg, tl)
	if err := Add(msg, tl); err == nil {
		t.Fatalf("duplicate test fail")
	}

	Send(msg)
	Send(&msg)
}
