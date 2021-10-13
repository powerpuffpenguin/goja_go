package net

import (
	"net"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`IPv4len`, f.getIPv4len, nil)
	f.Accessor(`IPv6len`, f.getIPv6len, nil)

	f.Accessor(`IPv4bcast`, f.getIPv4bcast, nil)
	f.Accessor(`IPv4allsys`, f.getIPv4allsys, nil)
	f.Accessor(`IPv4allrouter`, f.getIPv4allrouter, nil)
	f.Accessor(`IPv4zero`, f.getIPv4zero, nil)

	f.Accessor(`IPv6zero`, f.getIPv6zero, nil)
	f.Accessor(`IPv6unspecified`, f.getIPv6unspecified, nil)
	f.Accessor(`IPv6loopback`, f.getIPv6loopback, nil)
	f.Accessor(`IPv6interfacelocalallnodes`, f.getIPv6interfacelocalallnodes, nil)
	f.Accessor(`IPv6linklocalallnodes`, f.getIPv6linklocalallnodes, nil)
	f.Accessor(`IPv6linklocalallrouters`, f.getIPv6linklocalallrouters, nil)

	f.Accessor(`DefaultResolver`, f.getDefaultResolver, nil)
	f.Accessor(`ErrClosed`, f.getErrClosed, nil)
	f.Accessor(`ErrWriteToConnected`, f.getErrWriteToConnected, nil)

	f.Set(`JoinHostPort`, net.JoinHostPort)
	f.Set(`LookupAddr`, net.LookupAddr)
	f.Set(`LookupCNAME`, net.LookupCNAME)
	f.Set(`LookupHost`, net.LookupHost)
	f.Set(`LookupPort`, net.LookupPort)
	f.Set(`LookupTXT`, net.LookupTXT)
	f.Set(`ParseCIDR`, net.ParseCIDR)
	f.Set(`Pipe`, net.Pipe)
	f.Set(`SplitHostPort`, net.SplitHostPort)

	f.Set(`InterfaceAddrs`, net.InterfaceAddrs)

	f.Set(`Dial`, net.Dial)
	f.Set(`DialTimeout`, net.DialTimeout)
	f.Set(`FileConn`, net.FileConn)

	f.Set(`Dialer`, Dialer)

	f.Set(`Flags`, Flags)
	f.Accessor(`FlagUp`, f.getFlagUp, nil)
	f.Accessor(`FlagBroadcast`, f.getFlagBroadcast, nil)
	f.Accessor(`FlagLoopback`, f.getFlagLoopback, nil)
	f.Accessor(`FlagPointToPoint`, f.getFlagPointToPoint, nil)
	f.Accessor(`FlagMulticast`, f.getFlagMulticast, nil)

	f.Set(`ParseMAC`, net.ParseMAC)

	f.Set(`IPv4`, net.IPv4)
	f.Set(`LookupIP`, net.LookupIP)
	f.Set(`ParseIP`, net.ParseIP)

	f.Set(`ResolveIPAddr`, net.ResolveIPAddr)

	f.Set(`DialIP`, net.DialIP)
	f.Set(`ListenIP`, net.ListenIP)

	f.Set(`CIDRMask`, net.CIDRMask)
	f.Set(`IPv4Mask`, net.IPv4Mask)

	f.Set(`InterfaceByIndex`, net.InterfaceByIndex)
	f.Set(`InterfaceByName`, net.InterfaceByName)
	f.Set(`Interfaces`, net.Interfaces)

	f.Set(`FileListener`, net.FileListener)
	f.Set(`Listen`, net.Listen)

	f.Set(`LookupMX`, net.LookupMX)
	f.Set(`LookupNS`, net.LookupNS)

	f.Set(`FilePacketConn`, net.FilePacketConn)
	f.Set(`ListenPacket`, net.ListenPacket)

	f.Set(`Resolver`, Resolver)

	f.Set(`LookupSRV`, net.LookupSRV)

	f.Set(`ResolveTCPAddr`, net.ResolveTCPAddr)
	f.Set(`DialTCP`, net.DialTCP)
	f.Set(`ListenTCP`, net.ListenTCP)

	f.Set(`ResolveUDPAddr`, net.ResolveUDPAddr)
	f.Set(`DialUDP`, net.DialUDP)
	f.Set(`ListenMulticastUDP`, net.ListenMulticastUDP)
	f.Set(`ListenUDP`, net.ListenUDP)

	f.Set(`ResolveUnixAddr`, net.ResolveUnixAddr)
	f.Set(`DialUnix`, net.DialUnix)
	f.Set(`ListenUnixgram`, net.ListenUnixgram)
	f.Set(`ListenUnix`, net.ListenUnix)

	f.Set(`isAddr`, isAddr)
	f.Set(`isAddrError`, isAddrError)
	f.Set(`isBuffers`, isBuffers)
	f.Set(`isConn`, isConn)
	f.Set(`isDNSConfigErrorPointer`, isDNSConfigErrorPointer)
	f.Set(`isDNSErrorPointer`, isDNSErrorPointer)
	f.Set(`isDialerPointer`, isDialerPointer)
	f.Set(`isError`, isError)
	f.Set(`isFlags`, isFlags)
	f.Set(`isHardwareAddr`, isHardwareAddr)
	f.Set(`isIP`, isIP)
	f.Set(`isIPAddrPointer`, isIPAddrPointer)
	f.Set(`isIPConnPointer`, isIPConnPointer)
	f.Set(`isIPMask`, isIPMask)
	f.Set(`isIPNetPointer`, isIPNetPointer)
	f.Set(`isInterfacePointer`, isInterfacePointer)
	f.Set(`isInvalidAddrError`, isInvalidAddrError)
	f.Set(`isListenConfigPointer`, isListenConfigPointer)
	f.Set(`isListener`, isListener)
	f.Set(`isMXPointer`, isMXPointer)
	f.Set(`isNSPointer`, isNSPointer)
	f.Set(`isOpErrorPointer`, isOpErrorPointer)
	f.Set(`isPacketConn`, isPacketConn)
	f.Set(`isParseErrorPointer`, isParseErrorPointer)
	f.Set(`isResolverPointer`, isResolverPointer)
	f.Set(`isSRVPointer`, isSRVPointer)
	f.Set(`isTCPAddrPointer`, isTCPAddrPointer)
	f.Set(`isTCPConnPointer`, isTCPConnPointer)
	f.Set(`isTCPListenerPointer`, isTCPListenerPointer)
	f.Set(`isUDPAddrPointer`, isUDPAddrPointer)
	f.Set(`isUDPConnPointer`, isUDPConnPointer)
	f.Set(`isUnixAddrPointer`, isUnixAddrPointer)
	f.Set(`isUnixConnPointer`, isUnixConnPointer)
	f.Set(`isUnixListenerPointer`, isUnixListenerPointer)
	f.Set(`isUnknownNetworkError`, isUnknownNetworkError)
}

