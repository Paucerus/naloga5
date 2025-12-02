// Package redovalnica ponuja funkcije za upravljanje ocen študentov.
//
// Paket omogoča dodajanje ocen študentom, prikaz vseh ocen in izračun končnega uspeha.
//
// Osnovni podatkovni tip je Student, ki vsebuje ime, priimek in seznam ocen.
//
// # Izvožene funkcije
//
// DodajOceno doda oceno študentu in preveri, ali je v dovoljenih mejah.
//
// IzpisVsehOcen prikaže seznam vseh študentov z njihovimi ocenami.
//
// IzpisiKoncniUspeh izračuna povprečje ocen in kategorizira študente
// (odličen, povprečen, neuspešen).
//
// # Primer uporabe
//
//	studenti := map[string]redovalnica.Student{
//	    "63210001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8}},
//	}
//	redovalnica.DodajOceno(studenti, "63210001", 10, 0, 10)
//	redovalnica.IzpisVsehOcen(studenti)
//	redovalnica.IzpisiKoncniUspeh(studenti, 6)
package redovalnica
