package types


/*
Package types is used to store the context struct which
is passed while templates are executed.
*/
//Task is the struct used to identify tasks
type Task struct {
	Id           int           `json:"id"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	Created      string        `json:"created"`
	Priority     string        `json:"priority"`
	Category     string        `json:"category"`
	Referer      string        `json:"referer,omitempty"`

}
