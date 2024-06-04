package rsa

import (
	"crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	testPrvKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA01MbwDEV88NwSieTwf43s9L0n3co2yPixYtsDT8QW0VeqaGp
p0g4c1nMVmpP1h9w62p+cC0r52CWgFfR2INkRDbfMh7+wsSXPVAfJf75IMTrFX0B
Fpoi5aI24rXBROX3iA4fM+VvRo0zq68iGCHkh38937Fbd+ex61yNWb2EnkjS1oXY
rFV1BVa/rwaDm3ayxMbOshYgqalObMSnpmeMkC0u5TwRFHOCG5qrIr56/YTmgH1r
M1/R2GUqlJSiP3Wf3SPm/no1+toK+Ly4fqS9N/vSIUUlQsjNjJxWnxBOZKRFrmuo
9rEbRH6j/Nb6BI5Eyv/GM+7zdHpnLnNfkykBOQIDAQABAoIBAAGaQr6dYK49kC9v
x2YUOHXrpqFC9RNcXU7kNMiQ7t3SU9pzeGblXFA7oRoSqlJUYLXYU+cj4I5bpCiL
AfNjY7JEvEKUC5iMhx0W63WNTnRaKfI6xXRXbOZS6ZmkiTetisgEW2Y4WjStw3OM
23RVZXdFlsGIhcoN/sJrLEM+9a+BmFaRfBPytq9vzWRu6AT9FxrZoXyJy+vWsAV2
8NGDxq4/dE4lfEew34AaVQTr1Dq5I8hLRbJ7ZcIDttzWo1x2hWay1+RymxBW1gZ1
T/7+k82E7fFn/NZOsRbk5D/37veP8z84hIe7SdgLH42YdHAZmQ7OseEbaww56wcb
4SM6vjUCgYEA2XMtRytxuA63zli4LFMcZOHksOwA+WQ7qBaH6bGobqtMeyPaUn4Y
nFU9+1pgG92cdvcQQ1ikpgOtDfqPx6Y60g845bn5RmfMHner7G9dWcaqahVZOmqC
SrdIznfgxkYvKU7C0bCPxbokwwFsby6xXFn/fLcyo4n2n1DOV5S6vB8CgYEA+Mnw
s8qnqU/+1m5ch8FXrj0pBoHx4/t07uLgZSA5F7NcCfsuBAigVhO2P+Cn87OMfTiH
bgiGgidn0iTV4LpBkfDRdgsScWfWnxbIc05W8hKt4ncYmtqUFdjxE0yxXiQqhV00
jwhpqlwChXDaTL0ewh2vEGoKAtXtwi+Xgwtel6cCgYBkioqe44l7DuMRt5fAdtUk
GAZDf3ub1Cp5N6Gz/f7g3LdKIHOrvL3oWsmD1G+nsrnLj8SnPu8yC3USh42/RC7i
PUOThH+rfAa296I4ee1xuxfEYQaWqfSAU2qIfIkjZAQeV1pg1gBD7iNdPuVCKxa6
mqo0ogf58apkU1p0yEHnBQKBgQCGm6R7csOiTSEB3jZ8UTTilj7TQY2iH5SWB/UH
Yhbh5u5+jAPKtOwjhojOKPxWVChIPESyWTSFyVJYFgwOilgd4WFDcBwrddZev1H3
aUSLt2WTqYKLjoYfCADvw0gYOpMzE0nztcaOIThQAM5sRMsWlj75L5Z5EEyTC3L7
uTQV8wKBgQC0WnsJ1FVMId1cCd1wzzr/DnXO0D2ihEqYLvmpxxGXPKDlDHwTX+Rz
Rfbiihwt8sxLBszt0TMx+Yp36SUboVbdcOc1h/1gMcUP1g8sKLWkyd9iXwKwh2mT
KsN1Y5uLoJ8rokTkmdxkySCGYu9PYiRhTbTlSsGqI39KdtY0iLvUMg==
-----END RSA PRIVATE KEY-----`
	testPubKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA01MbwDEV88NwSieTwf43
s9L0n3co2yPixYtsDT8QW0VeqaGpp0g4c1nMVmpP1h9w62p+cC0r52CWgFfR2INk
RDbfMh7+wsSXPVAfJf75IMTrFX0BFpoi5aI24rXBROX3iA4fM+VvRo0zq68iGCHk
h38937Fbd+ex61yNWb2EnkjS1oXYrFV1BVa/rwaDm3ayxMbOshYgqalObMSnpmeM
kC0u5TwRFHOCG5qrIr56/YTmgH1rM1/R2GUqlJSiP3Wf3SPm/no1+toK+Ly4fqS9
N/vSIUUlQsjNjJxWnxBOZKRFrmuo9rEbRH6j/Nb6BI5Eyv/GM+7zdHpnLnNfkykB
OQIDAQAB
-----END RSA PUBLIC KEY-----`
	testPassword = "123456"
	testBits     = 2048
)

func TestGenRsaKey(t *testing.T) {
	prvKey, pubKey, err := GenKey(testBits)
	require.Nil(t, err)
	require.Greater(t, len(prvKey), 0)
	require.Greater(t, len(pubKey), 0)
}

func TestRsaCrypt(t *testing.T) {
	cipher, err := Encrypt([]byte(testPubKey), []byte(testPassword))
	require.Nil(t, err)
	require.Equal(t, len(cipher), testBits/8)

	plain, err := Decrypt([]byte(testPrvKey), cipher)
	require.Nil(t, err)
	require.Equal(t, string(plain), testPassword)
}

func TestRsaSign(t *testing.T) {
	sig, err := Sign([]byte(testPrvKey), crypto.SHA256, []byte(testPassword))
	require.Nil(t, err)
	require.Equal(t, len(sig), testBits/8)

	err = VerifySign([]byte(testPubKey), crypto.SHA256, []byte(testPassword), sig)
	require.Nil(t, err)
}
