package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type ArrayFlag []string

func (a *ArrayFlag) String() string {
	return fmt.Sprintf("%#v", a)
}

func (a *ArrayFlag) Set(value string) error {
	*a = append(*a, value)

	return nil
}

func main() {
	var name string
	var dir string
	var values ArrayFlag

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "GoTemplet v0.0.1\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(w, "-------\n")
	}

	flag.StringVar(&name, "name", "test", "Nombre de la plantilla a generar")
	flag.StringVar(&dir, "dir", "", "Ubicación de directorio extra donde se buscará la plantilla")
	flag.Var(&values, "V", "Indica una pareja `Clave:Valor` para utilizarla en la plantilla")

	flag.Parse()

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	log.Println("-------------------------")
}
