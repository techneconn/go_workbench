package concurrency

// ★ By using context package, this also can be achieved
// Combine any number of channels together into a single channel
// that will close as soon as any of its component channels are closed, or written to.
func orChannel(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	// termination criteria
	case 0:
		return nil
	// if variadic sile only contains one element, we just return that
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	// wait for messages on our channels without blocking
	go func() {
		defer close(orDone)

		switch len(channels) {
		// because of how we're recursing, every recursive call to or will at least
		// have two channels. As an optimization to keep the number of goroutines constrained,
		// we place a special case here for calls to or with only two channels
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		// recursively create an or-channel from all the channels in our slice after the 3rd index
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			// orDone is also passed to exit child trees when parent tree is done
			case <-orChannel(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
