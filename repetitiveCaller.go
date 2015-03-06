package main

import (
	"strconv"
	"time"

	"github.com/scalarm/scalarm_workers_manager/logger"
)

func RepetitiveCaller(f func() (interface{}, error), intervals []int, functionName string) (out interface{}, err error) {
	if intervals == nil {
		intervals = []int{15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15}
	}

	intervals = append(intervals, -1)

	for _, duration := range intervals {
		out, err = f()
		if err == nil || duration == -1 {
			return
		}
		logger.Info("RepetitiveCaller : call " + functionName + " failed, err: \n" + err.Error() + "\nReattempt in " + strconv.Itoa(duration) + "s")
		time.Sleep(time.Second * time.Duration(duration))
	}
	return
}
