package order

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestComputeTotal(t *testing.T) {
	t.Run("nominal cases", func(t *testing.T) {
		o := Order{
			ID:                "01",
			CurrencyAlphaCode: "INR",
			Items: []Item{
				{
					ID:        "01",
					Quantity:  2,
					UnitPrice: money.New(100, "INR"),
				},
				{
					ID:        "02",
					Quantity:  5,
					UnitPrice: money.New(500, "INR"),
				},
			},
		}
		total, err := o.ComputeTotal()
		assert.NoError(t, err)
		assert.Equal(t, int64(200)+int64(2500), total.Amount())
		assert.Equal(t, "INR", total.Currency().Code)
	})
	t.Run("currency issues", func(t *testing.T) {
		o := Order{
			ID:                "01",
			CurrencyAlphaCode: "INR",
			Items: []Item{
				{
					ID:        "01",
					Quantity:  2,
					UnitPrice: money.New(100, "INR"),
				},
				{
					ID:        "02",
					Quantity:  5,
					UnitPrice: money.New(500, "INR"),
				},
			},
		}
		_, err := o.ComputeTotal()
		assert.NoError(t, err)
	})
}
