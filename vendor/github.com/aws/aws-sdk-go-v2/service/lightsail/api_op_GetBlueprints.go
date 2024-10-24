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

// Returns the list of available instance images, or blueprints. You can use a
// blueprint to create a new instance already running a specific operating system,
// as well as a preinstalled app or development stack. The software each instance
// is running depends on the blueprint image you choose.
//
// Use active blueprints when creating new instances. Inactive blueprints are
// listed to support customers with existing instances and are not necessarily
// available to create new instances. Blueprints are marked inactive when they
// become outdated due to operating system updates or new application releases.
func (c *Client) GetBlueprints(ctx context.Context, params *GetBlueprintsInput, optFns ...func(*Options)) (*GetBlueprintsOutput, error) {
	if params == nil {
		params = &GetBlueprintsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBlueprints", params, optFns, c.addOperationGetBlueprintsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBlueprintsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBlueprintsInput struct {

	// Returns a list of blueprints that are specific to Lightsail for Research.
	//
	// You must use this parameter to view Lightsail for Research blueprints.
	AppCategory types.AppCategory

	// A Boolean value that indicates whether to include inactive (unavailable)
	// blueprints in the response of your request.
	IncludeInactive *bool

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetBlueprints request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string

	noSmithyDocumentSerde
}

type GetBlueprintsOutput struct {

	// An array of key-value pairs that contains information about the available
	// blueprints.
	Blueprints []types.Blueprint

	// The token to advance to the next page of results from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetBlueprints request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBlueprintsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBlueprints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBlueprints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBlueprints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlueprints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBlueprints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBlueprints",
	}
}