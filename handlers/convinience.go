package handlers

import (
	"context"
	"net/http"

	"github.com/FernandoJVideira/velox"
	"github.com/FernandoJVideira/velox/filesystems"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	return h.App.Render.Page(w, r, view, variables, data)
}

func (h *Handlers) uploadFile(r *http.Request, destination, field string, fs filesystems.FS) error {
	return h.App.UploadFile(r, destination, field, fs)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, value interface{}) {
	h.App.Session.Put(ctx, key, value)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) randomString(length int) string {
	return h.App.RandomString(length)
}

func (h *Handlers) encrypt(text string) (string, error) {
	enc := velox.Encription{Key: []byte(h.App.EncryptionKey)}

	encripted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}

	return encripted, nil
}

func (h *Handlers) decrypt(crypto string) (string, error) {
	enc := velox.Encription{Key: []byte(h.App.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}
