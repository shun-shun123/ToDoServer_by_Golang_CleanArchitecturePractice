package domain

type ToDo struct {
	ID 			int 	`json:"id"`
	Content 	string 	`json:"content"`
	CreateDate 	string 	`json:create_date`
	Author 		string 	`json:"author"`
}

type ToDos struct {
	Todos 		[]ToDo 	`json:"todos"`
}
