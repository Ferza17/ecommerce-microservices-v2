package usecase

import (
	"context"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

func (u *paymentUseCase) CreatePayment(ctx context.Context, requestId string, request *paymentRpc.CreatePaymentRequest) error {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.CreatePayment")
	defer span.End()

	//// Begin transaction
	//db := u.paymentRepository.OpenTransactionWithContext(ctx)
	//tx := db.Begin()
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback() // Roll back on panic
	//		panic(r)
	//	}
	//}()
	//
	//// Map the request into the ORM Payment model
	//payment := &orm.Payment{
	//	Code: util.GenerateInvoiceCode(),
	//	//TotalPrice: request.TotalPrice,
	//	Status: enum.PaymentStatusPending,
	//	//ProviderID: &request.ProviderId,
	//	UserID: request.UserId,
	//}
	//
	//// Call repository to create the Payment
	//paymentID, err := u.paymentRepository.CreatePayment(ctx, requestId, payment, tx)
	//if err != nil {
	//	tx.Rollback() // Roll back the transaction on error
	//	u.logger.Error(fmt.Sprintf("Failed to create payment, requestId: %s, error: %v", requestId, err))
	//	return fmt.Errorf("failed to create payment: %w", err)
	//}
	//
	//// Process PaymentItems
	//for _, item := range request.Items {
	//	paymentItem := &orm.PaymentItem{
	//		ID:        item.Id,
	//		PaymentID: paymentID, // Associate with the Payment
	//		ProductID: item.ProductId,
	//		Amount:    item.Amount,
	//		Qty:       item.Qty,
	//	}
	//
	//	// Call repository to create the PaymentItem
	//	if _, err := u.paymentRepository.CreatePaymentItem(ctx, paymentItem, tx); err != nil {
	//		tx.Rollback() // Roll back the transaction on error
	//		u.logger.Error(fmt.Sprintf("Failed to create payment item, requestId: %s, paymentItemId: %s, error: %v", requestId, paymentItem.ID, err))
	//		return fmt.Errorf("failed to create payment item: %w", err)
	//	}
	//}
	//
	//// Commit the transaction
	//if err := tx.Commit().Error; err != nil {
	//	u.logger.Error(fmt.Sprintf("Failed to commit transaction, requestId: %s, error: %v", requestId, err))
	//	return fmt.Errorf("failed to commit transaction: %w", err)
	//}
	//
	//// Log success
	//u.logger.Info(fmt.Sprintf("Successfully created payment and payment items, requestId: %s, paymentId: %s", requestId, paymentID))
	return nil
}
