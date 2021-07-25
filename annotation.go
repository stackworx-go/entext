package entext

import "entgo.io/ent/schema"

type Annotation struct {
	// PointInTime is entity a point in time entity
	PointInTime bool `json:"pointInTime,omitempty"`
	// Audited is the entity audited
	Audited bool `json:"audited,omitempty"`
}

func (Annotation) Name() string {
	return "StackworxOSS"
}

var _ schema.Annotation = (*Annotation)(nil)

// Merge implements the schema.Merger interface.
func (a Annotation) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}
	if b := ant.PointInTime; b {
		a.PointInTime = true
	}
	if b := ant.Audited; b {
		a.Audited = true
	}
	return a
}
