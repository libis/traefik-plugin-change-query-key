package change_query_key

import (
	"context"
	"errors"
	"net/http"
)

type changeQueryKey struct {
	keyName    string
	newKeyName string
	next       http.Handler
	name       string
	config     *Config
}

// Config holds configuration to be passed to the plugin.
type Config struct {
	KeyName    string `json:"keyname"`
	NewKeyName string `json:"newkeyname"`
}

// CreateConfig populates the Config data object.
func CreateConfig() *Config {
	return &Config{
		KeyName:    "",
		NewKeyName: "",
	}
}

func (q *changeQueryKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	value := r.URL.Query().Get(q.keyName)

	//Delete the old value and add a new one.
	query.Del(q.keyName)
	query.Add(q.newKeyName, value)

	r.URL.RawQuery = query.Encode()
	r.RequestURI = r.URL.RequestURI()
	q.next.ServeHTTP(rw, r)
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.KeyName == "" || config.NewKeyName == "" {
		return nil, errors.New("keyname and newkeyname must be specified in configuration")
	}

	return &changeQueryKey{
		keyName:    config.KeyName,
		newKeyName: config.NewKeyName,
		next:       next,
		name:       name,
		config:     config,
	}, nil
}
