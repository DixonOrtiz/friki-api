package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
)

var (
	testUserID     = 12345
	testExternalID = "12345"
)

var testSuperAdminUser = entities.User{
	ID:         1,
	RoleID:     consts.SUPER_ADMIN,
	Name:       "Charles Montgomery",
	LastName:   "Burns",
	Email:      "mr.burns@gmail.com",
	Cellphone:  "+56913207055",
	Rut:        "16927586-6",
	ExternalID: "SPR-USA-476",
	Picture:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ6e2rfWamM2jg4uhYv14cnr9TMoTCMnVgmXNkbAJl-qw&s",
}

var testNormalUser = entities.User{
	ID:         1,
	RoleID:     consts.USER,
	Name:       "Homer Jay",
	LastName:   "Simpson",
	Email:      "mr.x@gmail.com",
	Cellphone:  "+569131207055",
	Rut:        "11247566-2",
	ExternalID: "SPR-USA-476",
	Picture:    "https://imagenes.elpais.com/resizer/EL70gbJpew8S1LX26VnuikTEEL8=/1960x0/arc-anglerfish-eu-central-1-prod-prisa.s3.amazonaws.com/public/CKJODPGW4OA2KDYDF6SSLHZIJE.jpg",
}
