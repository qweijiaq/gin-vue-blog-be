package log_stash

import (
	"fmt"
	"net"
	"server/global"
)

// FormatBytes 格式化输出字节单位
func FormatBytes(size int64) string {
	_size := float64(size)
	uints := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	// 1
	// 1025 1.0KB
	//
	var i int = 0
	for _size >= 1024 && i < len(uints)-1 {
		_size /= 1024
		i++
	}
	return fmt.Sprintf("%.2f %s", _size, uints[i])

}

// ExternalIp 判断是否是外网地址
func ExternalIp(ip string) (ok bool) {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	ip4 := IP.To4()
	if ip4 == nil {
		return false
	}
	if !IP.IsPrivate() && !IP.IsLoopback() {
		return true
	}
	return false

}

func getAddr(ip string) (addr string) {
	if !ExternalIp(ip) {
		return "内网地址"
	}
	citys, err := global.AddrDB.City(net.ParseIP(ip))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 国家
	country := citys.Country.Names["zh-CN"]
	// 城市
	city := citys.City.Names["zh-CN"]
	// 省份
	var subdivisions string
	if len(citys.Subdivisions) > 0 {
		subdivisions = citys.Subdivisions[0].Names["zh-CN"]
		return fmt.Sprintf("%s-%s", subdivisions, city)
	}
	if city != "" {
		return fmt.Sprintf("%s-%s", country, city)
	}
	return "未知地址"
}
