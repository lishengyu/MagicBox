// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	. "tools/global"
	_ "tools/proc_string"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func main() {
	logFile, err := os.OpenFile("operate.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Printf("Init Log ......\n")

	log.Println("Tools Starting ....")
	mw := &MyMainWindow{model: NewFuncList()}

	ButtomClicked := func(text string) string {
		return GFuncList.Handlers[mw.lb.CurrentIndex()](text)
	}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Magic Box",
		MinSize:  Size{600, 800},
		Size:     Size{600, 800},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{
						AssignTo:              &mw.lb,
						Model:                 mw.model,
						MinSize:               Size{30, 100},
						OnCurrentIndexChanged: mw.lb_CurrentIndexChanged,
						OnItemActivated:       mw.lb_ItemActivated,
					},
					VSplitter{
						Children: []Widget{
							TextEdit{AssignTo: &mw.te},
							TextEdit{AssignTo: &mw.tt, ReadOnly: true},
							PushButton{
								Text: "convert",
								OnClicked: func() {
									text := mw.te.Text()
									log.Printf("PushButton:%s\n", text)
									mw.tt.SetText(ButtomClicked(text))
								},
							},
						},
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

type MyMainWindow struct {
	*walk.MainWindow
	model *FuncModel
	lb    *walk.ListBox
	te    *walk.TextEdit
	tt    *walk.TextEdit
}

func (mw *MyMainWindow) lb_CurrentIndexChanged() {
	i := mw.lb.CurrentIndex()
	if i < 0 {
		//点击空白区域
		return
	}
	item := mw.model.items[i]

	mw.te.SetText(item)

	fmt.Println("CurrentIndex: ", i)
	fmt.Println("CurrentFuncName: ", item)
}

func (mw *MyMainWindow) lb_ItemActivated() {
	i := mw.lb.CurrentIndex()
	if i < 0 {
		//点击空白区域
		return
	}
	value := mw.model.items[i]

	walk.MsgBox(mw, "Value", value, walk.MsgBoxIconInformation)
}

type FuncModel struct {
	walk.ListModelBase
	items []string
}

func NewFuncList() *FuncModel {
	m := &FuncModel{items: make([]string, GFuncList.Count)}

	log.Printf("register func Num:[%d]\n", GFuncList.Count)
	for i := 0; i < GFuncList.Count; i++ {
		m.items[i] = GFuncList.FuncList[i]
		log.Printf("register func[%d]:%s\n", i, GFuncList.FuncList[i])
	}

	return m
}

func (m *FuncModel) ItemCount() int {
	return len(m.items)
}

func (m *FuncModel) Value(index int) interface{} {
	return m.items[index]
}
