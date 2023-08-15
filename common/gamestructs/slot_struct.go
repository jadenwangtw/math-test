package gamestructs

import ()

type (
	// SpinResult
	SpinResult struct {
		TotalBet     int64
		TotalWin     int64
		Seq          int64
		StateResults []*StateResult
	}

	// StateResult
	StateResult struct {
		State        string
		SreenResults []*ScreenResult
		IsTrigger    bool
		Level        int64
	}

	//單一畫面資料，宣告定義用，若無則不會產生。
	ScreenResult struct {
		Anchors         []int64
		Symbols         [][]int64
		TotalWin        int64
		WinResults      []*WinResult
		ExtSreenResults []*ScreenResult
		IsTrigger       bool

		Round      int64
		AddRound   int64
		TotalRound int64
		Level      int64
		Multiplier int64
	}

	//單一贏分資料
	WinResult struct {
		Type   string // 贏分類型
		ID     int64  // 贏分 ID
		Symbol int64  // 贏分圖標
		Count  int64  // 贏分圖標個數
		Hit    int64
		Win    int64
		Msg    string
	}
)

var (
	// state type
	StateTypeBase = "BASE"
	StateTypeFeature1 = "FEATURE_1"
	StateTypeFeature2 = "FEATURE_2"
	StateTypeFeature3 = "FEATURE_3"
	StateTypeFeature4 = "FEATURE_4"
	StateTypeFeature5 = "FEATURE_5"

	// win type
	WinTypeScatter = "SCATTER"
	WinTypeLineLR = "LINE_LR"
	WinTypeLineRL = "LINE_RL"
	WinTypeWayLR = "WAY_LR"
	WinTypeWayRL = "WAY_RL"
)
