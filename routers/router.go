// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact yi.liu@esmart365.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hft_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/barcode",
			beego.NSInclude(
				&controllers.BarcodeController{},
			),
		),

		beego.NSNamespace("/cloudpay_statement_detail",
			beego.NSInclude(
				&controllers.CloudpayStatementDetailController{},
			),
		),

		beego.NSNamespace("/member",
			beego.NSInclude(
				&controllers.MemberController{},
			),
		),

		beego.NSNamespace("/member_face",
			beego.NSInclude(
				&controllers.MemberFaceController{},
			),
		),

		beego.NSNamespace("/shop_financial_order",
			beego.NSInclude(
				&controllers.ShopFinancialOrderController{},
			),
		),

		beego.NSNamespace("/shop_order",
			beego.NSInclude(
				&controllers.ShopOrderController{},
			),
		),

		beego.NSNamespace("/shop_sales_records",
			beego.NSInclude(
				&controllers.ShopSalesRecordsController{},
			),
		),

		beego.NSNamespace("/shop_sales_settings",
			beego.NSInclude(
				&controllers.ShopSalesSettingsController{},
			),
		),

		beego.NSNamespace("/shop_startup_store_costs",
			beego.NSInclude(
				&controllers.ShopStartupStoreCostsController{},
			),
		),

		beego.NSNamespace("/sys_barcode",
			beego.NSInclude(
				&controllers.SysBarcodeController{},
			),
		),

		beego.NSNamespace("/sys_barcode_new",
			beego.NSInclude(
				&controllers.SysBarcodeNewController{},
			),
		),

		beego.NSNamespace("/sys_barcode_old",
			beego.NSInclude(
				&controllers.SysBarcodeOldController{},
			),
		),

		beego.NSNamespace("/unit",
			beego.NSInclude(
				&controllers.UnitController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
