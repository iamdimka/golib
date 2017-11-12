package async

type status byte

const (
	pending status = iota
	resolved
	rejected
)

type Deferred interface {
	Resolve(interface{})
	Reject(interface{})
	Await() (interface{}, bool)
}

type dfd struct {
	await  chan struct{}
	status status
	data   interface{}
}

func NewDeferred() Deferred {
	return &dfd{
		await:  make(chan struct{}),
		status: pending,
	}
}

func (d *dfd) finish(data interface{}, status status) {
	if d.status == pending {
		d.status = status
		d.data = data
		d.await <- struct{}{}
		close(d.await)
	}
}

func (d *dfd) Resolve(data interface{}) {
	d.finish(data, resolved)
}

func (d *dfd) Reject(data interface{}) {
	d.finish(data, rejected)
}

func (d *dfd) Await() (interface{}, bool) {
	if d.status == pending {
		<-d.await
	}

	return d.data, d.status == resolved
}
