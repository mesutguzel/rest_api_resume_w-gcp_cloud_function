// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"log"
	"net/http"
  "bytes"
  "fmt"
)

type Resume struct {
	Basics       Basics        `json:"basics"`
	Work         []Work        `json:"work"`
	Volunteer    []Volunteer   `json:"volunteer"`
	Education    []Education   `json:"education"`
	Awards       []string      `json:"awards"`
	Publications []Publication `json:"publications"`
	Skills       []Skill       `json:"skills"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	References   []Reference   `json:"references"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Website  string    `json:"website"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Work struct {
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Website    string   `json:"website"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Website      string   `json:"website"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Gpa         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}

type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	Website     string `json:"website"`
	Summary     string `json:"summary"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

func getResume() *Resume {
	return &Resume{
		Basics: Basics{
			Name:    "Mesut Doganguzel",
			Label:   "Full Stack Developer",
			Picture: "https://media-exp1.licdn.com/dms/image/C4E03AQFCjzCTWYOo6g/profile-displayphoto-shrink_200_200/0/1564465950841?e=1648080000&v=beta&t=FDxvbrMWFgIf4RMhcMMnl2d7eifzCWrhLaZGjgP6ac8",
			Email:   "mesutguzel3501@gmail.com",
			Phone:   "217-390-2877",
			Website: "http://54.224.67.96/personal/navbar.html",
			Summary: "Student. AWS Enthusiast. Crypto King. Gained specialized knowledge in database building and management - Honed project management skills for IT systems Imaginative Designer with excellent analytical, research and collaboration skills. Experience working on Web Development.",
			Location: Location{
				Address:     "",
				PostalCode:  "",
				City:        "Toronto",
				CountryCode: "Canada",
				Region:      "ON",
			},
			Profiles: []Profile{
				{
					Network:  "Linkedin",
					Username: "Mesut Doganguzel",
					URL:      "https://www.linkedin.com/in/mesut-doganguzel-b5a04618b/",
				}, {
					Network:  "GitHub",
					Username: "mesutguzel",
					URL:      "https://github.com/mesutguzel",
				}, {
					Network:  "Twitter",
					Username: "MesutGuzel",
					URL:      "https://twitter.com/MesutGuzel1",
				},
			},
		},
		Work: []Work{
			{
				Company:   "Innunco Academy",
				Position:  "IT Infrastructure Project Coordinator(CO-OP)",
				Website:   "---",
				StartDate: "2021-09-01",
				EndDate:   "2021-12-31",
				Summary:   "Innunco Academy focuses on the students, our future. Every student is different in learning and we believe that everyone has the potential to grow. We will educate and assist students with the essential skills for post-secondary.",
				Highlights: []string{
					"• Maintained and edited company website, Managed Monday.com during Company Project.",
					"• Outlined work plans, determined resources, wrote timelines and generated initial budgets as part of project scope determination.",
					"• Oversaw large portfolio of projects to support teams, report progress and influence positive outcomes for key stakeholders.",
					"• Sped up and simplified application and resource provisioning.",
					"• Managed project teams in Agile environment, realizing success through application of SDLC methodologies and exceptional leadership skills",
				},
			},
		},
		Volunteer: []Volunteer{
			{
				Organization: "Estatetial",
				Position:     "AppDeveloper Co-Designer",
				Website:      "----",
				StartDate:    "2021-01-01",
				EndDate:      "2021-05-31",
				Summary:      "Real-Estate Application, Start-up Company",
				Highlights: []string{
					"• For Real Estate Application Developed and Android and IOS application in Android Studio with add, delete and download functions.",
					"• Obtained approval of concepts by submitting rough drafts to management or to client.",
					"• Developed design deliverables that elevated, differentiated and functioned on-brand and on-strategy.",
					"• Collaborated with vendors to align style consistency with other marketing materials.",
					"• Created over 100 visual concepts either by hand or with assistance of computer programs in 3 months.",
				},
			},
		},
		Education: []Education{
			{
				Institution: "Algonquin College",
				Area:        "Information Technology",
				StudyType:   "Advanced Diploma",
				StartDate:   "2020-09-01",
				EndDate:     "2022-08-26",
				Gpa:         "3.68/4.0",
				Courses:     []string{},
			},
		},
		Publications: []Publication{
			{
				Name:        "Tips for Working from Home",
				Publisher:   "YMCA",
				ReleaseDate: "2020-09-30",
				Website:     "https://newcomersincanada.ca/life-during-covid-19/blog-tips-of-working-from-home/",
				Summary:     "Tips for Working from Home",
			},
		},
		Skills: []Skill{
			{
				Name:  "AWS",
				Level: "Intermediate",
				Keywords: []string{
					"Cloud", "EC2", "S3", "IAM",
				},
			}, {
				Name:  "JavaScript",
				Level: "Beginner",
				Keywords: []string{
					"Web", "Front End",
				},
			}, {
				Name:  "Go",
				Level: "Beginner",
				Keywords: []string{
					"Golang",
				},
			}, {
				Name:  "Docker",
				Level: "Beginner",
				Keywords: []string{
					"Containers", "Images", "Networking",
				},
			}, {
				Name:  "Linux",
				Level: "Intermediate",
				Keywords: []string{
					"Bash", "Shell",
				},
			}, {
				Name:  "SQL",
				Level: "Intermediate",
				Keywords: []string{
					"PostgreSQL", "MySQL",
				},
			}, {
				Name:     "GitHub",
				Level:    "Intermediate",
				Keywords: []string{},
			}, {
				Name:     "Java",
				Level:    "Intermediate",
				Keywords: []string{},
			}, {
				Name:     "Python",
				Level:    "Beginner",
				Keywords: []string{},
			},
		},
		Languages: []Language{
			{

				Language: "English",
			},
		},
		Interests: []Interest{
			{
				Name:     "Crypto",
				Keywords: []string{},
			}, {
				Name:     "Travel",
				Keywords: []string{},
			},
		},
		References: []Reference{
			{
				Name: "Upon Request",
			},
		},
	}
}

func (input Resume) formatResume() string {
	bytesBuffer := new(bytes.Buffer)
	json.NewEncoder(bytesBuffer).Encode(&input)

	responseBytes := bytesBuffer.Bytes()

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, responseBytes, "", "  ")
	if error != nil {
		log.Println("JSON parse error: ", error)
	}
	formattedResume := string(prettyJSON.Bytes())
	return formattedResume
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	resume := getResume()
	formattedResume := resume.formatResume()

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, formattedResume)
}

