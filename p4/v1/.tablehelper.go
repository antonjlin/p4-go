package p4_v1
//
//const ()
//
//func CreateTableEntry() TableEntry {
//	return TableEntry{}
//}
//
//func (t *TableEntry) AddMatch(id uint32, matchType isFieldMatch_FieldMatchType) *TableEntry {
//
//	match := &FieldMatch{
//		FieldId:        id,
//		FieldMatchType: matchType,
//	}
//	t.Match = append(t.Match, match)
//	return t
//}
//
//func (t *TableEntry) AddAction(action *TableAction) *TableEntry {
//	t.Action = action
//	return t
//}
//
//func (t *TableEntry) SetPriority(priority int32) *TableEntry {
//	t.Priority = priority
//	return t
//}
//
//func (t *TableEntry) SetIsDefaultAction(isDefaultAction bool) *TableEntry {
//	t.IsDefaultAction = isDefaultAction
//	return t
//}
//
//func (t *TableEntry) SetMeterConfig(Cir int64, Cburst int64, Pir int64, Pburst int64) {
//	t.MeterConfig = &MeterConfig{
//		Cir:    Cir,
//		Cburst: Cburst,
//		Pir:    Pir,
//		Pburst: Pburst,
//	}
//}
//
//func CreateEmptyAction(id uint32) *TableAction {
//	t := &TableAction{
//		Type: &TableAction_Action{
//			Action: &Action{
//				ActionId: id,
//				Params:   []*Action_Param{},
//			},
//		},
//	}
//	return t
//}
//
////TODO support for
//// func (*TableAction_Action) isTableAction_Type() {}
//// func (*TableAction_ActionProfileMemberId) isTableAction_Type() {}
//// func (*TableAction_ActionProfileGroupId) isTableAction_Type() {}
//// func (*TableAction_ActionProfileActionSet) isTableAction_Type() {}
//
//func (action *TableAction) AddParam(param *Action_Param) *TableAction {
//	action.GetAction().Params = append(action.GetAction().GetParams(), param)
//	return action
//}
//
//func CreateParam(id uint32, value []byte) *Action_Param {
//	return &Action_Param{
//		ParamId: id,
//		Value:   value,
//	}
//}
//
//func CreateLPMMatch(value []byte) *FieldMatch_Lpm {
//	return &FieldMatch_Lpm{
//		Lpm: &FieldMatch_LPM{
//			Value: value,
//		},
//	}
//}
//
//func CreateTernaryMatch(value []byte, mask []byte) *FieldMatch_Ternary_ {
//	return &FieldMatch_Ternary_{
//		Ternary: &FieldMatch_Ternary{
//			Value: value,
//			Mask:  mask,
//		},
//	}
//}
//
//func CreateRangeMatch(high []byte, low []byte) *FieldMatch_Range_ {
//	return &FieldMatch_Range_{
//		Range: &FieldMatch_Range{
//			High: high,
//			Low:  low,
//		},
//	}
//}
//
//func CreateExactMatch(value []byte) *FieldMatch_Exact_ {
//	return &FieldMatch_Exact_{
//		Exact: &FieldMatch_Exact{
//			Value: value,
//		},
//	}
//}
