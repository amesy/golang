package main

import (
	"fmt"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
}

func(p *Ban) visit(ip string) bool {
	if _, ok := p.visitIPs[ip]; ok {
		return true
	}
	p.visitIPs[ip] = time.Now()
	return false
}

func main() {
	success := 0
	ban := &Ban{visitIPs:make(map[string]time.Time)}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}()
		}
	}
	fmt.Println("success: ", success)
}