func isUnknownNetworkError(i interface{}) bool {
	_, result := i.(net.UnknownNetworkError)
	return result
}
func isParseErrorPointer(i interface{}) bool {
	_, result := i.(*net.ParseError)
	return result
}
func isResolverPointer(i interface{}) bool {
	_, result := i.(*net.Resolver)
	return result
}
func isSRVPointer(i interface{}) bool {
	_, result := i.(*net.SRV)
	return result
}
func isTCPAddrPointer(i interface{}) bool {
	_, result := i.(*net.TCPAddr)
	return result
}
func isTCPConnPointer(i interface{}) bool {
	_, result := i.(*net.TCPConn)
	return result
}
func isTCPListenerPointer(i interface{}) bool {
	_, result := i.(*net.TCPListener)
	return result
}
func isUDPAddrPointer(i interface{}) bool {
	_, result := i.(*net.UDPAddr)
	return result
}
func isUDPConnPointer(i interface{}) bool {
	_, result := i.(*net.UDPConn)
	return result
}
func isUnixAddrPointer(i interface{}) bool {
	_, result := i.(*net.UnixAddr)
	return result
}
func isUnixConnPointer(i interface{}) bool {
	_, result := i.(*net.UnixConn)
	return result
}
func isUnixListenerPointer(i interface{}) bool {
	_, result := i.(*net.UnixListener)
	return result
}
func isPacketConn(i interface{}) bool {
	_, result := i.(net.PacketConn)
	return result
}
func isMXPointer(i interface{}) bool {
	_, result := i.(*net.MX)
	return result
}
func isNSPointer(i interface{}) bool {
	_, result := i.(*net.NS)
	return result
}
func isOpErrorPointer(i interface{}) bool {
	_, result := i.(*net.OpError)
	return result
}
func isListener(i interface{}) bool {
	_, result := i.(net.Listener)
	return result
}
func isListenConfigPointer(i interface{}) bool {
	_, result := i.(*net.ListenConfig)
	return result
}
func isInvalidAddrError(i interface{}) bool {
	_, result := i.(net.InvalidAddrError)
	return result
}
func isIPNetPointer(i interface{}) bool {
	_, result := i.(*net.IPNet)
	return result
}
func isInterfacePointer(i interface{}) bool {
	_, result := i.(*net.Interface)
	return result
}
func isIPMask(i interface{}) bool {
	_, result := i.(net.IPMask)
	return result
}
func isIPAddrPointer(i interface{}) bool {
	_, result := i.(*net.IPAddr)
	return result
}
func isIPConnPointer(i interface{}) bool {
	_, result := i.(*net.IPConn)
	return result
}
func isError(i interface{}) bool {
	_, result := i.(net.Error)
	return result
}
func isFlags(i interface{}) bool {
	_, result := i.(net.Flags)
	return result
}
func isHardwareAddr(i interface{}) bool {
	_, result := i.(net.HardwareAddr)
	return result
}
func isIP(i interface{}) bool {
	_, result := i.(net.IP)
	return result
}
func isDNSConfigErrorPointer(i interface{}) bool {
	_, result := i.(*net.DNSConfigError)
	return result
}
func isDNSErrorPointer(i interface{}) bool {
	_, result := i.(*net.DNSError)
	return result
}
func isDialerPointer(i interface{}) bool {
	_, result := i.(*net.Dialer)
	return result
}
func isAddr(i interface{}) bool {
	_, result := i.(net.Addr)
	return result
}
func isAddrError(i interface{}) bool {
	_, result := i.(net.AddrError)
	return result
}
func isBuffers(i interface{}) bool {
	_, result := i.(net.Buffers)
	return result
}
func isConn(i interface{}) bool {
	_, result := i.(net.Conn)
	return result
}
func Resolver() *net.Resolver {
	return &net.Resolver{}
}
func (f *factory) getFlagUp(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.FlagUp)
}
func (f *factory) getFlagBroadcast(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.FlagBroadcast)
}
func (f *factory) getFlagLoopback(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.FlagLoopback)
}
func (f *factory) getFlagPointToPoint(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.FlagPointToPoint)
}
func (f *factory) getFlagMulticast(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.FlagMulticast)
}
func Flags(v uint) net.Flags {
	return net.Flags(v)
}
func Dialer() *net.Dialer {
	return &net.Dialer{}
}
func (f *factory) getDefaultResolver(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.DefaultResolver)
}
func (f *factory) getErrClosed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.ErrClosed)
}
func (f *factory) getErrWriteToConnected(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.ErrWriteToConnected)
}
func (f *factory) getIPv6zero(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6zero)
}
func (f *factory) getIPv6unspecified(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6unspecified)
}
func (f *factory) getIPv6loopback(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6loopback)
}
func (f *factory) getIPv6interfacelocalallnodes(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6interfacelocalallnodes)
}
func (f *factory) getIPv6linklocalallnodes(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6linklocalallnodes)
}
func (f *factory) getIPv6linklocalallrouters(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6linklocalallrouters)
}
func (f *factory) getIPv4bcast(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv4bcast)
}
func (f *factory) getIPv4allsys(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv4allsys)
}
func (f *factory) getIPv4allrouter(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv4allrouter)
}
func (f *factory) getIPv4zero(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv4zero)
}
func (f *factory) getIPv4len(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv4len)
}
func (f *factory) getIPv6len(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(net.IPv6len)
}
