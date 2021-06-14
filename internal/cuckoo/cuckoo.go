package cuckoo

import (
	"CuckooGo/internal/config"
	"CuckooGo/internal/filter"
	"CuckooGo/internal/rpc"
	"time"
)

func Run() {
	conf := config.Read()
	filter := filter.NewFilter(conf.Capacity, conf.File)
	rpcSer := rpc.RpcServer(conf.RpcPort, filter)
	ticker := time.NewTicker(time.Minute * 2)
	go func() {
		previousCount := filter.Count()
		for range ticker.C {
			if previousCount != filter.Count() {
				previousCount = filter.Count()
				filter.SaveFile()
			}
		}
	}()
	rpcSer.Listen()
}
