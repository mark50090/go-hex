package core

// 1.สร้าง struct ของ data
type Order struct {
	Order_name string  `bson:"order_name" json:"order_name"` // MongoDB จะใช้ ObjectID เป็น primary key
	Total      float64 `bson:"total" json:"total"`
}
