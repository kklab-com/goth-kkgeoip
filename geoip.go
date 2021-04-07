package kkgeoip

import (
	"fmt"
	"net"
	"sync"

	kklogger "github.com/kklab-com/goth-kklogger"
	"github.com/oschwald/geoip2-golang"
)

var GeoIPDBDirPath = "/usr/share/GeoIP"
var asnDBInstance *geoip2.Reader
var cityDBInstance *geoip2.Reader
var countryDBInstance *geoip2.Reader
var dbInitOnce sync.Once

func ASN(ipAddr string) *geoip2.ASN {
	_DBInstanceInit()
	ip := net.ParseIP(ipAddr)
	if ip == nil || asnDBInstance == nil {
		return nil
	}

	if asn, err := asnDBInstance.ASN(ip); err == nil {
		return asn
	}

	return nil
}

func City(ipAddr string) *geoip2.City {
	_DBInstanceInit()
	ip := net.ParseIP(ipAddr)
	if ip == nil || cityDBInstance == nil {
		return nil
	}

	if city, err := cityDBInstance.City(ip); err == nil {
		return city
	}

	return nil
}

func Country(ipAddr string) *geoip2.Country {
	_DBInstanceInit()
	ip := net.ParseIP(ipAddr)
	if ip == nil || countryDBInstance == nil {
		return nil
	}

	if country, err := countryDBInstance.Country(ip); err == nil {
		return country
	}

	return nil
}

func CountryCode(ipAddr string) string {
	if country := Country(ipAddr); country != nil {
		return country.Country.IsoCode
	}

	return ""
}

func _DBInstanceInit() {
	dbInitOnce.Do(func() {
		var err error
		var path = fmt.Sprintf("%s/GeoLite2-ASN.mmdb", GeoIPDBDirPath)
		asnDBInstance, err = geoip2.Open(path)
		if err != nil {
			kklogger.Warn("KKGeoIP", fmt.Sprintf("ASN Load Fail, Path: %s, Err: %s", path, err.Error()))
		}

		path = fmt.Sprintf("%s/GeoLite2-City.mmdb", GeoIPDBDirPath)
		cityDBInstance, err = geoip2.Open(path)
		if err != nil {
			kklogger.Warn("KKGeoIP", fmt.Sprintf("City Load Fail, Path: %s, Err: %s", path, err.Error()))
		}

		path = fmt.Sprintf("%s/GeoLite2-Country.mmdb", GeoIPDBDirPath)
		countryDBInstance, err = geoip2.Open(path)
		if err != nil {
			kklogger.Warn("KKGeoIP", fmt.Sprintf("Country Load Fail, Path: %s, Err: %s", path, err.Error()))
		}
	})
}
