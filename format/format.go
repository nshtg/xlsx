package format

import (
	"github.com/plandem/xlsx/internal/ml"
	"reflect"
)

//DirectStyleID is alias of original ml.DirectStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type DirectStyleID = ml.DirectStyleID

//DiffStyleID is alias of original ml.DiffStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type DiffStyleID = ml.DiffStyleID

//NamedStyleID is alias of original ml.NamedStyleID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type NamedStyleID = ml.NamedStyleID

//DefaultDirectStyle is ID for any default direct style than depends on context:
// E.g. for cell it will be equal to NamedStyle 'Normal', for hyperlink - NamedStyle 'Hyperlink'
const DefaultDirectStyle = DirectStyleID(0)

//StyleFormat is objects that holds combined information about cell styling
type StyleFormat struct {
	styleInfo *ml.DiffStyle
	namedInfo *ml.NamedStyleInfo
}

type option func(o *StyleFormat)

//New creates and returns StyleFormat object with requested options
func New(options ...option) *StyleFormat {
	s := &StyleFormat{
		&ml.DiffStyle{
			NumberFormat: &ml.NumberFormat{},
			Font:         &ml.Font{},
			Fill: &ml.Fill{
				Pattern:  &ml.PatternFill{},
				Gradient: &ml.GradientFill{},
			},
			Border: &ml.Border{
				Left:       &ml.BorderSegment{},
				Right:      &ml.BorderSegment{},
				Top:        &ml.BorderSegment{},
				Bottom:     &ml.BorderSegment{},
				Diagonal:   &ml.BorderSegment{},
				Vertical:   &ml.BorderSegment{},
				Horizontal: &ml.BorderSegment{},
			},
			Alignment:  &ml.CellAlignment{},
			Protection: &ml.CellProtection{},
		},
		&ml.NamedStyleInfo{},
	}
	s.Set(options...)
	return s
}

//Set sets new options for style
func (s *StyleFormat) Set(options ...option) {
	for _, o := range options {
		o(s)
	}
}

//private method used by stylesheet manager to unpack StyleFormat
func fromStyleFormat(f *StyleFormat) (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, numFormat *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border, namedInfo *ml.NamedStyleInfo) {
	style := f.styleInfo
	named := f.namedInfo

	//copy non-empty namedInfo
	if *named != (ml.NamedStyleInfo{}) {
		namedInfo = &ml.NamedStyleInfo{}
		*namedInfo = *named
	}

	//copy non-empty alignment
	if *style.Alignment != (ml.CellAlignment{}) {
		alignment = &ml.CellAlignment{}
		*alignment = *style.Alignment
	}

	//copy non-empty font
	if (*style.Font != ml.Font{} && *style.Font != ml.Font{Size: 0, Family: 0, Charset: 0}) {
		font = &ml.Font{}
		*font = *style.Font
	}

	//copy non-empty numFormat
	if *style.NumberFormat != (ml.NumberFormat{}) {
		numFormat = &ml.NumberFormat{}
		*numFormat = *style.NumberFormat
	}

	//copy non-empty protection
	if *style.Protection != (ml.CellProtection{}) {
		protection = &ml.CellProtection{}
		*protection = *style.Protection
	}

	//copy non-empty border
	border = &ml.Border{}
	*border = *style.Border

	if reflect.DeepEqual(border.Left, &ml.BorderSegment{}) {
		border.Left = nil
	} else {
		border.Left = &ml.BorderSegment{}
		*border.Left = *style.Border.Left
	}

	if reflect.DeepEqual(border.Right, &ml.BorderSegment{}) {
		border.Right = nil
	} else {
		border.Right = &ml.BorderSegment{}
		*border.Right = *style.Border.Right
	}

	if reflect.DeepEqual(border.Top, &ml.BorderSegment{}) {
		border.Top = nil
	} else {
		border.Top = &ml.BorderSegment{}
		*border.Top = *style.Border.Top
	}

	if reflect.DeepEqual(border.Bottom, &ml.BorderSegment{}) {
		border.Bottom = nil
	} else {
		border.Bottom = &ml.BorderSegment{}
		*border.Bottom = *style.Border.Bottom
	}

	if reflect.DeepEqual(border.Diagonal, &ml.BorderSegment{}) {
		border.Diagonal = nil
	} else {
		border.Diagonal = &ml.BorderSegment{}
		*border.Diagonal = *style.Border.Diagonal
	}

	if reflect.DeepEqual(border.Vertical, &ml.BorderSegment{}) {
		border.Vertical = nil
	} else {
		border.Vertical = &ml.BorderSegment{}
		*border.Vertical = *style.Border.Vertical
	}

	if reflect.DeepEqual(border.Horizontal, &ml.BorderSegment{}) {
		border.Horizontal = nil
	} else {
		border.Horizontal = &ml.BorderSegment{}
		*border.Horizontal = *style.Border.Horizontal
	}

	//if border is actually empty, then nil it
	if *border == (ml.Border{}) {
		border = nil
	}

	//copy non-empty fill
	fill = &ml.Fill{}

	//copy pattern
	if !reflect.DeepEqual(style.Fill.Pattern, &ml.PatternFill{}) {
		fill.Pattern = &ml.PatternFill{}
		*fill.Pattern = *style.Fill.Pattern
	}

	//copy gradient
	if !reflect.DeepEqual(style.Fill.Gradient, &ml.GradientFill{}) {
		fill.Gradient = &ml.GradientFill{}
		*fill.Gradient = *style.Fill.Gradient
		copy(fill.Gradient.Stop, style.Fill.Gradient.Stop)
	}

	//if fill is actually empty, then nil it
	if *fill == (ml.Fill{}) {
		fill = nil
	}

	return
}

//private method used by to convert StyleFormat to ml.RichFont
func toRichFont(f *StyleFormat) *ml.RichFont {
	style := f.styleInfo

	//copy non-empty font
	if (*style.Font != ml.Font{} && *style.Font != ml.Font{Size: 0, Family: 0, Charset: 0}) {
		font := ml.RichFont(*style.Font)
		return &font
	}

	return nil
}
