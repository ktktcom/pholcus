package gui

import (
	. "github.com/ktktcom/pholcus/gui/model"
	"github.com/ktktcom/pholcus/runtime/cache"
	"github.com/ktktcom/pholcus/runtime/status"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

var (
	runStopBtn      *walk.PushButton
	pauseRecoverBtn *walk.PushButton
	setting         *walk.Composite
	mw              *walk.MainWindow
	runMode         *walk.GroupBox
	db              *walk.DataBinder
	ep              walk.ErrorPresenter
	mode            *walk.GroupBox
	host            *walk.Splitter
	spiderMenu      *SpiderMenu
)

// GUI输入
type Inputor struct {
	Spiders []*GUISpider
	*cache.AppConf
	BaseSleeptime     uint
	RandomSleepPeriod uint
}

var Input = &Inputor{
	// 默认值
	AppConf:           cache.Task,
	BaseSleeptime:     cache.Task.Pausetime[0],
	RandomSleepPeriod: cache.Task.Pausetime[1],
}

//****************************************GUI内容显示配置*******************************************\\

// 下拉菜单辅助结构体
type KV struct {
	Key    string
	Int    int
	Uint   uint
	String string
}

// 暂停时间选项及运行模式选项
var GuiOpt = struct {
	SleepTime           []*KV
	Mode                []*KV
	DeduplicationTarget []*KV
}{
	SleepTime: []*KV{
		{Key: "无暂停", Uint: 0},
		{Key: "0.1 秒", Uint: 100},
		{Key: "0.3 秒", Uint: 300},
		{Key: "0.5 秒", Uint: 500},
		{Key: "1 秒", Uint: 1000},
		{Key: "3 秒", Uint: 3000},
		{Key: "5 秒", Uint: 5000},
		{Key: "10 秒", Uint: 10000},
		{Key: "15 秒", Uint: 15000},
		{Key: "20 秒", Uint: 20000},
		{Key: "30 秒", Uint: 30000},
		{Key: "60 秒", Uint: 60000},
	},
	Mode: []*KV{
		{Key: "单机", Int: status.OFFLINE},
		{Key: "服务器", Int: status.SERVER},
		{Key: "客户端", Int: status.CLIENT},
	},
	DeduplicationTarget: []*KV{
		{Key: "去重样本位置: file", String: status.FILE},
		{Key: "去重样本位置: mgo", String: status.MGO},
	},
}

// 输出选项
var outputList []declarative.RadioButton
