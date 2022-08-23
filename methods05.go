package main

import (
    "fmt"
)

type Virtmach struct {

    ip string
    diskgb int
    ram int
    compsnno string
}


func (v *Virtmach)diskgbset(gb int) {
    v.diskgb = v.diskgb + gb
}

func (v *Virtmach)compsnnoset(compsnno string){
    v.compsnno = compsnno
}




func (v Virtmach)display(){
    fmt.Println("Disk GB:", v.diskgb)
    fmt.Println("Comp Serial No.", v.compsnno)
    fmt.Println("IP Address:", v.ip)
    fmt.Println("Ram:", v.ram)
}

func main() {
    vm1 := Virtmach{"10.0.0.12", 48, 8,"sn:Tdkjosjf89"}

    vm.display()

    vm1.diskgbset(5)

    vm1.compsnnoset("sn:127898hf")
    
    vm1.display()
}
