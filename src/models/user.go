package models

type User struct {
	UserID string `json:"userID,omitempty" dynamodbav:"UserID"`

	Username  string `json:"username,omitempty" dynamodbav:"Username"`
	Email     string `json:"email,omitempty" dynamodbav:"Email"`
	FirstName string `json:"firstName,omitempty" dynamodbav:"FirstName"`
	LastName  string `json:"lastName,omitempty" dynamodbav:"LastName"`
	Password  string `json:"password,omitempty" dynamodbav:"Password"`
	MainSort  string `json:"mainSort,omitempty" dynamodbav:"MainSort"`

	Follows       []string `json:"follows,omitempty" dynamodbav:"Follows"`
	TotalCaptures int      `json:"TotalCaptures,omitempty" dynamodbav:"TotalCaptures"`
	Color1        string   `json:"color1,omitempty" dynamodbav:"Color1"`
	Color2        string   `json:"color2,omitempty" dynamodbav:"Color2"`
}
