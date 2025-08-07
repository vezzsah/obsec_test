package nistlayer

import (
	"slices"
)

func ValidateIfCPEExists(cpe_to_check string) bool {

	existingCpes := []string{
		"cpe:2.3:a:mercadolibre:mercadolibre:3.8.7:*:*:*:*:android:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.7.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:4.1.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.4.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:3.0.17:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.6:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.4:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.3:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.9:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.8:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.7:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.5:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.2:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.7.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.6.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.5.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.4.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:-:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.3.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.3.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.2.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.1.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.0.2:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.0.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:6.0.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.8.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.7.6:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.7.5:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.7.4:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.7.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.6.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.5.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.4.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.4.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.3.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.1.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.0.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.0.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:4.6.4:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:4.6.3:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:4.6.2:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:4.6.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:4.2.2:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.3.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.2.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.2.0:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.6.1:*:*:*:*:wordpress:*:*",
		"cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:5.1.0:*:*:*:*:wordpress:*:*",
	}

	return slices.Contains(existingCpes, cpe_to_check)
}
