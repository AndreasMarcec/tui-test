package main

import (
	"io"
	"net/http"

	"github.com/rivo/tview"
)

type Tui struct {
	app   *tview.Application
	form  *tview.Form
	flex  *tview.Flex
	field *tview.TextView
}

func setup(t *Tui) {
	t.app = tview.NewApplication()
	t.form = tview.NewForm()
	t.flex = tview.NewFlex()
	t.field = tview.NewTextView().SetText("Teststring")
}

func create() *Tui {
	return new(Tui)
}

func get_url_from_body(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func get_url_headers(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	s := ""

	for k, v := range resp.Header {
		s += k
		for _, o := range v {
			s += " " + o
		}
		s += "\n"
	}
	return s
}

func main() {
	t := create()
	setup(t)

	t.flex.AddItem(t.field, 0, 1, false)

	t.form.AddInputField("Url", "https://example.com", 0, nil, nil).
		AddButton("GetBody", func() {
			// Update textfield in flex
			inputField := t.form.GetFormItem(0).(*tview.InputField).GetText()
			t.field.SetText(get_url_from_body(inputField))
		}).
		AddButton("GetHeaders", func() {
			inputField := t.form.GetFormItem(0).(*tview.InputField).GetText()
			t.field.SetText(get_url_headers(inputField))
		}).
		AddButton("Exit", func() { t.app.Stop() })

	t.flex.AddItem(t.form, 0, 1, false)
	t.form.SetBorder(true).SetTitle("Test")
	t.field.SetBorder(true)

	if err := t.app.SetRoot(t.flex, true).SetFocus(t.form).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
