package req

type CreateNewTodo struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
}
