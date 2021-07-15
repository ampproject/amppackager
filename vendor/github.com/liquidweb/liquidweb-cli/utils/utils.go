/*
Copyright © LiquidWeb

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/k0kubun/go-ansi"
)

func IpIsValid(ip string) bool {
	if parsedIp := net.ParseIP(ip); parsedIp == nil {
		return false
	}

	return true
}

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func FileExists(file string) bool {
	fileStat, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}

	return !fileStat.IsDir()
}

func PrintRed(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	if _, err := ansi.Print(red(msg)); err != nil {
		fmt.Printf("Error printing to console. Error was [%s] original message: [%s]\n", err, msg)
	}
}

func PrintTeal(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	if _, err := ansi.Print(teal(msg)); err != nil {
		fmt.Printf("Error printing to console. Error was [%s] original message: [%s]\n", err, msg)
	}
}

func PrintGreen(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	if _, err := ansi.Print(green(msg)); err != nil {
		fmt.Printf("Error printing to console. Error was [%s] original message: [%s]\n", err, msg)
	}
}

func PrintYellow(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	if _, err := ansi.Print(yellow(msg)); err != nil {
		fmt.Printf("Error printing to console. Error was [%s] original message: [%s]\n", err, msg)
	}
}

func PrintMagenta(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	if _, err := ansi.Print(magenta(msg)); err != nil {
		fmt.Printf("Error printing to console. Error was [%s] original message: [%s]\n", err, msg)
	}
}

func PrintPurple(m string, args ...interface{}) {
	msg := fmt.Sprintf(m, args...)
	if _, err := ansi.Print(purple(msg)); err != nil {
		fmt.Printf("Error printing to console. Error was [%s] original message: [%s]\n", err, msg)
	}
}

// private

var (
	teal    = colorize("\033[1;36m%s\033[0m")
	red     = colorize("\033[1;31m%s\033[0m")
	green   = colorize("\033[1;32m%s\033[0m")
	yellow  = colorize("\033[1;33m%s\033[0m")
	magenta = colorize("\033[1;35m%s\033[0m")
	purple  = colorize("\033[1;34m%s\033[0m")
)

func colorize(color string) func(...interface{}) string {
	colorized := func(args ...interface{}) string {
		return fmt.Sprintf(color,
			fmt.Sprint(args...))
	}

	return colorized
}
