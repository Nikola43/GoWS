package models

type User struct {
	FingerPrint     string `json:"fingerprint"`
}

func (o *User) InsertNewUser() error {
	return nil
}
