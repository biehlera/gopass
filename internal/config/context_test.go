package config

import (
	"context"
	"testing"

	"github.com/gopasspw/gopass/pkg/ctxutil"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx := context.Background()

	sc := &Config{
		ConfirmRecipients: false,
	}

	// should return the default value from the store config
	assert.Equal(t, false, ctxutil.IsConfirm(sc.WithContext(ctx)))

	// after overwriting the noconfirm value in the context,
	// it should not be overwritten by the store config value
	ctx = ctxutil.WithConfirm(ctx, true)
	assert.Equal(t, true, ctxutil.IsConfirm(sc.WithContext(ctx)))
}
