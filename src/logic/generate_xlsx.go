package logic

import (
	"github.com/tealeg/xlsx"
	"time"
	"errors"
	"../config"
)

func generate_xlsx(data *Log_data, save_file string) (err error) {
	file := xlsx.NewFile()
	sheet, add_sheet_error := file.AddSheet("Sheet1")
	if nil != add_sheet_error {
		return errors.New("generat_xlsx add sheet error: " + add_sheet_error.Error())
	}
	head_row := sheet.AddRow()
	for _, value := range config.StdConfigManager.GetConfig().Col_info_list  {
		head_cell := head_row.AddCell()
		head_cell.SetString(value.Title)
	}
	var id int = 0
	test_state_style := xlsx.NewStyle()
	test_state_style.Fill = xlsx.Fill{PatternType:"none", FgColor:"FFFFFFFF", BgColor:"FFFF00"}
	for _, value := range data.Log_entry_array  {
		var index int = 0
		id += 1
		data_row := sheet.AddRow()
		cell_id := data_row.AddCell()
		cell_id.SetInt(id)
		style_id := xlsx.NewStyle()
		style_id.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_id.SetStyle(style_id)

		index++
		cell_revision := data_row.AddCell()
		cell_revision.SetInt(value.Revision)
		style_revision := xlsx.NewStyle()
		style_revision.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_revision.SetStyle(style_revision)

		index++
		cell_auth := data_row.AddCell()
		cell_auth.SetString(value.Author)
		style_auth := xlsx.NewStyle()
		style_auth.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_auth.SetStyle(style_auth)

		index++
		date, _ := time.ParseInLocation(time.RFC3339Nano, value.Date, time.Local)
		cell_date := data_row.AddCell()
		cell_date.SetString(date.Format("2006-01-02 15:04:05"))
		style_date := xlsx.NewStyle()
		style_date.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_date.SetStyle(style_date)

		index++
		cell_desc := data_row.AddCell()
		cell_desc.SetString(value.Msg)
		style_desc := xlsx.NewStyle()
		style_desc.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_desc.SetStyle(style_desc)

		index++
		cell_test := data_row.AddCell()
		cell_test.SetString("未测试")
		style_test := xlsx.NewStyle()
		style_test.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_test.SetStyle(style_test)

		index++
		cell_comment := data_row.AddCell()
		cell_comment.SetString("")
		style_comment := xlsx.NewStyle()
		style_comment.Fill = xlsx.Fill{PatternType:"none",
			FgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Fg_color,
			BgColor:config.StdConfigManager.GetConfig().Col_info_list[index].Bg_color}
		cell_comment.SetStyle(style_comment)
	}
	save_error := file.Save(save_file)
	if nil != save_error {
		return errors.New("generat_xlsx save error: " + save_error.Error())
	}
	return nil
}
