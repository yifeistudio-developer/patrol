package e2e

import (
	"context"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/yifeistudio-developer/wharf/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

type CreateOrderTestSuite struct {
	suite.Suite
	compose tc.ComposeStack
}

func (c *CreateOrderTestSuite) SetupSuite() {
	composeFilePaths := []string{"resources/docker-compose.yml"}
	compose, err := tc.NewDockerCompose(composeFilePaths...)
	c.Nil(err)
	c.compose = compose
	ctx := context.Background()
	err = compose.Up(ctx, tc.Wait(true))
	c.Nil(err)
}

func (c *CreateOrderTestSuite) TearDownSuite() {
	ctx := context.Background()
	err := c.compose.Down(ctx)
	if err != nil {
		log.Fatalf("could not down compose stack: %v", err)
	}
}

func (c *CreateOrderTestSuite) Test_Should_Create_Order() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8080", opts...)
	c.Nil(err)
	defer conn.Close()
	client := order.NewOrderClient(conn)
	_, err = client.Create(context.Background(), &order.CreateOrderRequest{
		CustomerId: 1,
	})
	c.Nil(err)

}

func TestCreateOrderTestSuite(t *testing.T) {
	suite.Run(t, &CreateOrderTestSuite{})
}
