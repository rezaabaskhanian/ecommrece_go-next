package checkoutservcie

import (
	"context"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) Checkout(userID int) (entity.OrderWithItem, error) {
	const op = "checkout.Usecase.Checkout"

	ctx := context.Background()

	// 1️⃣ گرفتن cart
	cart, err := s.cartRepo.GetOrCreateCart(userID)
	if err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در دریافت سبد خرید")
	}

	cartItems, err := s.cartRepo.GetCartItems(cart.ID)
	if err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در دریافت آیتم‌های سبد")
	}

	if len(cartItems) == 0 {
		return entity.OrderWithItem{},
			richerror.New(op).WithMessage("سبد خرید خالی است")
	}

	// 2️⃣ ساخت OrderItem ها + محاسبه total
	var (
		total      int64
		orderItems []entity.OrderItem
	)

	for _, item := range cartItems {

		product, err := s.productRepo.GetProductWithID(item.ProductID)
		if err != nil {
			return entity.OrderWithItem{},
				richerror.New(op).WithErr(err).WithMessage("محصول یافت نشد")
		}

		if product.Stock < item.Quantity {
			return entity.OrderWithItem{},
				richerror.New(op).WithMessage("موجودی کافی نیست")
		}

		total += product.Price * int64(item.Quantity)

		orderItems = append(orderItems, entity.OrderItem{
			ProductID: product.ID,
			Name:      product.Name,
			Price:     product.Price, // snapshot
			Quantity:  item.Quantity,
		})
	}

	// 3️⃣ ساخت Order
	order := entity.Order{
		USerID: userID,
		CartID: cart.ID,
		Status: "pending_payment",
		Total:  total,
	}

	// 4️⃣ شروع Transaction
	tx, err := s.beginRepo.BeginTx(ctx)
	if err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در شروع transaction")
	}
	defer tx.Rollback(ctx)

	// 5️⃣ ذخیره Order
	orderID, err := s.orderRepo.CreateOrder(ctx, tx, order)
	if err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در ایجاد سفارش")
	}

	// set OrderID
	for i := range orderItems {
		orderItems[i].OrderID = orderID
	}

	// 6️⃣ کاهش موجودی
	for _, item := range orderItems {
		err = s.productRepo.DecreaseStock(ctx, tx, item.ProductID, item.Quantity)
		if err != nil {
			return entity.OrderWithItem{},
				richerror.New(op).WithErr(err).WithMessage("موجودی کافی نیست")
		}
	}

	// 7️⃣ ذخیره OrderItems
	err = s.orderRepo.CreateOrderItems(ctx, tx, orderItems)
	if err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در ثبت آیتم‌های سفارش")
	}

	// 8️⃣ پاک کردن cart
	err = s.cartRepo.ClearCart(ctx, tx, cart.ID)
	if err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در پاک کردن سبد خرید")
	}

	// 9️⃣ Commit
	if err := tx.Commit(ctx); err != nil {
		return entity.OrderWithItem{},
			richerror.New(op).WithErr(err).WithMessage("خطا در commit transaction")
	}

	// 10️⃣ response
	return entity.OrderWithItem{
		Order: entity.Order{
			ID:     orderID,
			USerID: userID,
			CartID: cart.ID,
			Status: "pending_payment",
			Total:  total,
		},
		Items: orderItems,
	}, nil
}
