// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new access key for the specified Amazon Lightsail bucket. Access keys
// consist of an access key ID and corresponding secret access key.
//
// Access keys grant full programmatic access to the specified bucket and its
// objects. You can have a maximum of two access keys per bucket. Use the [GetBucketAccessKeys]action
// to get a list of current access keys for a specific bucket. For more information
// about access keys, see [Creating access keys for a bucket in Amazon Lightsail]in the Amazon Lightsail Developer Guide.
//
// The secretAccessKey value is returned only in response to the
// CreateBucketAccessKey action. You can get a secret access key only when you
// first create an access key; you cannot get the secret access key later. If you
// lose the secret access key, you must create a new access key.
//
// [Creating access keys for a bucket in Amazon Lightsail]: https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-creating-bucket-access-keys
// [GetBucketAccessKeys]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBucketAccessKeys.html
func (c *Client) CreateBucketAccessKey(ctx context.Context, params *CreateBucketAccessKeyInput, optFns ...func(*Options)) (*CreateBucketAccessKeyOutput, error) {
	if params == nil {
		params = &CreateBucketAccessKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucketAccessKey", params, optFns, c.addOperationCreateBucketAccessKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketAccessKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketAccessKeyInput struct {

	// The name of the bucket that the new access key will belong to, and grant access
	// to.
	//
	// This member is required.
	BucketName *string

	noSmithyDocumentSerde
}

type CreateBucketAccessKeyOutput struct {

	// An object that describes the access key that is created.
	AccessKey *types.AccessKey

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBucketAccessKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBucketAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBucketAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBucketAccessKey"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBucketAccessKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucketAccessKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateBucketAccessKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBucketAccessKey",
	}
}
