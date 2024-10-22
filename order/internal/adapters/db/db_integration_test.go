package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"log"
	"testing"
	"time"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	DataSourceUrl string
}

func (suite *OrderDatabaseTestSuite) SetupSuite() {

	ctx := context.Background()
	port := "5432/tcp"
	dbUrl := func(host string, p nat.Port) string {
		return fmt.Sprintf("postgres://patrol:patrol@%s:%s/patrol?sslmode=disable&TimeZone=Asia/Shanghai", host, p.Port())
	}
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "patrol",
			"POSTGRES_DB":       "patrol",
			"POSTGRES_USER":     "patrol",
		},
		WaitingFor: wait.ForSQL(nat.Port(port), "postgres", dbUrl).WithStartupTimeout(time.Second * 30),
	}
	testContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal("failed to start test container", err)
	}
	endpoint, err := testContainer.Endpoint(ctx, "")
	if err != nil {
		log.Fatal("failed to connect to test container", err)
	}
	suite.DataSourceUrl = fmt.Sprintf("postgres://patrol:patrol@%s/patrol?sslmode=disable&TimeZone=Asia/Shanghai", endpoint)
}

func (suite *OrderDatabaseTestSuite) TearDownSuite() {

}

func (suite *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	adapter, err := NewAdapter(suite.DataSourceUrl)
	suite.Nil(err)
	err = adapter.Save(&domain.Order{})
	suite.Nil(err)
}

func (suite *OrderDatabaseTestSuite) Test_Should_Get_Order() {
	adapter, err := NewAdapter(suite.DataSourceUrl)
	suite.Nil(err)
	order := domain.NewOrder(2, []domain.OrderItem{
		{
			ProductCode: "abc",
			Quantity:    5,
			UnitPrice:   1.01,
		},
	})
	err = adapter.Save(&order)
	suite.Nil(err)
	record, err := adapter.Get(order.Id)
	if str, err := json.Marshal(record); err == nil {
		fmt.Println(string(str))
	}
	suite.Equal(order.CustomerId, record.CustomerId)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}
