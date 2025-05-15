package util

import "google.golang.org/grpc/metadata"

type MetadataHeaderCarrier struct {
	metadata.MD
}

func (c *MetadataHeaderCarrier) Get(key string) string {
	values := c.MD.Get(key)
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func (c *MetadataHeaderCarrier) Set(key string, value string) {
	c.MD.Set(key, value)
}

func (c *MetadataHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(c.MD))
	for k := range c.MD {
		keys = append(keys, k)
	}
	return keys
}
