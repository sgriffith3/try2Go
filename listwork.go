package main

import(
  "fmt"
  "strconv"
)

func main() {
  //use :8888/couses/stu?=python.60.79.22 to set courseName, hStudent, lStudent, numLabs
    hStudent := int(79)
    lStudent := int(60)
    numLabs := int(22)
    labList := []string{}
    for x := 0; x <= (numLabs); x ++ {
        conv := strconv.Itoa(x)
        labList = append(labList, conv)
    }
    initTable := `<tr><th>Student Name</th><th>Student ID</th>`
    for y := 0; y <= numLabs; y ++ {
      p := fmt.Sprintf(`<th>%d</td>`, y)
      initTable += p
    }
    initTable += `</tr>`
    fmt.Printf("This printed: %s", initTable)

}





//   // student := [name, id, [lab# [start, end]], [lab# [start, end]]
//   // for i :=
// }
