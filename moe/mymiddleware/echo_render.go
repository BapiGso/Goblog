package mymiddleware

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/labstack/echo/v4"
	"io"
	"text/template"
	"time"
)

type TemplateRenderWithCache struct {
	Template *template.Template //渲染模板
}

// Render todo 可以在这里检查data是否有cache，有的话用cache，没有的话添加进cache
func (t *TemplateRenderWithCache) Render(w io.Writer, name string, data any, c echo.Context) error {
	//先类型断言是不是qpu，然后检查是否有缓存，有的话渲染，没有的话defer添加进缓存

	return t.Template.ExecuteTemplate(w, name, data)
}

var cache = func() *bigcache.BigCache {
	config := bigcache.Config{
		//Shards: 分片数量(必须是2的幂),默认是1024,用于将缓存划分为多个分片,以减少锁竞争
		//LifeWindow: 缓存项的生存时间,默认是10分钟,也就是缓存项可以存活的最大时间
		//CleanWindow: 清理过期缓存项的时间间隔,默认是5分钟,如果设置为<=0则不会触发主动清理
		//MaxEntriesInWindow: 每个生命周期内的最大缓存数量,默认是10分钟内最多1000个缓存项,用于初始化内存分配
		//MaxEntrySize: 单个缓存项的最大大小(字节),默认500字节,同样用于初始化内存分配
		//Verbose: 是否打印额外的内存分配日志,默认false
		//HardMaxCacheSize: 缓存的最大内存大小限制(MB),默认8GB,如果达到限制会覆盖最老的缓存项
		//OnRemove: 当缓存项被移除时的回调函数,可以用于解包缓存值等操作
		//OnRemoveWithReason: 带有移除原因的回调函数,比OnRemove多了一个移除原因的参数
		Shards:             1024,
		LifeWindow:         24 * time.Hour,
		CleanWindow:        1 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   2048,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	}
	c, err := bigcache.New(context.Background(), config)
	if err != nil {
		panic(err)
	}

	return c
}()
