package main

import (
	"log"
	"os"
	"text/template"
)
type Client struct {
		Honorific string 
		LastName string
		Attend bool
		Donate bool
		FutureEvents []string
	}
	
func main() {

const letter =`Hello {{.Honorific}} {{.LastName}},
{{if .Attend}}
I am writing this letter to personally thanks you for attending the XYZ fundraiser. 
{{else}}
I am sorry to hear that you were not able to attend the XYZ fundraiser. 
{{end}}
{{if .Donate}}
Thank you for the generous donation. We truly appreciate your support.  
{{end}}
Just a friendly reminder,out upcoming events include the following, which you are more than welcome to attend:
{{range .FutureEvents}} {{.}}
{{end}}
Best Wishes,

XYZ
`

	var fEvents = []string{"Wake-a-thon on June 3rd",
		"Silent Auction on June 15th",
		"Car Wash on June 25th"}
	
	var ListOfClients = []Client{
		{"Mr", "Vang", true, true, fEvents},
		{"Mrs", "Smith", false, true, fEvents},
		{"Mr", "Johnson", true, false, fEvents},
		{"Ms", "Walker", false, false, fEvents},
	}

	t := template.Must(template.New("letter").Parse(letter))

	for _, r := range ListOfClients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
