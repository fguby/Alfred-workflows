package Workflow

import (
	"encoding/json"
	"fmt"
)

type Icon struct {
	Path string `json:"path"`
}

//单个alfred结构体
type AlfredWorkflow struct {
	Valid    bool   `json:"valid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Icon     Icon   `json:"icon"`
	Arg      string `json:"arg"`
}

//一组alfred结构体
type Alfreds struct {
	Arrs []AlfredWorkflow `json:"items"`
}

//创建alfred对象
func New() *Alfreds {
	return &Alfreds{}
}

//添加一组新通知
func (a *Alfreds) Add(valid bool, title, subtitle, icon, arg string) {
	alfred := AlfredWorkflow{
		Valid:    valid,
		Title:    title,
		Subtitle: subtitle,
		Icon: Icon{
			Path: icon,
		},
		Arg: arg,
	}
	a.Arrs = append(a.Arrs, alfred)

}

//回送通知
func (a *Alfreds) SendFeedback() {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

//旧版本xml
// func (a *Alfreds) SendFeedback() {
// 	var build strings.Builder
// 	build.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")
// 	build.WriteString("<items>\n")
// 	for _, value := range a.Arrs {
// 		build.WriteString("<item valid=\"")
// 		build.WriteString(value.Valid)
// 		build.WriteString("\">\n")
// 		build.WriteString("<title>")
// 		build.WriteString(value.Title)
// 		build.WriteString("</title>\n")
// 		build.WriteString("<subtitle>")
// 		build.WriteString(value.Subtitle)
// 		build.WriteString("</subtitle>\n")
// 		build.WriteString("<icon>")
// 		build.WriteString(value.icon)
// 		build.WriteString("</icon>\n")
// 		build.WriteString("</item>\n")
// 	}
// 	build.WriteString("</items>\n")
// 	fmt.Print(build.String())
// }
