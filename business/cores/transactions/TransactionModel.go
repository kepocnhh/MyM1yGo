package transactions

type Model interface {
    GetId() int
    GetDate() int64
    GetCount() int
}