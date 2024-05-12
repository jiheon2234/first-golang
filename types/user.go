package types

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  int(c.Age),
	}
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`

	UpdateAge int `json:"updateAge" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
}
