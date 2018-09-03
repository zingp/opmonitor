package main

import (
	"time"
	"fmt"
	"sync"
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

var waitGroup sync.WaitGroup

func main() {
	confFile := "./conf/app.cfg"
	err := initConfig(confFile)
	if err != nil {
		logs.Error("init config failed:%v", err)
		return
	}

	err = initLogs(appConf.LogFile)
	if err != nil {
		fmt.Printf("init log failed:%v\n", err)
		return
	}

	for k, v := range appConf.ProcMaP{
		waitGroup.Add(1)
		go checkProc(k, v.StartCmd, time.Duration(v.TimeInterval)* time.Second)
	}

	c := cron.New()
	c.AddFunc("1 * * * *", countKeyNum)
	c.Start()

	logs.Info("monitor start")
	waitGroup.Wait()
}
