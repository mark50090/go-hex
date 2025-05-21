package core

// port secondary ใครก็ตามที่ต้องการเป็น OrderRepository จะต้องมี method ที่ตรงกับที่กำหนดไว้ใน interface นี้
// Save(order Order) error เป็น method หนึ่งที่บอกว่า struct ที่ implement interface นี้ ต้องมี method ชื่อ Save ที่รับพารามิเตอร์เป็น Order และคืนค่าเป็น error
type OrderRepository interface {
	Save(order Order) error //เป็น การประกาศ "สัญญา" ว่า struct ที่ implement interface นี้ต้องมี method แบบนี้
}
