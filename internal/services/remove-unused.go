package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/marvinarlt/goshort/internal/consts"
)

func RemoveUnused() {
	intervalRaw := os.Getenv("APP_CLEANUP_INTERVAL")
	interval, err := strconv.Atoi(intervalRaw)

	if nil != err {
		log.Fatalln(consts.PREFIX, "APP_CLEANUP_INTERVAL needs to be an integer")
	}

	sql := fmt.Sprintf("DELETE FROM links WHERE latest_use < (NOW() - INTERVAL %s Day) OR latest_use IS NULL", os.Getenv("APP_CLEANUP_OLDER_THAN"))

	ticker := time.NewTicker(time.Duration(interval) * time.Hour)
	channel := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				DB.Exec(sql)
				fmt.Println(consts.PREFIX, "Clean up")
			case <-channel:
				ticker.Stop()
				return
			}
		}
	}()
}
