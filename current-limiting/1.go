/**
 * User: Trunks(Gao)
 * Date: 2020/06/10
 * goroutine协程
 */

package current_limiting

import (
	"fmt"
	"sync"
	"time"
)

// 限流器定义
type RequestLimitService struct {
	Interval time.Duration // 设置时间窗口大小
	MaxCount int           // 窗口内能支持的最大请求数（阈值）
	Lock     sync.Mutex    // 并发控制锁
	ReqCount int           // 当前窗口请求数（计数器）
}

// 判断当前窗口请求数是否大于最大请求数
func (reqLimit *RequestLimitService) IsAvailable() bool {
	reqLimit.Lock.Lock()
	defer reqLimit.Lock.Unlock()

	return reqLimit.ReqCount < reqLimit.MaxCount
}

// 对当前窗口请求数 +1
func (reqLimit *RequestLimitService) Increase() {
	reqLimit.Lock.Lock()
	defer reqLimit.Lock.Unlock()

	reqLimit.ReqCount += 1
}

func NewRequestLimitService(interval time.Duration, maxCnt int) *RequestLimitService {
	reqLimit := &RequestLimitService{
		Interval: interval,
		MaxCount: maxCnt,
	}

	go func() {
		ticker := time.NewTicker(interval) // 当达到窗口时间，将计数器清零
		for {
			<-ticker.C
			reqLimit.Lock.Lock()
			fmt.Println("Reset Count...")
			reqLimit.ReqCount = 0
			reqLimit.Lock.Unlock()
		}
	}()

	return reqLimit
}
