package smtp

import "net/smtp"

func (f *factory) register() {
	f.Set(`SendMail`, smtp.SendMail)

	f.Set(`CRAMMD5Auth`, smtp.CRAMMD5Auth)
	f.Set(`PlainAuth`, smtp.PlainAuth)
	f.Set(`Dial`, smtp.Dial)
	f.Set(`NewClient`, smtp.NewClient)
	f.Set(`isAuth`, isAuth)
	f.Set(`isClientPointer`, isClientPointer)
	f.Set(`isServerInfoPointer`, isServerInfoPointer)
}
func isAuth(i interface{}) bool {
	_, result := i.(smtp.Auth)
	return result
}
func isClientPointer(i interface{}) bool {
	_, result := i.(*smtp.Client)
	return result
}
func isServerInfoPointer(i interface{}) bool {
	_, result := i.(*smtp.ServerInfo)
	return result
}
