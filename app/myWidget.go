package app

import (
	"fyne.io/fyne/v2/widget"
)

type MyButton struct {
	widget.Button
}

func NewMyButton() *MyButton {
	btn := &MyButton{}
	btn.ExtendBaseWidget(btn)

	return nil
}
