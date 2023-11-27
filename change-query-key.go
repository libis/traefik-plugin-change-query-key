
import (
	"context"
	"errors"
	"net/http"
)

type changeQueryKey struct {
	keyName         string
	newKeyName      string
	createIfMissing bool
	value           string
	next            http.Handler
	name            string
	config          *Config
}

// Config holds configuration to be passed to the plugin.
type Config struct {
	KeyName         string `json:"keyname"`
	NewKeyName      string `json:"newkeyname"`
	CreateIfMissing bool   `json:"createifmissing"`  // If true, create a new key if it doesn't exist. Then value can not be empty
	Value           string `json:"value"`            // The value to set the new key to. Can't be empty if createifmissing is true.
}

// CreateConfig populates the Config data object.
func CreateConfig() *Config {
	return &Config{
		KeyName:         "",
		NewKeyName:      "",
		CreateIfMissing: false,
		Value:           "",
	}
}

func (q *changeQueryKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Get the value associated with the specified key (q.keyName) in the original query parameters.
	value := query.Get(q.keyName)

	// Check if the original key exists in the query parameters.
	if value != ""  {
		// Delete the original key from the query parameters.
		query.Del(q.keyName)
		// Add a new key (q.newKeyName) with the value obtained from the original key.
		query.Add(q.newKeyName, value)
	} else if q.createIfMissing {
		// If the original key doesn't exist, and createifmissing is set, we need to add a new key with a value.
		value = q.value
		query.Add(q.newKeyName, value)
	}

	// Encode the modified query parameters and update the request URL and request URI.
	r.URL.RawQuery = query.Encode()
	r.RequestURI = r.URL.RequestURI()

	// Call the ServeHTTP method of the next HTTP handler in the chain (q.next) with the modified request.
	q.next.ServeHTTP(rw, r)
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.KeyName == "" || config.NewKeyName == "" {
		return nil, errors.New("keyname and newkeyname must be specified in configuration")
	}

	// If createifmissing is set, we need to have a value.
	if config.CreateIfMissing && config.Value == "" {
		return nil, errors.New("createifmissing is set, but value is empty")
	}

	return &changeQueryKey{
		keyName:         config.KeyName,
		newKeyName:      config.NewKeyName,
		createIfMissing: config.CreateIfMissing,
		value:           config.Value,
		next:            next,
		name:            name,
		config:          config,
	}, nil
}
