package base

import "sync"

type umapface interface {
	Load(k string) interface{}
	Store(k string, v interface{})
	Delete(k string)
	LoadAndDelete(k string) interface{}
	LoadOrStore(k string, v interface{}) interface{}
}

// 第一种实现
type umap struct {
	m map[string]interface{}
	sync.RWMutex
}

func NewUmap() umapface {
	return &umap{
		m: make(map[string]interface{}),
	}
}

func (u *umap) Load(k string) interface{} {
	u.RLock()
	defer u.RUnlock()
	return u.m[k]
}

func (u *umap) Store(k string, v interface{}) {
	u.Lock()
	defer u.Unlock()
	u.m[k] = v
}

func (u *umap) Delete(k string) {
	u.Lock()
	defer u.Unlock()
	delete(u.m, k)
}

func (u *umap) LoadAndDelete(k string) interface{} {
	u.Lock()
	defer u.Unlock()
	v := u.m[k]
	delete(u.m, k)
	return v
}

func (u *umap) LoadOrStore(k string, v interface{}) interface{} {
	u.Lock()
	defer u.Unlock()

	ok, val := u.m[k]
	if ok == nil {
		u.m[k] = v
		return v
	}
	return val
}

type concurrentMapShared struct {
	items map[string]interface{}
	sync.RWMutex
}

type cmap []*concurrentMapShared

const sharedCount = 32

func NewCMap() umapface {
	m := make(cmap, sharedCount)
	for i := 0; i < sharedCount; i++ {
		m[i] = &concurrentMapShared{
			items: make(map[string]interface{}),
		}
	}
	return m
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func (c cmap) getShared(k string) *concurrentMapShared {
	h := fnv32(k)
	index := uint(h) % uint(sharedCount)
	return c[index]
}

func (c cmap) Load(k string) interface{} {
	cm := c.getShared(k)
	cm.RLock()
	defer cm.RUnlock()
	return cm.items[k]
}

func (c cmap) Store(k string, v interface{}) {
	cm := c.getShared(k)
	cm.Lock()
	defer cm.Unlock()
	cm.items[k] = v
}

func (c cmap) Delete(k string) {
	cm := c.getShared(k)
	cm.Lock()
	defer cm.Unlock()
	delete(cm.items, k)
}

func (c cmap) LoadAndDelete(k string) interface{} {
	cm := c.getShared(k)
	cm.Lock()
	defer cm.Unlock()
	ok, v := cm.items[k]
	if ok != nil {
		delete(cm.items, k)
	}
	return v
}

func (c cmap) LoadOrStore(k string, v interface{}) interface{} {
	cm := c.getShared(k)
	cm.Lock()
	defer cm.Unlock()
	ok, val := cm.items[k]
	if ok == nil {
		cm.items[k] = v
		return v
	}
	return val
}
