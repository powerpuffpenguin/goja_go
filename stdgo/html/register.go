package html

import "html"

func (f *factory) register() {
	f.Set(`EscapeString`, html.EscapeString)
	f.Set(`UnescapeString`, html.UnescapeString)
}
