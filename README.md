## HOWTO
* Assign `GeoIPDBDirPath` to your `GeoLite2-Country.mmdb`, `GeoLite2-ASN.mmdb`, `GeoLite2-City.mmdb` folder path before use it.

### Example
```
kkgeoip.GeoIPDBDirPath = conf.Config().DataStore.GeoIPPath
kkgeoip.CountryCode("1.1.1.1")
```

### Source
`maxmind.com`