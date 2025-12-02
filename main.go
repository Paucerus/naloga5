package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli/v3"
    "github.com/Paucerus/naloga5/naloga5"
)

func main() {
    app := &cli.App{
        Name:  "redovalnica",
        Usage: "Upravljanje ocen",
        Flags: []cli.Flag{
            &cli.IntFlag{Name: "stOcen", Value: 6},
            &cli.IntFlag{Name: "minOcena", Value: 1},
            &cli.IntFlag{Name: "maxOcena", Value: 10},
        },
        Action: func(ctx *cli.Context) error {
            minOcena := ctx.Int("minOcena")
            maxOcena := ctx.Int("maxOcena")
            stOcen := ctx.Int("stOcen")

            studenti := map[string]redovalnica.Student{
                "63210001": {"Ana", "Novak", []int{10, 9, 8, 9, 10, 8}},
                "63210002": {"Boris", "Kralj", []int{6, 7, 5, 8, 7, 6}},
            }

            redovalnica.DodajOceno(studenti, "63210001", 10, minOcena, maxOcena)
            redovalnica.DodajOceno(studenti, "63210002", 7, minOcena, maxOcena)

            redovalnica.IzpisVsehOcen(studenti)
            fmt.Println()
            redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

            return nil
        },
    }

    app.Run(os.Args)
}
