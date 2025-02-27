package cloudfront

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindFunctionByNameAndStage(conn *cloudfront.CloudFront, name, stage string) (*cloudfront.DescribeFunctionOutput, error) {
	input := &cloudfront.DescribeFunctionInput{
		Name:  aws.String(name),
		Stage: aws.String(stage),
	}

	output, err := conn.DescribeFunction(input)

	if tfawserr.ErrCodeEquals(err, cloudfront.ErrCodeNoSuchFunctionExists) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output, nil
}

// FindRealtimeLogConfigByARN returns the real-time log configuration corresponding to the specified ARN.
// Returns nil if no configuration is found.
func FindRealtimeLogConfigByARN(conn *cloudfront.CloudFront, arn string) (*cloudfront.RealtimeLogConfig, error) {
	input := &cloudfront.GetRealtimeLogConfigInput{
		ARN: aws.String(arn),
	}

	output, err := conn.GetRealtimeLogConfig(input)
	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, nil
	}

	return output.RealtimeLogConfig, nil
}

// FindMonitoringSubscriptionByDistributionID returns the monitoring subscription corresponding to the specified distribution id.
// Returns nil if no subscription is found.
func FindMonitoringSubscriptionByDistributionID(conn *cloudfront.CloudFront, id string) (*cloudfront.MonitoringSubscription, error) {
	input := &cloudfront.GetMonitoringSubscriptionInput{
		DistributionId: aws.String(id),
	}

	output, err := conn.GetMonitoringSubscription(input)
	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, nil
	}

	return output.MonitoringSubscription, nil
}

func FindResponseHeadersPolicyByID(conn *cloudfront.CloudFront, id string) (*cloudfront.GetResponseHeadersPolicyOutput, error) {
	input := &cloudfront.GetResponseHeadersPolicyInput{
		Id: aws.String(id),
	}

	output, err := conn.GetResponseHeadersPolicy(input)

	if tfawserr.ErrCodeEquals(err, cloudfront.ErrCodeNoSuchResponseHeadersPolicy) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output, nil
}
