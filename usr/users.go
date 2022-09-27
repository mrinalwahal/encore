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

type response struct {
	Message string
	Users   []*ent.User
}

type NewUser struct {
}

//encore:api method=POST public path=/system/user/add
func AddUser(ctx context.Context, user ent.User) (*ent.User, error) {

	//	Generate Password Hash
	hashedPassword, err := generatePasswordHash(user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return client.User.Create().
		SetID(rand.Int()).
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
		SetPasswordHash(string(hashedPassword)).
		SetNillablePhone(&user.Phone).
		SetNillableTotpSecret(&user.TotpSecret).
		Save(ctx)
}

//encore:api public path=/system/user/:id
func GetUser(ctx context.Context, id int) (*ent.User, error) {
	return client.User.Get(ctx, id)
}

//encore:api public path=/system/users
func GetAllUsers(ctx context.Context) (*response, error) {
	users, err := client.User.Query().Select().All(ctx)
	return &response{Users: users, Message: fmt.Sprintf("%v users fetched", len(users))}, err
}

//encore:api method=PATCH public path=/system/user/:id
func UpdateUser(ctx context.Context, id int, user ent.User) error {
	return client.User.Update().
		SetNillableUsername(&user.Username).
		SetNillableActiveMfaType(&user.ActiveMfaType).
		SetNillableLocale(&user.Locale).
		SetNillableAvatarURL(&user.AvatarURL).
		SetNillablePhone(&user.Phone).
		Exec(ctx)
}

//encore:api method=PATCH public path=/system/user/:id/metadata
func UpdateUserMetadata(ctx context.Context, id int, metadata schema.Metadata) error {
	return client.User.Update().
		SetMetadata(&metadata).
		Exec(ctx)
}

//encore:api method=PATCH public path=/system/user/:id/disable
func DisableUser(ctx context.Context, id int, disable bool) error {
	return client.User.Update().
		SetDisabled(disable).
		Exec(ctx)
}

//encore:api method=PATCH public path=/system/user/:id/anonymous
func SetAnonymity(ctx context.Context, id int, anonymous bool) error {
	return client.User.Update().
		SetIsAnonymous(anonymous).
		Exec(ctx)
}

//encore:api public path=/system/user/delete/:id
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
