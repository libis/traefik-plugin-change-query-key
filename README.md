# Traefik Query Key Change Middleware Plugin

This Traefik middleware plugin allows you to change the name of a query key. Optionally, if the query key is not present then it gets created with 'newkeyname' and 'value'.

## Static Configuration

### FILE

```yaml
experimental:
  plugins:
    change-query-key:
      moduleName: github.com/libis/traefik-plugin-change-query-key
      version: v0.2.0
```

## Dynamic Configuration

### FILE

```yaml
...
  middlewares:
    rename_query_key:                         #Change the key name, for example /foo?u=U123456 -> /foo?Username=U123456
      plugin:
        change-query-key:
          keyname:               "u"          # Required.
          newkeyname:            "Username"   # Required.
          createifmissing:       true         # Optional, default false.
          value:                 "foo"        # If createifmissing is true then this is required.
...
```
