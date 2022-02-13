package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zhanglt/myapp/myapp"
	"github.com/zhanglt/myapp/theme"
)

func main() {
	app := myapp.NewAPP()
	app.New().Run()

}
func main1() {
	myApp := app.New()
	wm := myApp.NewWindow("备份工具")
	wm.SetMainMenu(menu2())
	content := container.NewBorder(toolbar(), nil, nil, nil, widget.NewLabel("测试中文"))
	wm.SetContent(content)
	// lets create a label
	label1 := widget.NewLabel("...")
	// newselectentry widget
	// []string{} it take only slice of option
	//& you can your options also in run time. Not hardcoded
	select_entry := widget.NewSelectEntry([]string{"本地硬盘", "SFTP", "Minio", "S3"})

	// what to do with the selected entry ?
	// here is what we are going to define
	select_entry.OnSubmitted = func(s string) {
		fmt.Printf("my city %s is awesome", s)
		// update label with our values
		label1.Text = s
		label1.Refresh()
	}
	// container .. we have more than one widgets
	c := container.NewVBox(
		content,
		label1,
		select_entry,
	)
	wm.SetContent(c)
	wm.Show()
	wm.Resize(fyne.NewSize(400, 300))
	myApp.Settings().SetTheme(theme.MyTheme{})
	myApp.Run()

}

func toolbar() *widget.Toolbar {
	var r fyne.Resource
	r, _ = fyne.LoadResourceFromPath("img/add1.png")
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(r, func() {
			fmt.Println("create")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(r, func() {
			fmt.Println("create")
		}),
	)

	return toolbar
}

func menu() *fyne.MainMenu {
	item1 := fyne.NewMenuItem("打开", nil)
	item1.IsQuit = true
	item2 := fyne.NewMenuItem("保存", nil)
	menu1 := fyne.NewMenu("文件", item1, item2)

	menu2 := fyne.NewMenu("编辑", item1, item2)
	menu := fyne.NewMainMenu(menu1, menu2)

	return menu

}

func menu2() *fyne.MainMenu {
	// new menu items
	//first parameter is label, 2nd is function
	item1 := fyne.NewMenuItem("edit", nil)
	item2 := fyne.NewMenuItem("details", nil)
	item3 := fyne.NewMenuItem("home", nil)
	item4 := fyne.NewMenuItem("run", nil)
	// child menu
	item1.ChildMenu = fyne.NewMenu(
		"",                            // leave label blank
		fyne.NewMenuItem("copy", nil), // child menu items
		fyne.NewMenuItem("cut", nil),
		fyne.NewMenuItem("paste", nil),
	)
	// create child menu for 2nd item
	item2.ChildMenu = fyne.NewMenu(
		"",                             // leave label blank
		fyne.NewMenuItem("books", nil), // child menu items
		fyne.NewMenuItem("magzine", nil),
		fyne.NewMenuItem("notebook", nil),
	)
	// create child menu for third item
	item3.ChildMenu = fyne.NewMenu(
		"",                              // leave label blank
		fyne.NewMenuItem("school", nil), // child menu items
		fyne.NewMenuItem("college", nil),
		fyne.NewMenuItem("university", nil),
	)
	NewMenu1 := fyne.NewMenu("File", item1, item2, item3, item4)
	NewMenu2 := fyne.NewMenu("Help", item1, item2, item3, item4)
	// main menu
	menu := fyne.NewMainMenu(NewMenu1, NewMenu2)
	// setup menu
	return menu
}
