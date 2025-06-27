package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"math"
	"strconv"
	"strings"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "first-app-command-line"
	app.Usage = "Get IPs address of a host"
	app.Version = "0.0.1"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "convert",
			Value: "todec", // Value is the default value for the flag.
		},
		cli.StringFlag{
			Name:  "system",
			Value: "bin", // Value is the default value for the flag.
		},
		cli.StringFlag{
			Name:  "code",
			Value: "300", // Value is the default value for the flag.
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "converter",
			Usage:  "number systems converter",
			Flags:  flags,
			Action: converter,
		},
	}

	return app
}

func converter(c *cli.Context) error {
	system := c.String("system")
	convert := c.String("convert")

	if convert == "todec" {
		decimal := 0
		switch system {
		case "bin":
			code := c.Int("code")
			decimal = binaryToDecimal(code)
		case "oct":
			code := c.Int("code")
			decimal = octalToDecimal(code)
		case "dec":
			code := c.Int("code")
			decimal = decimalToDecimal(code)
		case "hex":
			str := c.String("code")
			decimal = int(hexadecimalToDecimal(str))
		}

		fmt.Println(decimal)
	} else if convert == "decto" {
		conveted := ""
		switch system {
		case "bin":
			code := c.Int("code")
			conveted = DecimalToBinary(code)
		case "oct":
			code := c.Int("code")
			conveted = DecimalToOctal(code)
		case "hex":
			str := c.Int("code")
			conveted = DecimalToHexadecimal(str)
		}

		fmt.Println(conveted)

	}

	return nil
}

func binaryToDecimal(code int) int {
	decimal := 0
	i := 0
	rem := 0

	for code > 0 {
		rem = code % 10
		decimal += rem * int(math.Pow(2, float64(i)))
		code /= 10
		i++
	}

	return decimal
}

func octalToDecimal(code int) int {
	decimal := 0
	i := 0
	rem := 0

	for code > 0 {
		rem = code % 10
		decimal += rem * int(math.Pow(8, float64(i)))
		code /= 10
		i++
	}

	return decimal

}

func decimalToDecimal(code int) int {
	return code
}

func hexadecimalToDecimal(hexaString string) uint64 {
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	output, err := strconv.ParseUint(numberStr, 16, 64)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return output
}

func DecimalToBinary(code int) string {
	binary := ""
	i := 0
	rem := 0
	for code > 0 {
		rem = code % 2
		binary = strconv.Itoa(rem) + binary
		code /= 2
		i++
	}
	return binary
}

func DecimalToOctal(code int) string {
	octal := ""
	i := 0
	rem := 0
	for code > 0 {
		rem = code % 8
		octal = strconv.Itoa(rem) + octal
		code /= 8
		i++
	}
	return octal
}

func DecimalToHexadecimal(code int) string {
	hexadecimal := ""
	i := 0
	rem := 0
	for code > 0 {
		rem = code % 16
		hexadecimal = strconv.Itoa(rem) + hexadecimal
		code /= 16
		i++
	}
	return hexadecimal
}
