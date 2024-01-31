package interactor

import (
	"currency-rate/internal/usecase"
	mock_usecase "currency-rate/internal/usecase/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestReplacingDash(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"2022-01-01", "2022/01/01"},
		{"2022-02-02", "2022/02/02"},
		{"2022-03-03", "2022/03/03"},
	}

	for _, tt := range tests {
		out := replacingDash(tt.in)
		if out != tt.out {
			t.Errorf("Expected %s, got %s", tt.out, out)
		}
	}
}

type SuiteCurrency struct {
	suite.Suite
	repo *mock_usecase.MockCurrencyCache
	w    *mock_usecase.MockCurrencyGateway
}

func (this *SuiteCurrency) SetupSuite() {
	ctl := gomock.NewController(this.T())
	defer ctl.Finish()
	this.repo = mock_usecase.NewMockCurrencyCache(ctl)
	this.w = mock_usecase.NewMockCurrencyGateway(ctl)
}

func (this *SuiteCurrency) TestGetCurrencyWithOKCache() {
	respCurrency := NewGetCurrency(this.w, this.repo)
	date := time.Now()
	nameCur := "AUD"
	expectedCurrency := &usecase.CurrencyDTO{
		StartDate: date,
		EndDate:   date,
		Name:      "Австралийский доллар",
		CharCode:  "AUD",
		Value:     58.2831,
	}
	this.repo.EXPECT().GetCurrencyDate("2024/01/31", "AUD").Return(expectedCurrency, nil).Times(1)
	currency, err := respCurrency.GetCurrency(date, nameCur)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), expectedCurrency, currency)
}
func (this *SuiteCurrency) TestGetCurrencyWithNotOKCacheError() {
	respCurrency := NewGetCurrency(this.w, this.repo)
	date := time.Now()
	nameCur := "AUD"
	errCust := errors.New("Error")
	dto := usecase.CurrencyDTO{}
	dtoCur := usecase.AllCur{}
	this.repo.EXPECT().GetCurrencyDate("2024/01/31", "AUD").Return(&dto, errCust).Times(1)
	this.w.EXPECT().GetQuotes("2024/01/31").Return(&dtoCur, errCust).Times(1)
	currency, err := respCurrency.GetCurrency(date, nameCur)
	assert.Error(this.T(), err, errCust)
	assert.Nil(this.T(), currency)
}

func (this *SuiteCurrency) TestGetCurrencyWithNotOKCache() {
	respCurrency := NewGetCurrency(this.w, this.repo)
	date := time.Now()
	nameCur := "AUD"
	errCust := errors.New("Error")
	dto := usecase.CurrencyDTO{}
	mapValute := make(map[string]usecase.Valute)
	valuteWeb := usecase.Valute{
		ID:       "",
		NumCode:  "",
		CharCode: nameCur,
		Nominal:  1,
		Name:     "Австралийский доллар",
		Value:    58.283,
		Previous: 0.1,
	}
	mapValute[nameCur] = valuteWeb
	webDTO := &usecase.AllCur{
		Date:         date,
		PreviousDate: date,
		PreviousURL:  "",
		Timestamp:    date,
		Valute:       mapValute,
	}

	expectedCurrency := &usecase.CurrencyDTO{
		StartDate: date,
		EndDate:   date,
		Name:      "Австралийский доллар",
		CharCode:  "AUD",
		Value:     58.283,
	}
	currencyList := []usecase.CurrencyDTO{*expectedCurrency}
	this.repo.EXPECT().GetCurrencyDate("2024/01/31", "AUD").Return(&dto, errCust).Times(1)
	this.w.EXPECT().GetQuotes("2024/01/31").Return(webDTO, nil).Times(1)
	this.repo.EXPECT().InsertCurrencyDate(currencyList).Return(errCust).Times(1)
	currency, err := respCurrency.GetCurrency(date, nameCur)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), expectedCurrency, currency)
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, new(SuiteCurrency))
}
