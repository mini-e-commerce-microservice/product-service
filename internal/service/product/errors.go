package product

import "errors"

var ErrInvalidProductCondition = errors.New("invalid product condition")
var ErrMustHavePrimaryProduct = errors.New("at least 1 must have a primary product")
var ErrMustHavePrimaryMedia = errors.New("at least 1 must have a primary media")
var ErrOnlyChooseOnePrimaryProduct = errors.New("can only choose primary product 1")
var ErrOnlyChooseOnePrimaryMedia = errors.New("can only choose primary media 1")
var ErrInvalidSubCategoryItem = errors.New("invalid subcategory item")
var ErrMustHaveSizeGuide = errors.New("must have a size guide")
var ErrVariantValueCannotBeEmpty = errors.New("variant value cannot be empty")
