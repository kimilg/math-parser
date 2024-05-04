package math

type Element struct {
	Classifications []Classification `json:"classifications" gorm:"many2many:element_classifications;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
}

