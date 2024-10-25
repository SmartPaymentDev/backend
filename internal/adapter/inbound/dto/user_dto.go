package dto

import "smart-payment-dev-be/internal/core/domain"

type UserDTO struct{}

type User struct {
	CUSTID     int    `json:"custid"`
	NOCUST     string `json:"nocust"`
	NMCUST     string `json:"nmcust"`
	NUM2ND     string `json:"num_2_nd"`
	STCUST     string `json:"stcust"`
	CODE01     string `json:"code_01"`
	DESC01     string `json:"desc_01"`
	CODE02     string `json:"code_02"`
	DESC02     string `json:"desc_02"`
	CODE03     string `json:"code_03"`
	DESC03     string `json:"desc_03"`
	CODE04     string `json:"code_04"`
	DESC04     string `json:"desc_04"`
	CODE05     string `json:"code_05"`
	DESC05     string `json:"desc_05"`
	TOTPAY     string `json:"totalpay"`
	GENUS      string `json:"genus"`
	LastUpdate string `json:"last_update"`
	NO_WA      string `json:"noWa"`
}

type LoginUserRequest struct {
	NOCUST         string `json:"nocust"`
	MOBILEPASSWORD string `json:"mobilepassword"`
}

type UserResponse struct {
	JwtToken string `json:"jwt_token"`
}

type UserResponseByMe struct {
	User User `json:"user"`
}

func (ud *UserDTO) LoginUserTransformIn(req LoginUserRequest) domain.User {
	return domain.User{
		NOCUST:         req.NOCUST,
		MOBILEPASSWORD: req.MOBILEPASSWORD,
	}
}

func (ud *UserDTO) LoginUserTransformOut(token string) UserResponse {
	return UserResponse{
		JwtToken: token,
	}
}

func (ud *UserDTO) GetUserByMe(user domain.User) UserResponseByMe {
	return UserResponseByMe{
		User: User{
			CUSTID:     user.CUSTID,
			NOCUST:     user.NOCUST,
			NMCUST:     user.NMCUST,
			NUM2ND:     user.NUM2ND,
			STCUST:     user.STCUST,
			CODE01:     user.CODE01,
			DESC01:     user.DESC01,
			CODE02:     user.CODE02,
			DESC02:     user.DESC02,
			CODE03:     user.CODE03,
			DESC03:     user.DESC03,
			CODE04:     user.CODE04,
			DESC04:     user.DESC04,
			CODE05:     user.CODE05,
			DESC05:     user.DESC05,
			TOTPAY:     user.TOTPAY,
			GENUS:      user.GENUS,
			LastUpdate: user.LastUpdate,
			NO_WA:      user.NO_WA,
		},
	}
}
