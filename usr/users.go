package usr

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mrinalwahal/encore/usr/ent"
	"github.com/mrinalwahal/encore/usr/ent/schema"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string       `json:",omitempty"`
	Users   []*ParamUser `json:",omitempty"`
	User    *ParamUser   `json:",omitempty"`
}

type ParamUser struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale string `json:"locale,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// DefaultRole holds the value of the "default_role" field.
	DefaultRole string `json:"default_role,omitempty"`
	// IsAnonymous holds the value of the "is_anonymous" field.
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// TotpSecret holds the value of the "totp_secret" field.
	TotpSecret string `json:"totp_secret,omitempty"`
	// ActiveMfaType holds the value of the "active_mfa_type" field.
	ActiveMfaType string `json:"active_mfa_type,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata *schema.Metadata `json:"metadata,omitempty"`
}

//encore:api method=POST public path=/system/users/add
func AddUser(ctx context.Context, user ParamUser) (*Response, error) {

	//	Generate Password Hash
	hashedPassword, err := generatePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}

	data, err := client.User.Create().
		SetID(rand.Int()).
		SetPasswordHash(string(hashedPassword)).
		SetNillableUsername(&user.Username).
		SetName(user.Name).
		SetNillableActiveMfaType(&user.ActiveMfaType).
		SetCreatedAt(time.Now()).
		SetDisabled(user.Disabled).
		SetEmail(user.Email).
		SetIsAnonymous(user.IsAnonymous).
		SetNillableLocale(&user.Locale).
		SetMetadata(user.Metadata).
		SetNillableAvatarURL(&user.AvatarURL).
		SetNillablePhone(&user.Phone).
		SetNillableTotpSecret(&user.TotpSecret).
		Save(ctx)

	return &Response{User: marshalResourceUser(*data)}, err
}

//encore:api public path=/system/user/:id
func GetUser(ctx context.Context, id int) (*Response, error) {
	user, err := client.User.Get(ctx, id)
	return &Response{User: marshalResourceUser(*user)}, err
}

//encore:api public path=/system/users
func GetAllUsers(ctx context.Context) (*Response, error) {
	data, err := client.User.Query().Select().All(ctx)

	var users []*ParamUser
	for _, item := range data {
		users = append(users, marshalResourceUser(*item))
	}

	return &Response{Users: users, Message: fmt.Sprintf("%v users fetched", len(users))}, err
}

//encore:api method=PATCH public path=/system/user/:id/update
func UpdateUser(ctx context.Context, id int, user ParamUser) error {
	return client.User.Update().
		SetNillableUsername(&user.Username).
		SetNillableActiveMfaType(&user.ActiveMfaType).
		SetNillableLocale(&user.Locale).
		SetNillableAvatarURL(&user.AvatarURL).
		SetNillablePhone(&user.Phone).
		Exec(ctx)
}

//encore:api method=PATCH public path=/system/user/:id/update/metadata
func UpdateUserMetadata(ctx context.Context, id int, metadata schema.Metadata) error {
	return client.User.Update().
		SetMetadata(&metadata).
		Exec(ctx)
}

//encore:api method=PATCH public path=/system/user/:id/update/disable
func DisableUser(ctx context.Context, id int) error {

	user, err := GetUser(ctx, id)
	if err != nil {
		return err
	}

	return client.User.Update().
		SetDisabled(!user.User.Disabled).
		Exec(ctx)
}

//encore:api method=PATCH public path=/system/user/:id/update/anonymous
func SetAnonymity(ctx context.Context, id int) error {

	user, err := GetUser(ctx, id)
	if err != nil {
		return err
	}
	return client.User.Update().
		SetIsAnonymous(!user.User.IsAnonymous).
		Exec(ctx)
}

//encore:api public path=/system/user/:id/delete
func DeleteUser(ctx context.Context, id int) error {
	return client.User.DeleteOneID(id).Exec(ctx)
}

func generatePasswordHash(payload string) ([]byte, error) {

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(payload))
	return hashedPassword, err
}

func marshalResourceUser(payload ent.User) *ParamUser {
	return &ParamUser{
		ID:            payload.ID,
		Name:          payload.Name,
		Email:         payload.Email,
		Phone:         payload.Phone,
		Disabled:      payload.Disabled,
		AvatarURL:     payload.AvatarURL,
		Locale:        payload.Locale,
		DefaultRole:   payload.DefaultRole,
		IsAnonymous:   payload.IsAnonymous,
		TotpSecret:    payload.TotpSecret,
		ActiveMfaType: payload.ActiveMfaType,
		Metadata:      payload.Metadata,
	}
}
