# rttables

Package rttables use to parse and change `/etc/iproute2/rt_tables` .

Recently i'm writing a robot software, it controls Ethernet and mobile network interfaces. Its function is to ensure that the Ethernet interface is the default route. When the Ethernet cannot connect to the internet, it switch the default route to mobile network.

After search about linux with multiple interfaces, i think the mobile network interface needs a seperating routing table, so that we can define a default route for it. Now when i dial to a remote server by setting specified mobile network interface, routing service will match the packets and send it out by corresponding interface.

That's why i created this package, it would use to ensure the user-defined routing tables are created, and we couldn't worry its duplicated.

## Install

```shell
go get github.com/laeo/go-rttables
```

## Example

```golang
// 1. Parse system-wide rt_tables file.
rows, _ := rttables.Parse()
// 2. Check if table eth2 exists.
exists := rttables.ContainsIn(rows, Table{ID: 1, Name: "eth2"})
// 3. If not exist, add it.
if !exists {
    rows = append(rows, Table{ID: 1, Name: "eth2"})
    err := rttables.Patch(rows)
    if err != nil {
        // So handle error here.
        log.Fatal(err)
    }
}
```

## License

The MIT License.
