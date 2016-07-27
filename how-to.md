# Golang

## Commandes
    go run main.go // Exécute le script main.go
    go build main.go // Crée un executable
    go run --work main.go // Exécute le script et montre où se trouve l'executable temporaire
    godoc fmt Prinln // Affiche les informations sur la commande Println du paquet fmt

## Documentation
    godoc -http:=6060 // Documentation offline disponible à l'adresse http://localhost:6060

## Installation
    sudo tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin // Dans $HOME/.profile
    . ~/.profile // Pour MAJ du chemin vers go OU
    source ~/.profile // MAJ du chemin (Variante)

Si une précédente installation de Golang a été faites, il faut la supprimer :

    sudo rm -Rf /usr/local/go

## Imports
    // Plusieurs librairies
    import (
      "fmt"
      "os"
      )

    // Une seule libraire
    import "fmt"

## variables
  var power int = 9000
  //OU
  power := 9000 // Le programme détermine automatiquement le type
  //OU
  var power int // Ici power sera à zéro

## Fonctions
     func log(message string){

     }

     func add(a int, b int) int {

     }

     // Ou (Version courte)
     func add(a, b int) int {}

     func power(name string) (int, bool){

     }
     // Utilisation #1 :
     value, exists := power("goku")
     // Utilisation #2 :
      _ , exists := power("goku")  // Ici, on n'est pas intéressé de récupérer la valeur (int)


## Structures
    // Déclaration
    type Saiyan struct {
      Name string
      Power int
      Father * Saiyan // Possible d'avoir un une structure à l'intérieur d'une autre structure
    }

    // Utilisation (Plusieur façon)
    goku := Saiyan{}
    goku := Saiyan {
      Name := "Goku",
      Power := 9000,
    }

    goku := Saiyan{Name: "Goku"}
    goku.Power = 9000
    goku := Saiyan("Goky", 9000)

    // Structure et fonction
    func (s * Saiyan) Super() {

    }

    Utilisation de la fonction
    goku.Super()

    // Pointeurs sur objet
    goku := &Saiyan{}
    // OU
    goku := new(Saiyan) // Est égal à ce qu'il y a au dessus.

    // Composition
    type Person struct() {
      Name string
    }

    fun (p * Person) Introduce() {
      fmt.Println("Hi, I'm %s\", p.Name)
    }

    type Saiyan struct {
      * Person
      Power int
    }

    // And to use it
    goku := &Saiyan {
      Person: &Person{"Goku"},
      Power: 9001,
    }

    goku.Introduce()

## Tableaux et Slice

### Tableaux
  // Déclaration : En go, les tableaux ont des  taille fixes
    var scores [10]int // Tableau vide
    score := [4]int{50, 105, 512, 98} // Tableau avec des valeurs

  // Affectation
    scores[0] = 50

  // Parcours tableau
    for index, valeur := range scores {
    }

### Slices
  // Il s'agit de portion de tableau, plus souple que les tableaux
  // Déclaration #1
    scores := []int{1,4,293,4,9}

  // # Déclaration #2
    scores := make([]int, 10) // make(), alloue la mémoire (comme new()) ET initialise le slice. Ici, on crée un slice d'une longueur de 10 et d'une capacité de 10. La longueur est la taille du slice et la capacité la taille du tableau sous-jacent.

  // # Déclaration #3
  // Il est possible de spécifier les deux :
    scores := make([]int, 0, 10) // Ici, le slice à une taille de 0 et le tableau une capacité de 10.

  // # Déclaration #4
    var scores []int

  // Ajouter un élément à un slice
    scores = append(scores, 5) // https://golang.org/pkg/builtin/#append
  // append() augmente la capacité du tableau au besoin

  // Changer la taille du slice
    scores = scores[0:8] // ATTENTION : Ne peut pas dépasser la capacité du tableau !!

  // cap() : Retourne la capacité du tableau
    c := cap(scores) // https://golang.org/pkg/builtin/#cap
## Astuces
 Un projet doit toujours être composé comme ceci :
    /code/src/mon_projet

 Il y a TOUJOURS un paquet "main" et un fonction "main". C'est par là que va commencer le programme.

 Les variables sont toujours passées par valeurs et non par référence. Donc une variable passée dans une fonction sera une copie de la variable appelante et ne sera pas modifié. Si on veut la modifier, il faut utiliser les pointeurs.
