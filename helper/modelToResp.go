package helper

import (
	"self-payrol/model"
	"self-payrol/response"
	"strconv"
)

// NewUserResp returns user data in the user response format
func NewUserResp(user *model.User) *response.UserResp {
	var position *response.PositionResp
	if user.Position != nil {
		position = NewPositionResp(user.Position)
	}

	return &response.UserResp{
		Id:         int(user.ID),
		SecretId:   user.SecretID,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Address:    user.Address,
		PositionId: user.PositionID,
		Position:   position,
		CreatedAt:  user.CreatedAt.String(),
		UpdatedAt:  user.UpdatedAt.String(),
	}
}

// NewPositionResp returns position data in the position response format
func NewPositionResp(position *model.Position) *response.PositionResp {
	return &response.PositionResp{
		Id:        strconv.Itoa(int(position.ID)),
		Name:      position.Name,
		Salary:    position.Salary,
		CreatedAt: position.CreatedAt.String(),
		UpdatedAt: position.UpdatedAt.String(),
	}
}

// NewCompanyResp returns company data in the company response format
func NewCompanyResp(company *model.Company) *response.CompanyResp {
	return &response.CompanyResp{
		Id:        int(company.ID),
		Name:      company.Name,
		Address:   company.Address,
		Balance:   company.Balance,
		CreatedAt: company.CreatedAt.String(),
		UpdatedAt: company.UpdatedAt.String(),
	}
}

// NewTransactionResp returns transaction data in the transaction response format
func NewTransactionResp(transaction *model.Transaction) *response.TransactionResp {
	return &response.TransactionResp{
		Id:        int(transaction.ID),
		Amount:    transaction.Amount,
		Note:      transaction.Note,
		Type:      transaction.Type,
		CreatedAt: transaction.CreatedAt.String(),
		UpdatedAt: transaction.UpdatedAt.String(),
	}
}
