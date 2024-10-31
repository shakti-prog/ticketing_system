package utils

type Credentials struct {
	Useremail string `json:"useremail"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

type ProjectParams struct {
	ProjectName string `json:"project_name"`
	Created_By  int64  `json:"created_by"`
}

type TicketParams struct {
	AssigneeId  int64  `json:"assigneeId"`
	ReporterId  int64  `json:"reporterId"`
	ProjectId   int64  `json:"projectId"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
}

type CommentParams struct {
	TicketId    int64  `json:"ticketId"`
	Description string `json:"description"`
	UserId      int64  `json:"userId"`
}

type UpdateTicketParams struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
