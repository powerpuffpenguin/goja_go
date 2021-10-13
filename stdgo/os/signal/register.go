package signal

import "os/signal"

func (f *factory) register() {
	f.Set(`Ignore`, signal.Ignore)
	f.Set(`Ignored`, signal.Ignored)
	f.Set(`Notify`, signal.Notify)
	f.Set(`NotifyContext`, signal.NotifyContext)
	f.Set(`Reset`, signal.Reset)
	f.Set(`Stop`, signal.Stop)
}
