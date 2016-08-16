# example

This is example of `httptraceutils`,

```bash
$ go run main.go
2016/08/16 09:10:44 [GetConn] HostPort: google.com:443
2016/08/16 09:10:44 [DNSStart] Host: google.com
2016/08/16 09:10:44 [DNSDone] Addrs: 172.217.25.78
2016/08/16 09:10:44 [DNSDone] Coalesced: false
2016/08/16 09:10:44 [DNSDone] Error: nil
2016/08/16 09:10:44 [ConnectStart] Network: tcp
2016/08/16 09:10:44 [ConnectStart] Addr: 172.217.25.78:443
2016/08/16 09:10:44 [ConnectDone] Network: tcp
2016/08/16 09:10:44 [ConnectDone] Addr: 172.217.25.78:443
2016/08/16 09:10:44 [ConnectDone] Error: nil
2016/08/16 09:10:45 [GotConn] LocalAddr: 10.0.1.16:49280
2016/08/16 09:10:45 [GotConn] RemoteAddr: 172.217.25.78:443
2016/08/16 09:10:45 [GotConn] Reused: false
2016/08/16 09:10:45 [GotConn] WasIdle: false
2016/08/16 09:10:45 [WroteHeaders]
2016/08/16 09:10:45 [WroteRequest] Error: nil
2016/08/16 09:10:45 [GotFirstResponseByte]
```
