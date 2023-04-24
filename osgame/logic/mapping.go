package logic

import (
	model "osgame/model/osserver/proto/resdata"
)

var (
	DataBinBytesFileStruct = map[string]interface{}{
		//"ResBattlePlayerLevelDataBin.bytes":   &[]*model.ResBattlePlayerLevelDataBin{},
		//"ResChessHeroConfDataBin.bytes":       &[]*model.ResChessHeroConfDataBin{},
		//"ResEquipDataBin.bytes":               &[]*model.ResEquipDataBin{},
		//"ResFetterConfDataBin.bytes":          &[]*model.ResFetterConfDataBin{},
		//"ResHeroPropertyIconCfgDataBin.bytes": &[]*model.ResHeroPropertyIconCfgDataBin{},
		"ResLordSkillDataBin.bytes": &[]*model.ResLordSkillDataBin{}, //问题线上的
		//"ResLordSkillLevelDataBin.bytes":      &[]*model.ResLordSkillLevelDataBin{},
		//"ResSkillCfgInfo.bytes":               &[]*model.ResSkillCfgInfo{}, //有问题
		//"ResSkillCombineCfgInfo.bytes":        &[]*model.ResSkillCombineCfgInfo{},
	}
)

var (
	DataBinBytesFileStructBasic = map[string]interface{}{
		//"ResLordSkinDataBin.bytes":     &[]*model.ResLordSkinDataBin{},
		//"ResSeasonLineupDataBin.bytes": &[]*model.ResSeasonLineupDataBin{},
	}
)
