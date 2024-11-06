package product

import "errors"

var ErrMustHavePrimaryProduct = errors.New("at least 1 must have a primary product")
var ErrMustHavePrimaryMedia = errors.New("at least 1 must have a primary media")
var ErrOnlyChooseOnePrimaryProduct = errors.New("can only choose primary product 1")
var ErrOnlyChooseOnePrimaryMedia = errors.New("can only choose primary media 1")
var ErrInvalidSubCategoryItem = errors.New("invalid subcategory item")
var ErrMustHaveSizeGuide = errors.New("must have a size guide")
var ErrVariantValue1IsRequired = errors.New("variant value 1 is required")
var ErrVariantValue2IsRequired = errors.New("variant value 2 is required")
var ErrOutletNotFound = errors.New("outlet not found")
