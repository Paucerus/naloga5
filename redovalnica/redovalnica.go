package redovalnica

import "fmt"

// Student predstavlja podatke o študentu.
// Vsebuje ime, priimek in seznam ocen.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno študentu z dano vpisno številko.
// Funkcija preveri, ali je ocena v mejah [minOcena, maxOcena].
// Če študent ne obstaja ali je ocena izven mej, izpiše opozorilo.
//
// Parametri:
//   - studenti: mapa študentov, kjer je ključ vpisna številka
//   - vpis: vpisna številka študenta
//   - ocena: ocena, ki jo dodajamo
//   - minOcena: najmanjša dovoljena ocena
//   - maxOcena: največja dovoljena ocena
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

// povprecje je interna funkcija, ki izračuna povprečno oceno študenta.
// Vrne 0, če študent nima nobene ocene.
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

// IzpisVsehOcen izpiše seznam vseh študentov z njihovimi ocenami.
// Za vsakega študenta izpiše vpisno številko, ime, priimek in seznam ocen.
//
// Parametri:
//   - studenti: mapa študentov, kjer je ključ vpisna številka
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("Redovalnica:")
	for vpis, s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpis, s.Ime, s.Priimek, s.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh vseh študentov.
// Za vsakega študenta izračuna povprečje in ga razvršča v kategorije:
//   - "Odličen študent" če je povprečje >= 9
//   - "Povprečen študent" če je povprečje >= 6
//   - "Neuspešen študent" če je povprečje < 6
//   - "Premalo ocen." če študent nima dovolj ocen
//
// Parametri:
//   - studenti: mapa študentov, kjer je ključ vpisna številka
//   - minStOcen: minimalno število ocen potrebnih za ocenjevanje
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
