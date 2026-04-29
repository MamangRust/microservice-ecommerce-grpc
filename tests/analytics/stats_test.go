package analytics_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	categoryhandler "github.com/MamangRust/monolith-ecommerce-grpc-apigateway/handler/category"
	"github.com/MamangRust/monolith-ecommerce-shared/errors"
	tests "github.com/MamangRust/monolith-ecommerce-test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type AnalyticsTestSuite struct {
	tests.BaseTestSuite
	echo *echo.Echo
}

func (s *AnalyticsTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupStatsReaderService() // This sets up s.Conns["stats-reader"]

	s.echo = echo.New()
	apiHandler := errors.NewApiHandler(s.Obs, s.Log)

	// Register Category Stats Handlers
	categoryhandler.RegisterCategoryHandler(&categoryhandler.DepsCategory{
		Client:      nil, // Not needed for stats-only tests
		StatsReader: s.Conns["stats-reader"],
		E:           s.echo,
		Logger:      s.Log,
		CacheStore:  s.GetCacheStore(),
		UploadImage: &tests.MockImageUpload{},
		ApiHandler:  apiHandler,
	})
}

func (s *AnalyticsTestSuite) TestCategoryStatsAPI() {
	// 1. Monthly Total Price
	req := httptest.NewRequest(http.MethodGet, "/api/category-stats/monthly-total-pricing?year=2024&month=1", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	
	s.Equal(http.StatusOK, rec.Code, rec.Body.String())
	
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
}

func (s *AnalyticsTestSuite) TestOrderStatsAPI() {
	// Test Monthly Revenue
	req := httptest.NewRequest(http.MethodGet, "/api/category-stats/monthly-pricing?year=2024", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	
	s.Equal(http.StatusOK, rec.Code, rec.Body.String())
	
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
}

func TestAnalyticsTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(AnalyticsTestSuite))
}
