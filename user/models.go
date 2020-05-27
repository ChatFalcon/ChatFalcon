package user

// Permissions defines the model which is used to check a users permission.
type Permissions struct {
	// Admin overrides all.
	Admin bool `json:"admin" bson:"admin"`
}

// User defines the model which is used for the user.
type User struct {
	Username string `json:"username" bson:"_id"`
	HashedSaltedPassword string `json:"hashedSaltedPassword" bson:"hashedSaltedPassword"`
	Email string `json:"email" bson:"email"`
	Confirmed bool `json:"confirmed" bson:"confirmed"`
	Perms *Permissions `json:"perms" bson:"perms"`
	Roles []string `json:"roles" bson:"roles"`
	PFPUrl string `json:"pfpUrl" bson:"pfpUrl"`
}
