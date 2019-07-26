package p4_config_v1

func CreateTable() Table {
	return Table{}
}
func (t *Table) AddPreamble(preamble *Preamble) {
	t.Preamble = preamble
}
func (t *Table) AddMatchField(field *MatchField) {
	t.MatchFields = append(t.MatchFields, field)
}

func (t *Table) AddActionRef(ref *ActionRef) {
	t.ActionRefs = append(t.ActionRefs, ref)
}

func CreatePreamble(id uint32, name string, alias string) *Preamble {
	return &Preamble{
		Id:    id,
		Name:  name,
		Alias: alias,
	}
}

func CreateActionRef(id uint32, scope ActionRef_Scope, annotations []string) *ActionRef {
	return &ActionRef{
		Id:          id,
		Scope:       scope,
		Annotations: annotations,
	}
}

func MatchType(matchType MatchField_MatchType) *MatchField_MatchType_ {
	return &MatchField_MatchType_{
		MatchType: matchType,
	}
}

func CreateTypeName(name string) *P4NamedType {
	return &P4NamedType{
		Name: name,
	}
}

func CreateMatchField(
	id uint32,
	name string,
	annotations []string,
	bitwidth int32,
	match MatchField_MatchType,
) *MatchField {
	return &MatchField{
		Id:          id,
		Name:        name,
		Annotations: annotations,
		Bitwidth:    bitwidth,
		Match: &MatchField_MatchType_{
			MatchType: match,
		},
	}
}
