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

// Allows the update of one or more parameters of a database in Amazon Lightsail.
//
// Parameter updates don't cause outages; therefore, their application is not
// subject to the preferred maintenance window. However, there are two ways in
// which parameter updates are applied: dynamic or pending-reboot . Parameters
// marked with a dynamic apply type are applied immediately. Parameters marked
// with a pending-reboot apply type are applied only after the database is
// rebooted using the reboot relational database operation.
//
// The update relational database parameters operation supports tag-based access
// control via resource tags applied to the resource identified by
// relationalDatabaseName. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags
func (c *Client) UpdateRelationalDatabaseParameters(ctx context.Context, params *UpdateRelationalDatabaseParametersInput, optFns ...func(*Options)) (*UpdateRelationalDatabaseParametersOutput, error) {
	if params == nil {
		params = &UpdateRelationalDatabaseParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRelationalDatabaseParameters", params, optFns, c.addOperationUpdateRelationalDatabaseParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRelationalDatabaseParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRelationalDatabaseParametersInput struct {

	// The database parameters to update.
	//
	// This member is required.
	Parameters []types.RelationalDatabaseParameter

	// The name of your database for which to update parameters.
	//
	// This member is required.
	RelationalDatabaseName *string

	noSmithyDocumentSerde
}

type UpdateRelationalDatabaseParametersOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRelationalDatabaseParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRelationalDatabaseParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRelationalDatabaseParameters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRelationalDatabaseParameters"); err != nil {
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
	if err = addOpUpdateRelationalDatabaseParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRelationalDatabaseParameters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRelationalDatabaseParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRelationalDatabaseParameters",
	}
}
