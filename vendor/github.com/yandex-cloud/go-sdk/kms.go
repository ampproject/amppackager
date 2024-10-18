package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/kms"
	kmsasymmetricencryptioncrypto "github.com/yandex-cloud/go-sdk/gen/kmsasymmetricencryptioncrypto"
	kmsasymmetricencryption "github.com/yandex-cloud/go-sdk/gen/kmsasymmetricencryptionkey"
	kmsasymmetricsignaturecrypto "github.com/yandex-cloud/go-sdk/gen/kmsasymmetricsignaturecrypto"
	kmsasymmetricsignature "github.com/yandex-cloud/go-sdk/gen/kmsasymmetricsignaturekey"
	kmscrypto "github.com/yandex-cloud/go-sdk/gen/kmscrypto"
)

const (
	KMSServiceID                           = "kms"
	KMSCryptoServiceID                     = "kms-crypto"
	KMSAsymmetricEncryptionServiceID       = "kms"
	KMSAsymmetricEncryptionCryptoServiceID = "kms-crypto"
	KMSAsymmetricSignatureServiceID        = "kms"
	KMSAsymmetricSignatureCryptoServiceID  = "kms-crypto"
)

func (sdk *SDK) KMS() *kms.KMS {
	return kms.NewKMS(sdk.getConn(KMSServiceID))
}

func (sdk *SDK) KMSCrypto() *kmscrypto.KMSCrypto {
	return kmscrypto.NewKMSCrypto(sdk.getConn(KMSCryptoServiceID))
}

func (sdk *SDK) KMSAsymmetricEncryption() *kmsasymmetricencryption.KMSAsymmetricEncryption {
	return kmsasymmetricencryption.NewKMSAsymmetricEncryption(sdk.getConn(KMSAsymmetricEncryptionServiceID))
}

func (sdk *SDK) KMSAsymmetricEncryptionCrypto() *kmsasymmetricencryptioncrypto.KMSAsymmetricEncryptionCrypto {
	return kmsasymmetricencryptioncrypto.NewKMSAsymmetricEncryptionCrypto(sdk.getConn(KMSAsymmetricEncryptionCryptoServiceID))
}

func (sdk *SDK) KMSAsymmetricSignature() *kmsasymmetricsignature.KMSAsymmetricSignature {
	return kmsasymmetricsignature.NewKMSAsymmetricSignature(sdk.getConn(KMSAsymmetricSignatureServiceID))
}

func (sdk *SDK) KMSAsymmetricSignatureCrypto() *kmsasymmetricsignaturecrypto.KMSAsymmetricSignatureCrypto {
	return kmsasymmetricsignaturecrypto.NewKMSAsymmetricSignatureCrypto(sdk.getConn(KMSAsymmetricSignatureCryptoServiceID))
}
