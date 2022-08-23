package main
  
import (
    "fmt"
)

type Astro struct {
    
    name string
    age int
    mission string
    isneeded bool
}

type nasaMission struct {
    
    people []Astro
    number int
    message string
}

func main() {

        astr0 := Astro{"Michael", 37, "comp", false}
        astr1 := Astro{"Thomas", 23, "engineer", true}
        astr2 := Astro{"Sidsel", 29, "data", true}

        fmt.Println(astr0)
        fmt.Println(astr1)
        fmt.Println(astr2)

        p := []Astro{astr0, astr1, astr2}

        fmt.Println(p)

        fmt.Println(p[1].mission)

        s := nasaMission{p, 3, "failure"}

        fmt.Println(s)

        fmt.Println("%+v",s)
}
