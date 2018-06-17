package strfmt

import (
	github_com_go_courier_validator "github.com/go-courier/validator"
)

var ASCIIValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringASCII, "ascii")
var AlphaValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringAlpha, "alpha")
var AlphaNumericValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringAlphaNumeric, "alpha-numeric", "alphaNumeric")
var AlphaUnicodeValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringAlphaUnicode, "alpha-unicode", "alphaUnicode")
var AlphaUnicodeNumericValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringAlphaUnicodeNumeric, "alpha-unicode-numeric", "alphaUnicodeNumeric")
var Base64Validator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringBase64, "base64")
var Base64URLValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringBase64URL, "base64-url", "base64URL")
var BtcAddressValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringBtcAddress, "btc-address", "btcAddress")
var BtcAddressLowerValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringBtcAddressLower, "btc-address-lower", "btcAddressLower")
var BtcAddressUpperValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringBtcAddressUpper, "btc-address-upper", "btcAddressUpper")
var DataURIValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringDataURI, "data-uri", "dataURI")
var EmailValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringEmail, "email")
var EthAddressValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringEthAddress, "eth-address", "ethAddress")
var EthAddressLowerValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringEthAddressLower, "eth-address-lower", "ethAddressLower")
var EthAddressUpperValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringEthAddressUpper, "eth-address-upper", "ethAddressUpper")
var HslValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringHSL, "hsl")
var HslaValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringHSLA, "hsla")
var HexAdecimalValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringHexAdecimal, "hex-adecimal", "hexAdecimal")
var HexColorValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringHexColor, "hex-color", "hexColor")
var HostnameValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringHostname, "hostname")
var HostnameXValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringHostnameX, "hostname-x", "hostnameX")
var Isbn10Validator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringISBN10, "isbn10")
var Isbn13Validator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringISBN13, "isbn13")
var LatitudeValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringLatitude, "latitude")
var LongitudeValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringLongitude, "longitude")
var MultibyteValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringMultibyte, "multibyte")
var NumberValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringNumber, "number")
var NumericValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringNumeric, "numeric")
var PrintableASCIIValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringPrintableASCII, "printable-ascii", "printableASCII")
var RgbValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringRGB, "rgb")
var RgbaValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringRGBA, "rgba")
var SsnValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringSSN, "ssn")
var UUIDValidator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringUUID, "uuid")
var Uuid3Validator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringUUID3, "uuid3")
var Uuid4Validator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringUUID4, "uuid4")
var Uuid5Validator = github_com_go_courier_validator.NewRegexpStrfmtValidator(regexpStringUUID5, "uuid5")