package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/mynet1314/nlan-admin/controllers"
	"github.com/mynet1314/nlan-admin/jobs"
	utility "github.com/mynet1314/nlan-admin/utils"
	"github.com/mynet1314/nlan/utils"
	"github.com/robfig/cron"
)

var (
	confPath    = flag.String("c", "./conf/config.toml", "config file")
	versionFlag = flag.Bool("v", false, "version")
	version     = "unknown"
)

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		return
	}

	utility.InitConf(*confPath)
	db := utils.InitDB(utility.DB_Driver, utility.DB_Connect)
	// For normal mode
	if len(os.Args) == 1 {
		go initRouter(db)
		go initJob(db)
		select {}
	} else if os.Args[1] == "recover" {
		jobs.RecoverTask(db)
	}
}

func initRouter(db *xorm.Engine) {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*.html")
	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r, db)
}

func initJob(db *xorm.Engine) {
	jobs.SetDB(db)
	c := cron.New()
	cj := &jobs.CronJob{}

	c.AddFunc("0 */5 * * * *", cj.InstantStats)
	c.AddFunc("0 0 0 * * *", cj.DailyStats)
	c.AddFunc("0 0 0 1 * *", cj.MonthlyStats)
	c.Start()
	select {}
}
