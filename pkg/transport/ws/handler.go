package ws

// Handler 消息处理器
type Handler struct {
	size   int             //work池大小
	engine *Engine         //路由处理
	queue  []chan *Request //Worker负责取任务的消息队列
}

// NewHandler 实例化消息处理器
func NewHandler(size int, e *Engine) *Handler {
	return &Handler{
		engine: e,
		size:   size,
		queue:  make([]chan *Request, size),
	}
}

// Init 初始化work池
func (h *Handler) Init(len int) {
	for i := 0; i < h.size; i++ {
		//初始化
		h.queue[i] = make(chan *Request, len)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go h.startWorker(h.queue[i])
	}
}

// Execute 消息处理
func (h *Handler) Execute(r *Request) {
	h.engine.Start(r)
}

// AsyncExecute 异步消息处理
func (h *Handler) AsyncExecute(r *Request) {
	//轮询分发任务
	h.queue[r.Conn().GetID()%uint64(h.size)] <- r
}

func (h *Handler) startWorker(queue chan *Request) {
	for {
		select {
		case r := <-queue:
			h.Execute(r)
		}
	}
}
