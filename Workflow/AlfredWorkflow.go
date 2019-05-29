package Workflow

import (
	"fmt"
	"strings"
)

//Alfred通知结构体
type AlfredWorkflow struct {
	Valid    string
	Title    string
	Subtitle string
	icon     string
}

//存放切片
type Alfreds struct {
	Arrs []AlfredWorkflow
}

//构造方法
func New() *Alfreds {
	return &Alfreds{}
}

//添加一组回显信息
func (a *Alfreds) Add(valid, title, subtitle, icon string) {
	alfred := AlfredWorkflow{
		Valid:    valid,
		Title:    title,
		Subtitle: subtitle,
		icon:     icon,
	}
	a.Arrs = append(a.Arrs, alfred)

}

//展示所有的回显信息
func (a *Alfreds) SendFeedback() {
	var build strings.Builder
	build.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")
	build.WriteString("<items>\n")
	for _, value := range a.Arrs {
		build.WriteString("<item valid=\"")
		build.WriteString(value.Valid)
		build.WriteString("\">\n")
		build.WriteString("<title>")
		build.WriteString(value.Title)
		build.WriteString("<title>\n")
		build.WriteString("<subtitle>")
		build.WriteString(value.Subtitle)
		build.WriteString("</subtitle>\n")
		build.WriteString("<icon>")
		build.WriteString(value.icon)
		build.WriteString("</icon>\n")
		build.WriteString("</item>\n")
	}
	build.WriteString("</items>\n")
	fmt.Print(build.String())
}
