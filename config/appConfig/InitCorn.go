package appconfig

import (
	"sync"
	"time"

	"github.com/lenna-ai/bni-iproc/app/kernel/console"
	"github.com/robfig/cron/v3"
)

func initCornJob() {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	var wg sync.WaitGroup

	go scheduler.AddFunc("* * * * *", func() {
		wg.Add(1)
		console.RemoveFileStorage("./storage/logs", 7)
	})

	wg.Wait()
	go scheduler.Start()

	scheduler.Stop()

}
