package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fpabl0/sparky-go/swid"
	mytheme "github.com/zhanglt/myapp/theme"
)

type MyAPP struct {
	app     fyne.App
	win     fyne.Window
	menu    *fyne.MainMenu
	toolbar *widget.Toolbar
}

func NewAPP() *MyAPP {
	app := app.NewWithID("备份主页")
	win := app.NewWindow("备份工具")
	win.Resize(fyne.NewSize(400, 300))
	myapp := &MyAPP{
		app: app,
		win: win,
	}
	myapp.setToolbar()
	myapp.setTheme()
	myapp.setMenu()

	return myapp
}

func (myapp *MyAPP) Run() {
	myapp.win.ShowAndRun()
}
func (myapp *MyAPP) setTheme() {
	myapp.app.Settings().SetTheme(&mytheme.MyTheme{})
}
func (myapp *MyAPP) setToolbar() {
	var r fyne.Resource
	r, _ = fyne.LoadResourceFromPath("img/add1.png")
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(r, func() {
			fmt.Println("first toolbar button")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(r, func() {

		}),
	)
	button := widget.NewButtonWithIcon("测试图标", r, func() {})
	button.IconPlacement = 0
	button.IconPosition = layout.NewVBoxLayout()

	//file := NewFileChoice("文件选择", "bbb", "....")
	a := []string{"c", "b"}
	file := swid.NewSelectField(a, func(string) {})
	content := container.NewBorder(toolbar, nil, nil, nil, file, file)
	myapp.win.SetContent(content)
	myapp.toolbar = toolbar
}
func (myapp *MyAPP) setMenu() {
	item1 := fyne.NewMenuItem("打开", nil)
	item1.IsQuit = true
	item2 := fyne.NewMenuItem("保存", nil)
	menu1 := fyne.NewMenu("文件", item1, item2)
	menu2 := fyne.NewMenu("编辑", item1, item2)
	menu := fyne.NewMainMenu(menu1, menu2)
	myapp.win.SetMainMenu(menu)
	myapp.menu = menu

}
