package resolver

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/api-gateway/infrastructure/telemetry"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/usecase"
	cartUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/commerce/cart/usecase"
	paymentUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/payment/payment/usecase"
	paymentProviderUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/payment/provider/usecase"
	productUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/product/usecase"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUseCase             userUseCase.IUserUseCase
	ProductUseCase          productUseCase.IProductUseCase
	CartUseCase             cartUseCase.ICartUseCase
	AuthUseCase             authUseCase.IAuthUseCase
	PaymentUseCase          paymentUseCase.IPaymentUseCase
	PaymentProviderUseCase  paymentProviderUseCase.IPaymentProviderUseCase
	TelemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	Logger                  pkg.IZapLogger
}
