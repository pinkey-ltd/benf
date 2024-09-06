package ven_service

type cesiumWidgetOption struct {
	fullscreenButton     bool // 全屏按钮
	homeButton           bool // 主页按钮，回默认视口
	infoBox              bool // 信息框
	navigationHelpButton bool // 帮助按钮
	sceneModePicker      bool // 场景选择器，2D、3D、投影
	timeline             bool // 时间轴
	animation            bool // 动画控件，左下角的码表
	geocoder             bool // 地理编码器，搜索框
	baseLayerPicker      bool // 底图选择器
	vrButton             bool // VR按钮
}

type cesiumOption struct {
	widget cesiumWidgetOption
}

type Option struct {
	cesium cesiumOption
}
