package main

import (
	"flag"
	"fmt"
	"github.com/ryo-endo/go-trans/translator"
	"os"
	"strings"
)

func main() {
	if err := Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Run(args []string) error {
	from := flag.String("from", "en", `Set the language code. "en" "ja" "vn"`)
	to := flag.String("to", "ja", `Set the language code. "en" "ja" "vn"`)
	flag.Parse()

	if len(args) <= 1 {
		return fmt.Errorf(`usage: trans "Hello world!"`)
	}

	key := os.Getenv("TRANS_API_KEY")
	if key == "" {
		return fmt.Errorf("You may need to set TRANS_API_KEY.\n# export TRANS_API_KEY=your-api-key")
	}

	transtalor := translator.NewAzure(key)

	input := strings.Join(flag.Args(), " ")
	out, err := transtalor.Trans(input, *from, *to)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}
