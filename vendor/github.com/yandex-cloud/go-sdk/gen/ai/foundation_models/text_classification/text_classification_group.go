// Code generated by sdkgen. DO NOT EDIT.

package text_classification

import (
	"context"

	"google.golang.org/grpc"
)

// FoundationModelsTextClassification provides access to "text_classification" component of Yandex.Cloud
type FoundationModelsTextClassification struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewFoundationModelsTextClassification creates instance of FoundationModelsTextClassification
func NewFoundationModelsTextClassification(g func(ctx context.Context) (*grpc.ClientConn, error)) *FoundationModelsTextClassification {
	return &FoundationModelsTextClassification{g}
}

// TextClassification gets TextClassificationService client
func (f *FoundationModelsTextClassification) TextClassification() *TextClassificationServiceClient {
	return &TextClassificationServiceClient{getConn: f.getConn}
}