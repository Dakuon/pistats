package main

import (
	"os/exec"
)

type Data struct {
	IP string
	CPU string
	Mem string
	SD string
	SSD string
	Temp string
}


func run(cmd string) string {
	run := exec.Command("bash", "-c", cmd)
	stdout, err := run.Output()
	if err != nil {
	  panic(err)
	}
	return string(stdout)
}

func getData() (*Data) {
	var data = &Data{
		IP: run("hostname -I | cut -d' ' -f1 | head --bytes -1"),
		CPU: run("top -bn1 | grep load | awk '{printf \"%.2f\", $(NF-2)}'"),
		Mem: run("free -m | awk 'NR==2{printf \"%.2f%%\", $3*100/$2 }'"),
		SD: run("df -h | awk '$NF==\"/\"{printf \"%s\", $5}'"),
		Temp: run("vcgencmd measure_temp | cut -d '=' -f 2 | head --bytes -1"),
	}

	return data
}