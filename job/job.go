package job

import (
	"github.com/robfig/cron"
	"log"
)

func Setup() {
	go func() {
		c := cron.New()
		spec := "*/20 * * * * ?"
		c.AddFunc(spec, func() {

		})

		c.Start()
		log.Println("+++ cron job stared +++")
		defer c.Stop()
	}()
}
