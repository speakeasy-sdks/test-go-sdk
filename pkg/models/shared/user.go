// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/test-go-sdk/v3/pkg/utils"
	"time"
)

// User - The details of a typical user account
type User struct {
	Country    string     `json:"country"`
	Createdate *time.Time `json:"createdate,omitempty"`
	Email      string     `json:"email"`
	Firstname  string     `json:"firstname"`
	ID         *string    `json:"id,omitempty"`
	Lastname   string     `json:"lastname"`
	Nickname   string     `json:"nickname"`
	Password   string     `json:"password"`
	Updatedate *time.Time `json:"updatedate,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *User) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *User) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *User) GetCreatedate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Createdate
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetFirstname() string {
	if o == nil {
		return ""
	}
	return o.Firstname
}

func (o *User) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *User) GetLastname() string {
	if o == nil {
		return ""
	}
	return o.Lastname
}

func (o *User) GetNickname() string {
	if o == nil {
		return ""
	}
	return o.Nickname
}

func (o *User) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *User) GetUpdatedate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Updatedate
}
