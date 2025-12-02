package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Paucerus/naloga5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "naloga5",
		Usage: "Demonstracija redovalnice (preneseno iz naloga1)",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stOcen",
				Value:   6,
				Usage:   "Najmanjše število ocen potrebnih za pozitivno oceno",
				Aliases: []string{"n"},
			},
			&cli.IntFlag{
				Name:    "minOcena",
				Value:   0,
				Usage:   "Najmanjša možna ocena",
				Aliases: []string{"min"},
			},
			&cli.IntFlag{
				Name:    "maxOcena",
				Value:   10,
				Usage:   "Največja možna ocena",
				Aliases: []string{"max"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Preberemo vrednosti iz stikal
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			studenti := map[string]redovalnica.Student{
				"63210001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8, 9, 10, 8}},
				"63210002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8, 7, 6}},
				"63210003": {Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3, 5, 4, 5}},
			}

			// Dodajanje ocen (posnemam naloga1 izvajanje)
			redovalnica.DodajOceno(studenti, "63210001", 10, minOcena, maxOcena)
			redovalnica.DodajOceno(studenti, "63210004", 7, minOcena, maxOcena)
			redovalnica.DodajOceno(studenti, "63210002", 11, minOcena, maxOcena)

			// Izpis redovalnice in končnih uspehov
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
