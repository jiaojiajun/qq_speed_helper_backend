package car_test

import (
	"context"

	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jiaojiajun/qq_speed_helper_backend/biz/handler/helper/car"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCarsInner_happy_case_1(t *testing.T) {
	var assert = assert.New(t)
	var ctx = context.TODO()
	var req = &app.RequestContext{}

	//act
	resp, err := car.GetAllCarsInner(ctx, req)
	assert.NoError(err)
	assert.NotNil(resp)

	assert.Equal(len(resp.Cars), 100)
}
