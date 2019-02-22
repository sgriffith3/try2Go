package main

import (
	"html/template"
	"os"
	"log"
	"net/http"
	"fmt"
	"time"
	"path"
	// "net/url"
	"strconv"
	// "strings"
)


func errorHandler(status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		if status == http.StatusNotFound {
			r.URL.Path = "404.html"
			Template().ServeHTTP(w, r)
		}
	})
}


func iterTable(SName string, ID int, LLabs int)(studRow string) {
	type StudentProgress struct{
		Name       string
		StudentID  int
		LabNum     int
	}
	type Labs struct {
		Start    time.Time
		End      time.Time
	}

	progress := StudentProgress{Name: SName, StudentID: ID, LabNum: LLabs}
  studRow = ``
	for i := 0; i < progress.StudentID; i ++ {
		studRow = fmt.Println(`<tr><td>%s</td><td>%d</d>`, progress.Name, progress.StudentID)
		for j := range progress.Labs {
			c := ColorMe(progress.StudentID[i].LabNum[j].Start, progress.StudentID[i].LabNum[j].End)
      studRow += c
		}
		studRow += `</tr>`
  	}
	return
}

func ColorMe(st time.Time, et time.Time)(color string) {
  loc, _ := time.LoadLocation("UTC")
	ct := time.Now().In(loc)
  hourAgo := ct.Add(-1 * time.Hour)
	if et > (hourAgo){
		fmt.Printf(`<td bgcolor="#11FF00"></td>`)
    color := `<td bgcolor="#11FF00"></td>`
	// } else if et > (ct -(2 * hourAgo)) {
	// 	fmt.Printf(`<td bgcolor="#00FFAA"></td>`)
	// 	color := `<td bgcolor="#00FFAA"></td>`
	// } else if et > (ct - hourAgo) {
	// 	fmt.Printf(`<td bgcolor="#00FF00"></td>`)
	// 	color := `<td bgcolor="#00FF00"></td>`
	// } else if st > (ct -hourAgo) {
	// 	fmt.Printf(`<td bgcolor="#FFFF00"></td>`)
	// 	color := `<td bgcolor="#FFFF00"></td>`
	// } else if st > (ct -(3 * hourAgo)) {
	// 	fmt.Printf(`<td bgcolor="#000000"></td>`)
	// 	color := `<td bgcolor="#000000"></td>`
	} else {
			fmt.Printf(`<td bgcolor="#FFFFFF"></td>`)
			color := `<td bgcolor="#FFFFFF"></td>`
		}
		return color
	}



func Template() http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "" {
	                        r.URL.Path = "graph.html"
			} else if r.URL.Path == "deploy/robots.txt" {
				http.ServeFile(w, r, "robots.txt")
				return
			} else if r.URL.Path == "deploy/sitemap.xml" {
				http.ServeFile(w, r, "sitemap.xml")
				return
			}

		  //example := "alta3.com/graph/sam?labs=37&lst=60&hst=79"
      labs := r.FormValue("labs")
			lowStudent := r.FormValue("lst")
			highStudent := r.FormValue("hst")

			fmt.Printf("Labs = %s, Lowest Student Number = %s, Highest Student Number = %s", labs, lowStudent, highStudent)
		  labDigit, err := strconv.Atoi(labs)
			    if err != nil {println("ERROR IN STRING CONVERT")}
			lStudent, err := strconv.Atoi(lowStudent)
				 if err != nil {println("ERROR IN STRING CONVERT")}
			hStudent, err := strconv.Atoi(highStudent)
			   if err != nil {println("ERROR IN STRING CONVERT")}
			fmt.Printf("\n Labs in digits: %d, L: %d, H: %d", labDigit, lStudent, hStudent)


			labList := []int{}
			for x := 0; x <= (labDigit); x ++ {
					labList = append(labList, x)
			}

			fmt.Println(labList)
      studentList := []int{}
      studRange := hStudent - lStudent
			for e := lStudent; e <= hStudent; e ++ {
					studentList = append(studentList, e)
			}

			fmt.Println(studentList)

      initTable := iterTable("Sam", 1, 12)


			// initTable := `<table><tr><th>Student Name</th><th>Student ID</th>`
			// // for y := 0; y <= (labDigit); y ++ {
			// // 	p := fmt.Sprintf(`<th>%d</th>`, y)
			// // 	initTable += p
			// // }
			//
			// initTable += `</tr>`
			// labGridyx := ""
			// studentRow := fmt.Sprintf(`<tr><td>Name</td><td>ID</td>`)
			//
      // for n := 0; n < studRange ; n++ {
			//     for num := 0; num <= 20 ; num ++ {
			// 		    studentRow += `<td></td>`
			//     }
			//     studentRow += `</tr><tr><td>Name</td><td>ID</td>`
		  //     }
			//
			// words := strings.Fields(labGridyx)
			// fmt.Println(words, len(words))
    	// initTable += studentRow
			// fmt.Printf("\nLab Grids: %s", labGridyx)


			t := template.New("graph.html")
			t, _ = t.Parse(initTable)
			t.Execute(w, "graph.html")
      tp := path.Join("deploy/templates", "graph.html")


			info, err := os.Stat(tp)
			if err != nil {
				if os.IsNotExist(err) {
					log.Printf("404 - This file does NOT exist: %s", r.URL.Path)
					errorHandler(http.StatusNotFound).ServeHTTP(w, r)
					return
				}
			}
	                //if fp is a directory, send 404
			if info.IsDir() {
	                        fmt.Println("404 - The file request is pointing to a directory! ")
				errorHandler(http.StatusNotFound).ServeHTTP(w, r)
				return
			}
			return
		})
	}


func main() {
	http.Handle("/graph/", http.StripPrefix("/graph/", Template()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
