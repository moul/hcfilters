package hcfilters

// MaxSize filters file that are too big to be cached
func MaxSize(next Cache, limit int) Cache {
	return maxSizeFilter{
		next:  next,
		limit: limit,
	}
}

type maxSizeFilter struct {
	next  Cache
	limit int
}

func (f maxSizeFilter) Get(key string) (responseBytes []byte, ok bool) {
	return f.next.Get(key)
}
func (f maxSizeFilter) Delete(key string) {
	f.next.Delete(key)
}
func (f maxSizeFilter) Set(key string, responseBytes []byte) {
	if len(responseBytes) < f.limit {
		f.next.Set(key, responseBytes)
	}
}

// MinSize filters file that are too small to be cached
func MinSize(next Cache, limit int) Cache {
	return minSizeFilter{
		next:  next,
		limit: limit,
	}
}

type minSizeFilter struct {
	next  Cache
	limit int
}

func (f minSizeFilter) Get(key string) (responseBytes []byte, ok bool) {
	return f.next.Get(key)
}
func (f minSizeFilter) Delete(key string) {
	f.next.Delete(key)
}
func (f minSizeFilter) Set(key string, responseBytes []byte) {
	if len(responseBytes) > f.limit {
		f.next.Set(key, responseBytes)
	}
}

// Cache is a clone of the httpcache.Cache struct to avoid having a dependency on the original repo
type Cache interface {
	// Get returns the []byte representation of a cached response and a bool
	// set to true if the value isn't empty
	Get(key string) (responseBytes []byte, ok bool)
	// Set stores the []byte representation of a response against a key
	Set(key string, responseBytes []byte)
	// Delete removes the value associated with the key
	Delete(key string)
}
