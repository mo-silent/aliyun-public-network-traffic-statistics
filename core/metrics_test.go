package core

import (
	cms20190101 "github.com/alibabacloud-go/cms-20190101/v8/client"
	"github.com/alibabacloud-go/tea/tea"
	"os"
	"testing"
)

func TestTrafficCount(t *testing.T) {
	type args struct {
		client                    *cms20190101.Client
		describeMetricListRequest *cms20190101.DescribeMetricListRequest
	}
	clientConfig := ClientConfig{
		Type:            "access_key",
		AccessKeyId:     os.Getenv("ALICLOUD_ACCESS_KEY"),
		AccessKeySecret: os.Getenv("ALICLOUD_SECRET_KEY"),
	}
	client, err := clientConfig.CreateClient()
	if err != nil {
		t.Errorf("CreateClient() error = %v", err)
	}
	describeMetricListRequest := &cms20190101.DescribeMetricListRequest{
		Namespace:  tea.String("acs_alb"),
		MetricName: tea.String("ListenerInBits"),
		Period:     tea.String("3600"),
		StartTime:  tea.String("2024-03-01 00:00:00"),
		EndTime:    tea.String("2024-03-31 23:59:59"),
		Dimensions: tea.String("[{\"loadBalancerId\":\"alb-b28yqg6ryq2q2xzlpn\"}]"),
		Length:     tea.String("1440"),
		Express:    tea.String("{\"groupby\":[\"userId\",\"loadBalancerId\"]}"),
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			args:    args{client: client, describeMetricListRequest: describeMetricListRequest},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TrafficCount(tt.args.client, tt.args.describeMetricListRequest); (err != nil) != tt.wantErr {
				t.Errorf("TrafficCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
