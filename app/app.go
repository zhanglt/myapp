package app

import (
	"fmt"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
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

	myapp := &MyAPP{
		app: app,
		win: app.NewWindow("备份工具"),
	}
	myapp.setToolbar()
	myapp.setTheme()
	myapp.setMenu()
	return myapp
}

func (myapp *MyAPP) run() {
	myapp.win.ShowAndRun()
}
func (myapp *MyAPP) setTheme() {
	myapp.app.Settings().SetTheme(&mytheme.MyTheme{})
}

func (myapp *MyAPP) setToolbar() {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("first toolbar button")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.HomeIcon(), func() {

		}),
	)
	myapp.toolbar = toolbar
}

func (myapp *MyAPP) setMenu() {
	item1 := fyne.NewMenuItem("打开", nil)
	item1.IsQuit = true
	item2 := fyne.NewMenuItem("保存", nil)
	menu1 := fyne.NewMenu("文件", item1, item2)

	menu2 := fyne.NewMenu("编辑", item1, item2)
	menu := fyne.NewMainMenu(menu1, menu2)
	myapp.menu = menu

}
