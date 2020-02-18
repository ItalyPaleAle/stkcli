## stkcli deploy

Deploy an app

### Synopsis

Deploys an app to a site.

This command tells the node to deploy the app (already uploaded beforehand) with the specific name and version to a site identified by the domain option.


```
stkcli deploy [flags]
```

### Options

```
  -a, --app string       app's bundle name (required)
  -d, --domain string    primary domain name (required)
  -h, --help             help for deploy
  -S, --http             use HTTP protocol, without TLS
  -k, --insecure         disable TLS certificate validation (default true)
  -n, --node string      node address or IP (default "localhost")
  -P, --port string      port the node listens on (default "2265")
  -v, --version string   app's bundle version (required)
```

### SEE ALSO

* [stkcli](stkcli.md)	 - Manage a Statiko node
