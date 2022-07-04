package main

import (
	"fmt"

	"github.com/gogf/gf/net/ghttp"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	client := ghttp.NewClient()
	client.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
	// client.SetTimeOut(time.Second * 10)

	url := "https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=35&provinceId=0&pageSize=50&isVerify=1&pageNo=1"
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer response.Body.Close()
	var auto PlsAutoGenerated
	jsoniter.NewDecoder(response.Body).Decode(&auto)
	println(len(auto.Value.List))
}

type PlsAutoGenerated struct {
	DataFrom     string `json:"dataFrom"`
	EmptyFlag    bool   `json:"emptyFlag"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
	Value        struct {
		List []struct {
			DrawFlowFund            string        `json:"drawFlowFund"`
			DrawFlowFundRj          string        `json:"drawFlowFundRj"`
			EstimateDrawTime        string        `json:"estimateDrawTime"`
			IsGetKjpdf              int           `json:"isGetKjpdf"`
			IsGetXlpdf              int           `json:"isGetXlpdf"`
			LotteryDrawNum          string        `json:"lotteryDrawNum"`
			LotteryDrawResult       string        `json:"lotteryDrawResult"`
			LotteryDrawStatus       int           `json:"lotteryDrawStatus"`
			LotteryDrawTime         string        `json:"lotteryDrawTime"`
			LotteryEquipmentCount   int           `json:"lotteryEquipmentCount"`
			LotteryGameName         string        `json:"lotteryGameName"`
			LotteryGameNum          string        `json:"lotteryGameNum"`
			LotteryGamePronum       int           `json:"lotteryGamePronum"`
			LotteryPaidBeginTime    string        `json:"lotteryPaidBeginTime"`
			LotteryPaidEndTime      string        `json:"lotteryPaidEndTime"`
			LotteryPromotionFlag    int           `json:"lotteryPromotionFlag"`
			LotterySaleBeginTime    string        `json:"lotterySaleBeginTime"`
			LotterySaleEndTimeUnix  int           `json:"lotterySaleEndTimeUnix"`
			LotterySaleEndtime      string        `json:"lotterySaleEndtime"`
			LotterySuspendedFlag    int           `json:"lotterySuspendedFlag"`
			LotteryUnsortDrawresult string        `json:"lotteryUnsortDrawresult"`
			MatchList               []interface{} `json:"matchList"`
			PdfType                 int           `json:"pdfType"`
			PoolBalanceAfterdraw    string        `json:"poolBalanceAfterdraw"`
			PoolBalanceAfterdrawRj  string        `json:"poolBalanceAfterdrawRj"`
			PrizeLevelList          []struct {
				AwardType        int    `json:"awardType"`
				Group            string `json:"group"`
				LotteryCondition string `json:"lotteryCondition"`
				PrizeLevel       string `json:"prizeLevel"`
				Sort             int    `json:"sort"`
				StakeAmount      string `json:"stakeAmount"`
				StakeCount       string `json:"stakeCount"`
				TotalPrizeamount string `json:"totalPrizeamount"`
			} `json:"prizeLevelList"`
			PrizeLevelListRj  []interface{} `json:"prizeLevelListRj"`
			RuleType          int           `json:"ruleType"`
			TermList          []interface{} `json:"termList"`
			TotalSaleAmount   string        `json:"totalSaleAmount"`
			TotalSaleAmountRj string        `json:"totalSaleAmountRj"`
			Verify            int           `json:"verify"`
			VtoolsConfig      struct {
			} `json:"vtoolsConfig"`
		} `json:"list"`
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
		Pages    int `json:"pages"`
		Total    int `json:"total"`
	} `json:"value"`
}
