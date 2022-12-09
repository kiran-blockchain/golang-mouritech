package main

import "fmt"

type Config struct {
	Env   string
	Proxy ProxyInfo
}

type ProxyInfo struct {
	Address string
	Port    string
}

func (conf Config) ConfInfo() string {
	fmt.Println("Env: ", conf.Env)
	fmt.Println("Proxy: ", conf.Proxy)
	return "----------------------"
}

func main() {
	c := &Config{
		Env: "DEBUG:TRUE",
		
		Proxy: ProxyInfo{
			Address: "addr",
			Port:    "port",
		},
	}
    c.Proxy.Port="4000"

	fmt.Println(c.ConfInfo())
	fmt.Println(c.Proxy)
}
