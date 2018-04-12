package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:BarcodeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"] = append(beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"] = append(beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"] = append(beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"] = append(beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"] = append(beego.GlobalControllerRouter["hft_go/controllers:CloudpayStatementDetailController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"] = append(beego.GlobalControllerRouter["hft_go/controllers:MemberFaceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopFinancialOrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopOrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesRecordsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopSalesSettingsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"] = append(beego.GlobalControllerRouter["hft_go/controllers:ShopStartupStoreCostsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeNewController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"] = append(beego.GlobalControllerRouter["hft_go/controllers:SysBarcodeOldController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:UnitController"] = append(beego.GlobalControllerRouter["hft_go/controllers:UnitController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:UnitController"] = append(beego.GlobalControllerRouter["hft_go/controllers:UnitController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:UnitController"] = append(beego.GlobalControllerRouter["hft_go/controllers:UnitController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:UnitController"] = append(beego.GlobalControllerRouter["hft_go/controllers:UnitController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hft_go/controllers:UnitController"] = append(beego.GlobalControllerRouter["hft_go/controllers:UnitController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
