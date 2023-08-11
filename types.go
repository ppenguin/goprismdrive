package client

type User struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
}

type UserFull struct {
	Id          uint64 `json:"id"`
	Token       string `json:"access_token"`
	DisplayName string `json:"display_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	TsCreated   string `json:"created_at"`
	TsUpdated   string `json:"updated_at"`
}

type ReqLogin struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	DeviceName string `json:"device_name"`
}

type ResLoginOK struct {
	Status string   `json:"status"`
	User   UserFull `json:"user"`
}

type ResLoginErr struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// Errors struct {}
}

type FileEntry struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	BlobName    string  `json:"file_name"`
	Size        uint64  `json:"file_size"`
	ParentId    uint64  `json:"parent_id"`
	Parent      string  `json:"parent"` // type???
	Thumbnail   string  `json:"thimbnail"`
	Mime        string  `json:"mime"`
	Url         string  `json:"url"`
	Hash        string  `json:"hash"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	TsDeleted   string  `json:"deleted_at"`
	TsCreated   string  `json:"created_at"`
	TsUpdated   string  `json:"updated_at"`
	Path        string  `json:"path"`
	Users       []*User `json:"users"`
}

type FileEntries []*FileEntry
