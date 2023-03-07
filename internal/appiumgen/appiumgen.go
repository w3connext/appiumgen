package appiumgen

import (
	"fmt"
	"time"

	"github.com/w3connext/appiumgen/pkg/config"
)

func Run(conf config.Config) {
	fGen := NewFlutterGenerator()
	b := NewBuilder(fGen)

	startTime := time.Now()
	fmt.Println("[INFO] Generating build script...")

	err := b.Build(conf)

	stopTime := time.Now()
	timeUsage := stopTime.Sub(startTime).Milliseconds()
	if err != nil {
		fmt.Println("[ERROR] Failure after", timeUsage, "ms", err)
		return
	} else {
		fmt.Println("[INFO] Succeeded after", timeUsage, "ms")
	}

}
