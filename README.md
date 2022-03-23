# go-vtm

To import the library:
```go
import "github.com/pulse/vtm"
```

To connect directly to a Virtual Traffic Manager instance:
```go
trafficManager := vtm.NewVirtualTrafficManager("https://vtm1:9070/api", "admin", "P@55w0rD!")
```

To connect to a Virtual Traffic Manager instance via a Services Director proxy:
```go
trafficManager := vtm.NewVirtualTrafficManager("https://sd1:8100/api/tmcm/2.5/instance/vtm1", "sd_admin", "SD_P@55w0rD!")
```

## Single-occurrence resources
Some resources, for example global settings, occur as single occurrences.  These resources
support the Get<*ResourceType*> function and the <*Instance*>.Apply() method.  NB. values are 
stored with pointers.
```go
globalSettings, err := trafficManager.GetGlobalSettings()
*globalSettings.Connection.IdleTimeout = 22
*globalSettings.FaultTolerance.AutoFailback = false
globalSettings.Apply()
```
NB. settings will not be applied until the Apply() method is called.

## Multiple-occurrence resources
Most resource types, for example pools, virtual servers and TrafficScript rules, have many instances.
These resources support the List<*ResourceType*>s, Get<*ResourceType*>, New<*ResourceType*> and Delete<*ResourceType*>
functions.

### Listing objects of a certain type
To get a list of the names of each instance of a configuration object:
```go
vsList, err := trafficManager.ListVirtualServers()
for _, vsName := range vsList {
	fmt.Println("Virtual server: ", vsName)
}
```

### Getting a specific configuration object
To get a specific instance of a configuration object we use the Get<*ResourceType*> function.
This applies to both JSON-based and text-based resources:
```go
myVs, err := trafficManager.GetVirtualServer("my_virtual_server")
myRule, err := trafficManager.GetRule("my_trafficscript_rule")
```

### Reading values from a configuration object
Some configuration objects are data structures (ie. JSON-based in the API) whilst some 
are text-based.  Get<*ResourceType*>() calls return structures or strings respectively.

For JSON-based objects (NB. values are stored with pointers):
```go
fmt.Println("Virtual server port: ", *myVs.Basic.Port)
fmt.Println("Virtual server default pool: ", *myVs.Basic.Pool)
```

For text-based object:
```go
fmt.Printf("My TrafficScript Rule: \n%v\n", myRule)
```

### Altering values on a configuration object
JSON-based and text-based resources have different methods for altering their data.
JSON-based resources can have their field values changed locally, and then with a call to
the <*Instance*>.Apply() method, the changes are applied to the vTM:
```go
*myVs.Basic.Port = 443
*myVs.WebCache.Enabled = true
myVs.Apply()
```

As text-based resources are a single piece of data, they use the Set<*ResourceType*> function
to either create a new or overwrite an existing resource:
```go
err := trafficManager.SetRule("my_trafficscript_rule", "log.info('I overwrote an existing rule');")
err := trafficManager.SetRule("new_rule", "log.info('I created a new rule');")
```

### Creating new configuration objects
Creating new text-based objects was shown in the previous section (using the Set<*ResourceType*> function).
Creating JSON-based objects is done using the New<*ResourceType*> functions.  Depending on the resource type,
initialization parameters may be required (see the function signatures or the REST API manual).
*NOTE: the New<*ResourceType*>

To create a new instance of a configuration object with initialization parameters (in this case, the 
default pool and the listen port, as required to create a virtual server resource):
```go
newVs, err := trafficManager.NewVirtualServer("new_virtual_server", "discard", 8080)
```

To create a new instance of an object that does not require initialization parameters:
```go
newPool, err := trafficManager.NewPool("new_pool")
```

### Deleting configuration objects
To delete a specific instance of a configuration object (same for JSON- and text-based objects):
```go
err := trafficManager.DeleteVirtualServer("new_virtual_server")
err := trafficManager.DeleteRule("new_rule")
```

## Statistics objects
Virtual Traffic Manager provides many read-only counters and statistics.  These can be accessed 
using the Get<*Resource*>Statistics functions.

Functions for stats on single-occurrence objects take no arguments:
```go
globalStats, err := trafficManager.GetGlobalsStatistics()
fmt.Printf("Data memory usage: %d\n", globalStats.DataMemoryUsage)
fmt.Printf("Uptime: %d\n", globalStats.UpTime)
```

Functions for stats on multiple-occurrence objects take an object name as the only argument:
```go
vsStats, err := trafficManager.GetVirtualServersStatistics("my_vs")
fmt.Printf("Bytes In: %d\n", vsStats.BytesIn)
fmt.Printf("Bytes Out: %d\n", vsStats.BytesOut)

```

## System information resources
Virtual Traffic Manger provides some resources containing current statuses:
```go
vtmState, err := trafficManager.GetSystemState()
for _, vtmError := range vtmState.State.Errors {
	fmt.Printf("Error: %s\n", vtmError)
}

vtmInformation, err := trafficManager.GetSystemInformation()
fmt.Printf("vTM UUID: %s\n", vtmInformation.Information.Uuid)
```

## Errors
All functions/methods will return a pointer to an error as either their second or only return value.
If this is a nil pointer, no error occurred.  If not, an error object will be available:
```go
myVs, err := trafficManager.GetVirtualServer("my_non_existent_vs")
if err != nil {
	fmt.Println("ErrorId: ", err.ErrorId)
	fmt.Println("ErrorText: ", err.ErrorText)
	fmt.Println("ErrorInfo: ", err.ErrorInfo)
}
```
would result in:
```
ErrorId:  resource.not_found
ErrorText:  Resource '/api/tm/5.0/config/active/virtual_servers/my_non_existent_vs' does not exist
ErrorInfo:
```
Note that the ErrorInfo field is only populated in the case of particular errors; often it is empty.

## Copyright and License Acknowledgement

Copyright &copy; 2022, Pulse Secure LLC. Licensed under the terms of the
MPL 2.0. See the LICENSE file for details.
