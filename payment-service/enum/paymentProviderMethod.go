package enum

type ProviderMethod string

const (
	ProviderMethodBank           ProviderMethod = "BANK"
	ProviderMethodCryptoCurrency ProviderMethod = "CRYPTO_CURRENCY"
	ProviderMethodDebit          ProviderMethod = "DEBIT"
	ProviderMethodCredit         ProviderMethod = "CREDIT"
	ProviderMethodCashOnDelivery ProviderMethod = "CASH_ON_DELIVERY"
)
