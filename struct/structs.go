package _struct

type English struct {
	ID                string `json:"id"`
	Word              string `json:"word"`
	Transcription     string `json:"transcription"`
	Translate         string `json:"translate"`
	Example1          string `json:"example1"`
	TranslateExample1 string `json:"translateExample1"`
	Example2          string `json:"example2"`
	TranslateExample2 string `json:"translateExample2"`
}

type Words struct {
	CountOfWords string `json:"countOfWords"`
	Login        string `json:"login"`
}

type RegistrationRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type BoolAnswer struct {
	BoolAnswer bool `json:"hasRegLog"`
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
