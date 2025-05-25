package core

// OrderRepository ต้องมี method ชื่อ Save ที่รับพารามิเตอร์เป็น Order และ return error
type OrderRepository interface {
	Save(order Order) error //รับ Order และ return error
}
