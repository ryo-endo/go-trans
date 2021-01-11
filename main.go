package main

import (
	"bufio"
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
	flg := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	from := flg.String("from", "en", `Set the language code. "en" "ja" "vn"`)
	to := flg.String("to", "ja", `Set the language code. "en" "ja" "vn"`)
	interactive := flg.Bool("i", false, `interactive mode.`)
	flag.Parse()

	if len(args) <= 1 {
		return fmt.Errorf(`usage: trans "Hello world!"`)
	}

	key := os.Getenv("TRANS_API_KEY")
	if key == "" {
		return fmt.Errorf("You may need to set TRANS_API_KEY.\n# export TRANS_API_KEY=your-api-key")
	}

	trans := translator.NewAzure(key)

	if *interactive {
		err := interactiveTranslate(trans, *from, *to)
		if err != nil {
			return err
		}
	} else {
		err := translate(trans, *from, *to)
		if err != nil {
			return err
		}
	}

	return nil
}

func translate(translator translator.Translator, from string, to string) error {

	input := strings.Join(flag.Args(), " ")
	out, err := translator.Trans(input, from, to)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func interactiveTranslate(translator translator.Translator, from string, to string) error {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		out, err := translator.Trans(scanner.Text(), from, to)
		if err != nil {
			return err
		}

		fmt.Println(out)
	}

	//return nil
}
