# Go Metaname

A Go library implementing the [Metaname](https://metaname.net.nz) API.

## Usage

Create a client with NewMetanameClient, passing it your account reference and API key, then use it to call the needed functions.

```go
client := NewMetanameClient(os.Getenv("ACCOUNT_REF"), os.Getenv("API_KEY"))
client.DeleteDnsRecord(ctx, "example.org", "1234")
```

### Implemented functions

These functions have been implemented:

* [CreateDnsRecord](https://metaname.net/api/1.1/doc#create_dns_record)
* [UpdateDnsRecord](https://metaname.net/api/1.1/doc#update_dns_record)
* [DeleteDnsRecord](https://metaname.net/api/1.1/doc#delete_dns_record)
* [DnsZone](https://metaname.net/api/1.1/doc#dns_zone)
* [ConfigureZone](https://metaname.net/api/1.1/doc#configure_zone)

## Known issues

* ResourceRecords for accounts where the API key was first created before 2023-07-11 will receive Aux values of 0 on MX and SRV records when calling DnsZone. If this affects you, you can email Metaname support and ask them to "flag my account for the numeric aux fix".
