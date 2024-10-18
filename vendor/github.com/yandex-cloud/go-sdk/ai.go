// Copyright (c) 2020 YANDEX LLC.

package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/ai/foundation_models/embedding"
	"github.com/yandex-cloud/go-sdk/gen/ai/foundation_models/image_generation"
	"github.com/yandex-cloud/go-sdk/gen/ai/foundation_models/text_classification"
	"github.com/yandex-cloud/go-sdk/gen/ai/foundation_models/text_generation"
	"github.com/yandex-cloud/go-sdk/gen/ai/ocr"
	"github.com/yandex-cloud/go-sdk/gen/ai/stt"
	sttv3 "github.com/yandex-cloud/go-sdk/gen/ai/sttv3"
	"github.com/yandex-cloud/go-sdk/gen/ai/translate"
	"github.com/yandex-cloud/go-sdk/gen/ai/vision"
)

const (
	AITranslate Endpoint = "ai-translate"
	AIVision    Endpoint = "ai-vision"
	AISTT       Endpoint = "ai-stt"
	AISTTV3     Endpoint = "ai-stt-v3"
	AIOCR       Endpoint = "ai-vision-ocr"
	AIFM        Endpoint = "ai-foundation-models"
)

type AI struct {
	sdk *SDK
}

func (m *AI) Translate() *translate.Translate {
	return translate.NewTranslate(m.sdk.getConn(AITranslate))
}

func (m *AI) Vision() *vision.Vision {
	return vision.NewVision(m.sdk.getConn(AIVision))
}

func (m *AI) STT() *stt.STT {
	return stt.NewSTT(m.sdk.getConn(AISTT))
}

func (m *AI) STTV3() *sttv3.STT {
	return sttv3.NewSTT(m.sdk.getConn(AISTTV3))
}

func (m *AI) TextGeneration() *text_generation.FoundationModelsTextGeneration {
	return text_generation.NewFoundationModelsTextGeneration(m.sdk.getConn(AIFM))
}

func (m *AI) TextClassification() *text_classification.FoundationModelsTextClassification {
	return text_classification.NewFoundationModelsTextClassification(m.sdk.getConn(AIFM))
}

func (m *AI) ImageGeneration() *image_generation.FoundationModelsImageGeneration {
	return image_generation.NewFoundationModelsImageGeneration(m.sdk.getConn(AIFM))
}

func (m *AI) Embeddings() *embedding.FoundationModelsEmbedding {
	return embedding.NewFoundationModelsEmbedding(m.sdk.getConn(AIFM))
}

func (m *AI) OCR() *ocr.OCR {
	return ocr.NewOCR(m.sdk.getConn(AIOCR))
}
