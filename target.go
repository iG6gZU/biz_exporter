package main

import (
	"encoding/json"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type Target struct {
	mes      string
	topic    string
	cluster  string
	biz_name string
	success  bool
	timeout  bool
	fail     bool
	code     string
	err_mes  string
}

type rsp_param_jingangqu struct {
	Code     string `json:"code"`
	ErrCode  string `json:"errCode"`
	ErrorMsg string `json:"errorMsg"`
	Message  string `json:"message"`
}

type req_param_jingangqu struct {
	Requestlocation string `json:"Requestlocation"`
	IsShowTab       string `json:"isShowTab"`
}

type mes_jingangqu struct {
	CostTime     float64 `json:"costTime"`
	RequestParam string  `json:"requestParam"`
	ResponseParm string  `json:"responseParm"`
}

type mes_zsyh struct {
	Requestlocation string
	Cost_time       float64
	ResPonsecode    string
	Response_param  rsp_mes_zsyh
}

type rsp_mes_zsyh struct {
	Code string
	Msg  string
}

type mes_number struct {
	Cost_time      float64
	ResponseCode   string
	Response_param rsp_mes_number
}

type rsp_mes_number struct {
	RspCode string
	RspDesc string
}

type mes_double11 struct {
	Request_type   string
	Cost_time      float64
	Response_param string
}

type mes_pagebroad struct {
	Requestlocation string
	Cost_time       float64
	ResponseCode    string
	Response_param  rsp_mes_pagebroad
}

type rsp_mes_pagebroad struct {
	Code string
	Msg  string
}

type param_double11_result struct {
	Code         string
	Msg          string
	UNI_BSS_HEAD param_double11_result_nlpt
}

type param_double11_result_nlpt struct {
	RESP_CODE string
	RESP_DESC string
}

type param_double11 struct {
	RspResult param_double11_result
}

func (t *Target) target(counters []*prometheus.CounterVec) {
	switch t.cluster {
	case "JINGANGQU_LOG":
		switch t.topic {
		case "xxx":
			var m mes_jingangqu
			json.Unmarshal([]byte(t.mes), &m)
			var req req_param_jingangqu
			var rsp rsp_param_jingangqu
			req_param := strings.Replace(m.RequestParam, "\\", "", -1)
			rsp_param := strings.Replace(m.ResponseParm, "\\", "", -1)
			json.Unmarshal([]byte(req_param), &req)
			json.Unmarshal([]byte(rsp_param), &rsp)
			switch req.Requestlocation {
			case "1":
				t.biz_name = "手机数码-百亿补贴"
				if m.CostTime > 1000 {
					t.timeout = true
				} else {
					t.code = rsp.ErrCode
					if t.code == "1" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = rsp.ErrorMsg
					}
				}
			case "2":
				t.biz_name = "商城特惠购机"
				if m.CostTime > 1000 {
					t.timeout = true
				} else {
					t.code = rsp.ErrCode
					if t.code == "1" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = rsp.ErrorMsg
					}
				}

			case "5":
				if req.IsShowTab == "2" {
					t.biz_name = "智慧双推-调用资源中心"
				} else {
					t.biz_name = "智慧双推-不调用资源中心"
				}
				if m.CostTime > 1000 {
					t.timeout = true
				} else {
					t.code = rsp.Code
					if t.code == "1" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = rsp.Message
					}
				}
			case "6":
				t.biz_name = "五入口实时排序"
				if m.CostTime > 1000 {
					t.timeout = true
				} else {
					t.success = true
				}

			default:
				t.biz_name = req.Requestlocation
			}
		default:
			return
		}

	case "zntj-new":
		switch t.topic {
		case "APP_ZSYH_LOG":
			var m mes_zsyh
			json.Unmarshal([]byte(t.mes), &m)
			switch m.Requestlocation {
			case "1":
				t.biz_name = "专属优惠接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = m.Response_param.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = m.Response_param.Msg
					}
				}
			case "2":
				t.biz_name = "瀑布流接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = m.Response_param.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = m.Response_param.Msg
					}
				}
			case "3":
				t.biz_name = "靓号刷新接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = m.Response_param.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = m.Response_param.Msg
					}
				}
			case "4":
				t.biz_name = "异网专属推荐靓号接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = m.Response_param.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = m.Response_param.Msg
					}
				}
			default:
				t.biz_name = m.Requestlocation
			}
		case "APP_BENUMBER_LOG":
			var m mes_number
			json.Unmarshal([]byte(t.mes), &m)
			t.biz_name = "靓号获取接口"
			if m.Cost_time > 1000 {
				t.timeout = true
			} else {
				t.code = m.Response_param.RspCode
				if t.code == "0000" {
					t.success = true
				} else {
					t.fail = true
					t.err_mes = m.Response_param.RspDesc
				}
			}
		case "APP_DOUBLE11_LOG":
			var m mes_double11
			json.Unmarshal([]byte(t.mes), &m)
			var p param_double11
			rsp_param := strings.Replace(m.Response_param, "\\", "", -1)
			json.Unmarshal([]byte(rsp_param), &p)

			switch m.Request_type {
			case "3":
				t.biz_name = "宽带tab页获取宽带接口"

				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}

			case "4":
				t.biz_name = "专属优惠宽带已推荐接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "5":
				t.biz_name = "手机tab合约机推荐接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "6":
				t.biz_name = "宽带tab页瀑布流接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "7":
				t.biz_name = "我的页面-我的设备接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "10":
				t.biz_name = "家庭tab页瀑布流推荐"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "12":
				t.biz_name = "手机详情页为你推荐"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "13":
				t.biz_name = "合约续约下单"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "18":
				t.biz_name = "终端页面搜索模块接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "19":
				t.biz_name = "手机tab页-瀑布流"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "50_1":
				t.biz_name = "618-首页弹窗"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "50_2":
				t.biz_name = "618-Top5弹窗"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "50_3":
				t.biz_name = "618-主活动主图"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "50_4":
				t.biz_name = "618-商城百亿补贴"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "51":
				t.biz_name = "618-更多机型"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "52":
				t.biz_name = "618-更多机型-查询政策信息"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "53":
				t.biz_name = "618-新品楼层"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "54":
				t.biz_name = "618-捡便宜楼层"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "55":
				t.biz_name = "618-挽留弹窗"
				if m.Cost_time > 1900 {
					t.timeout = true
				} else {
					t.code = p.RspResult.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = p.RspResult.Msg
					}
				}
			case "56":
				t.biz_name = "618-能开商城接口"
				if m.Cost_time > 1600 {
					t.timeout = true
				} else {
					t.code = p.RspResult.UNI_BSS_HEAD.RESP_CODE
					if t.code == "00000" {
						t.success = true
					} else {
						t.fail = true
						if strings.Contains(p.RspResult.UNI_BSS_HEAD.RESP_CODE, "调用组合逻辑出错") {
							t.err_mes = "调用组合逻辑出错"
						} else {
							t.err_mes = p.RspResult.UNI_BSS_HEAD.RESP_DESC
						}
					}
				}
			default:
				t.biz_name = m.Request_type
			}
		case "MY_PAGEBROAD_LOG":
			var m mes_pagebroad
			json.Unmarshal([]byte(t.mes), &m)
			switch m.Requestlocation {
			case "1":
				t.biz_name = "我的页面-小黑条宽带推荐接口"
				if m.Cost_time > 1000 {
					t.timeout = true
				} else {
					t.code = m.Response_param.Code
					if t.code == "0" {
						t.success = true
					} else {
						t.fail = true
						t.err_mes = m.Response_param.Msg
					}
				}
			default:
				t.biz_name = m.Requestlocation
			}
		default:
			return
		}
	}

	counters[0].WithLabelValues(t.topic, t.biz_name).Inc()

	if t.timeout == true {
		counters[2].WithLabelValues(t.topic, t.biz_name).Inc()
	}

	if t.success == true {
		counters[1].WithLabelValues(t.topic, t.biz_name).Inc()
	}

	if t.fail == true {
		counters[3].WithLabelValues(t.topic, t.biz_name, t.code, t.err_mes).Inc()
	}

}
