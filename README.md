# Traefik Query Key Change Middleware Plugin

This Traefik middleware plugin you to change the name of a query key.

## Static Configuration

### FILE

```yaml
experimental:
  plugins:
    change-query-key:
      moduleName: github.com/libis/traefik-plugin-change-query-key
      version: v0.1.1
```

## Dynamic Configuration

### FILE

```yaml
...
  middlewares:
    rename_query_key:                #Change the key name, for example /foo?u=U123456 -> /foo?Username=U123456
      plugin:
        change-query-key:
          keyname:               "u"
          newkeyname:            "Username"
...
```
