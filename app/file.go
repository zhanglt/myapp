package app

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

type FileChoice struct {
	widget.BaseWidget
	path                               binding.String
	labelText, placeHolder, buttonText string
}

func (fc FileChoice) Text() string {
	text, _ := fc.path.Get()
	return text
}

func NewFileChoice(labelText, placeHolder, buttonText string) *FileChoice {
	f := &FileChoice{}
	f.path = binding.NewString()
	f.labelText = labelText
	f.placeHolder = placeHolder
	f.buttonText = buttonText
	f.ExtendBaseWidget(f)
	return f
}

type fileChoiceRender struct {
	fc     *FileChoice
	label  *widget.Label
	entry  *widget.Entry
	button *widget.Button
}

func (fcr fileChoiceRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{fcr.label, fcr.entry, fcr.button}
}

func (fcr fileChoiceRender) MinSize() fyne.Size {
	var minHeight float32
	var totalWidth float32
	for _, item := range fcr.Objects() {
		if item.MinSize().Height > minHeight {
			minHeight = item.MinSize().Height
		}
		totalWidth += item.MinSize().Width
	}
	return fyne.NewSize(totalWidth, minHeight)
}

func (fcr fileChoiceRender) Layout(size fyne.Size) {
	topLeft := fyne.NewPos(0, 0)
	fcr.label.Move(topLeft)
	fcr.label.Resize(fyne.NewSize(fcr.label.MinSize().Width, size.Height))
	topLeft = topLeft.Add(fyne.NewPos(fcr.label.MinSize().Width, 0))
	fcr.entry.Move(topLeft)
	fcr.entry.Resize(fyne.NewSize(size.Width-fcr.label.MinSize().Width-fcr.button.MinSize().Width, size.Height))
	topLeft = topLeft.Add(fyne.NewPos(size.Width-fcr.label.MinSize().Width-fcr.button.MinSize().Width, 0))
	fcr.button.Move(topLeft)
	fcr.button.Resize(fyne.NewSize(fcr.button.MinSize().Width, size.Height))
}

func (fcr fileChoiceRender) Destroy() {}

func (fcr fileChoiceRender) Refresh() {}

func (fc FileChoice) CreateRenderer() fyne.WidgetRenderer {
	label := widget.NewLabel(fc.labelText)
	entry := widget.NewEntryWithData(fc.path)
	entry.PlaceHolder = fc.placeHolder
	button := widget.NewButton(fc.buttonText, func() {
		cwd, _ := os.Getwd()
		defer os.Chdir(cwd)
		file, err := dialog.File().Load()
		if err != nil {
			file = ""
		}
		if file != "" {
			entry.SetText(file)
		}
	})
	return &fileChoiceRender{label: label, entry: entry, button: button}
}
