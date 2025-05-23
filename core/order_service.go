package core

import "errors"

// Primary port
type OrderService interface {
	CreateOrder(order Order) error //โดยจะ return entities Order ตัวเดียวกับของ repo ออกไป
}

//---------------  ส่วนของ service ---------------//

//สร้าง struct เพื่อรับ business logicore เพราะ service ไม่มี data เป็นของตัวเองเป็นตัวกลางรับลูกส่งลูก เลยต้องตัวเก็บข้อมูลเพื่อใช้ดึง data
type orderServiceImpl struct {
	repo OrderRepository //การอ้างถึง interface
}

// ทำหน้าที่เหมือน Constructor คือ รับ OrderRepository และ return OrderService
func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo} //return การ intance orderServiceImpl
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}
	// Business logic...
	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
