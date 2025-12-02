package naloga5

import "fmt"

type Student struct {
    Ime     string
    Priimek string
    Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpis string, ocena int, minOcena int, maxOcena int) {
    if ocena < minOcena || ocena > maxOcena {
        fmt.Printf("Ocena mora biti med %d in %d.\n", minOcena, maxOcena)
        return
    }

    s, ok := studenti[vpis]
    if !ok {
        fmt.Println("Študent ne obstaja.")
        return
    }

    s.Ocene = append(s.Ocene, ocena)
    studenti[vpis] = s
}

func povprecje(s Student) float64 {
    if len(s.Ocene) == 0 {
        return 0
    }

    vsota := 0
    for _, o := range s.Ocene {
        vsota += o
    }
    return float64(vsota) / float64(len(s.Ocene))
}

func IzpisVsehOcen(studenti map[string]Student) {
    fmt.Println("Redovalnica:")
    for vpis, s := range studenti {
        fmt.Printf("%s - %s %s: %v\n", vpis, s.Ime, s.Priimek, s.Ocene)
    }
}

func IzpisiKoncniUspeh(studenti map[string]Student, minStOcen int) {
    for _, s := range studenti {
        p := povprecje(s)

        fmt.Printf("%s %s -> povprečje %.1f: ", s.Ime, s.Priimek, p)

        if len(s.Ocene) < minStOcen {
            fmt.Println("Premalo ocen.")
        } else if p >= 9 {
            fmt.Println("Odličen študent")
        } else if p >= 6 {
            fmt.Println("Povprečen študent")
        } else {
            fmt.Println("Neuspešen študent")
        }
    }
}
