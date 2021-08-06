package demo_etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

// Watch A watch only tells the latest revision
type Watch struct {
	revision      int64                // 保存最新的 revision 号
	cancel        context.CancelFunc   // 控制 watcher 退出
	eventChan     chan *clientv3.Event // 返回给上层的数据 channel
	eventChanSize int
	lock          *sync.RWMutex

	incipientKVs []*mvccpb.KeyValue // 保存了目前 prefix 下的所有值
}

type CustomClient struct {
	clientv3.Client
}

func (client *CustomClient) WatchPrefix(ctx context.Context, prefix string) (*Watch, error) {
	// 初始化请求 WithPrefix
	resp, err := client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	// 初始化 Watch 结构
	var w = &Watch{
		eventChanSize: 64,
		revision:      resp.Header.Revision,
		eventChan:     make(chan *clientv3.Event, 64),
		incipientKVs:  resp.Kvs,
	}

	go func() {
		ctx, cancel := context.WithCancel(context.Background())

		// 注意：给外部的 cancel 方法，用于取消下面的 watch
		w.cancel = cancel

		// 注意，client.Watch 是一个子协程
		rch := client.Client.Watch(ctx, prefix, clientv3.WithPrefix(), clientv3.WithCreatedNotify(), clientv3.WithRev(w.revision))
		for {
			for n := range rch {
				// 一般情况下，协程的逻辑会阻塞在此
				if n.CompactRevision > w.revision {
					w.revision = n.CompactRevision
				}
				// 是否需要更新当前的最新的 revision
				if n.Header.GetRevision() > w.revision {
					w.revision = n.Header.GetRevision()
				}
				if err := n.Err(); err != nil {
					fmt.Println(err)
					continue
				}
				for _, ev := range n.Events {
					select {
					// 将事件 event 通过 eventChan 通知上层
					case w.eventChan <- ev:
					default:
						fmt.Println("watch etcd with prefix, block event chan drop event message")
					}
				}
			}
			// 当 watch() 被上层取消时，逻辑会走到此
			ctx, cancel := context.WithCancel(context.Background())
			w.cancel = cancel
			if w.revision > 0 {
				// 如果 revision 非 0，那么使用 WithRev 从 revision 的位置开始监听好了
				rch = client.Watch(ctx, prefix, clientv3.WithPrefix(), clientv3.WithCreatedNotify(), clientv3.WithRev(w.revision))
			} else {
				rch = client.Watch(ctx, prefix, clientv3.WithPrefix(), clientv3.WithCreatedNotify())
			}
		}
	}()

	// 返回 w（控制 channel 和数据 channel 给上层应用）
	return w, nil
}
