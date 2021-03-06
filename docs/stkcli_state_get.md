## stkcli state get

Retrieve state and save to file

### Synopsis

Dumps the state of the node to a file or stdout (if the `--output` parameter is not set).

The state is a JSON document containing the list of sites and apps currently configured in the web server. You can store the state in a file for backups, or to restore it to another node using the `state set` command.


```
stkcli state get [flags]
```

### Options

```
  -h, --help          help for get
  -N, --node string   node address or IP
  -o, --out string    output file where to store state; if not set, print to stdout
  -P, --port string   port the node listens on
```

### SEE ALSO

* [stkcli state](stkcli_state.md)	 - Get or restore state

