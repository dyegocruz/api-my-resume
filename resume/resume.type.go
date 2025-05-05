package resume

type Position struct {
	Company   string   `json:"company" bson:"company"`
	Title     string   `json:"title" bson:"title"`
	Skills    []string `json:"skills" bson:"skills"`
	IsCurrent bool     `json:"isCurrent" bson:"isCurrent"`
	StartDate string   `json:"startDate" bson:"start_date"`
	EndDate   *string  `json:"endDate" bson:"end_date"` // Usando ponteiro para permitir valores nulos
	Summary   []string `json:"summary" bson:"summary"`
}

type Education struct {
	CourseName  string `json:"course_name" bson:"course_name"`
	Degree      string `json:"degree" bson:"degree"`
	FromYear    string `json:"fromYear" bson:"from_year"`
	Institution string `json:"institution" bson:"institution"`
	ToYear      string `json:"toYear" bson:"to_year"`
}

type Social struct {
	Github   string `json:"github" bson:"github"`
	Linkedin string `json:"linkedin" bson:"linkedin"`
}

type MyResume struct {
	Username        string      `json:"username" bson:"username"`
	Name            string      `json:"name" bson:"name"`
	Email           string      `json:"email" bson:"email"`
	Location        string      `json:"location" bson:"location"`
	Bio             string      `json:"bio" bson:"bio"`
	AvatarURL       string      `json:"avatarUrl" bson:"avatar_url"`
	CurrentPosition string      `json:"currentPosition" bson:"current_position"`
	Skills          []string    `json:"skills" bson:"skills"`
	Social          Social      `json:"social" bson:"social"`
	Positions       []Position  `json:"positions" bson:"positions"`
	Education       []Education `json:"education" bson:"education"`
}
