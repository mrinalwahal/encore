package usr

import (
	"context"
	"math/rand"

	"github.com/mrinalwahal/encore/usr/ent"
)

//encore:api method=POST public path=/system/user/add
func AddUser(ctx context.Context, user ent.User) (*ent.User, error) {
	return client.User.Create().
		SetID(rand.Int()).
		SetName(user.Name).
		Save(ctx)
}
