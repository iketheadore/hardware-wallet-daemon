// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteWipe clean all the configurations.
*/
func (a *Client) DeleteWipe(params *DeleteWipeParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWipeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWipeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteWipe",
		Method:             "DELETE",
		PathPattern:        "/wipe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWipeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWipeOK), nil

}

/*
GetCsrf Returns csrf token
*/
func (a *Client) GetCsrf(params *GetCsrfParams) (*GetCsrfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCsrfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCsrf",
		Method:             "GET",
		PathPattern:        "/csrf",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCsrfReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCsrfOK), nil

}

/*
GetFeatures Returns device information.
*/
func (a *Client) GetFeatures(params *GetFeaturesParams, authInfo runtime.ClientAuthInfoWriter) (*GetFeaturesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeaturesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFeatures",
		Method:             "GET",
		PathPattern:        "/features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFeaturesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFeaturesOK), nil

}

/*
PostApplySettings Apply hardware wallet settings.
*/
func (a *Client) PostApplySettings(params *PostApplySettingsParams, authInfo runtime.ClientAuthInfoWriter) (*PostApplySettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostApplySettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostApplySettings",
		Method:             "POST",
		PathPattern:        "/applySettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostApplySettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostApplySettingsOK), nil

}

/*
PostBackup Start seed backup procedure.
*/
func (a *Client) PostBackup(params *PostBackupParams, authInfo runtime.ClientAuthInfoWriter) (*PostBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostBackup",
		Method:             "POST",
		PathPattern:        "/backup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBackupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostBackupOK), nil

}

/*
PostCheckMessageSignature Check a message signature matches the given address.
*/
func (a *Client) PostCheckMessageSignature(params *PostCheckMessageSignatureParams, authInfo runtime.ClientAuthInfoWriter) (*PostCheckMessageSignatureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCheckMessageSignatureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCheckMessageSignature",
		Method:             "POST",
		PathPattern:        "/checkMessageSignature",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostCheckMessageSignatureReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCheckMessageSignatureOK), nil

}

/*
PostGenerateAddresses Generate addresses for the hardware wallet seed.
*/
func (a *Client) PostGenerateAddresses(params *PostGenerateAddressesParams, authInfo runtime.ClientAuthInfoWriter) (*PostGenerateAddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostGenerateAddressesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostGenerateAddresses",
		Method:             "POST",
		PathPattern:        "/generateAddresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostGenerateAddressesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostGenerateAddressesOK), nil

}

/*
PostGenerateMnemonic Generate mnemonic can be used to initialize the device with a random seed.
*/
func (a *Client) PostGenerateMnemonic(params *PostGenerateMnemonicParams, authInfo runtime.ClientAuthInfoWriter) (*PostGenerateMnemonicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostGenerateMnemonicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostGenerateMnemonic",
		Method:             "POST",
		PathPattern:        "/generateMnemonic",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostGenerateMnemonicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostGenerateMnemonicOK), nil

}

/*
PostIntermediatePassPhrase passphrase ack request.
*/
func (a *Client) PostIntermediatePassPhrase(params *PostIntermediatePassPhraseParams, authInfo runtime.ClientAuthInfoWriter) (*PostIntermediatePassPhraseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIntermediatePassPhraseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIntermediatePassPhrase",
		Method:             "POST",
		PathPattern:        "/intermediate/passPhrase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIntermediatePassPhraseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostIntermediatePassPhraseOK), nil

}

/*
PostIntermediatePinMatrix pin matrix ack request.
*/
func (a *Client) PostIntermediatePinMatrix(params *PostIntermediatePinMatrixParams, authInfo runtime.ClientAuthInfoWriter) (*PostIntermediatePinMatrixOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIntermediatePinMatrixParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIntermediatePinMatrix",
		Method:             "POST",
		PathPattern:        "/intermediate/pinMatrix",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIntermediatePinMatrixReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostIntermediatePinMatrixOK), nil

}

/*
PostIntermediateWord word ack request.
*/
func (a *Client) PostIntermediateWord(params *PostIntermediateWordParams, authInfo runtime.ClientAuthInfoWriter) (*PostIntermediateWordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIntermediateWordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIntermediateWord",
		Method:             "POST",
		PathPattern:        "/intermediate/word",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIntermediateWordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostIntermediateWordOK), nil

}

/*
PostRecovery Recover existing wallet using seed.
*/
func (a *Client) PostRecovery(params *PostRecoveryParams, authInfo runtime.ClientAuthInfoWriter) (*PostRecoveryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRecoveryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRecovery",
		Method:             "POST",
		PathPattern:        "/recovery",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostRecoveryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRecoveryOK), nil

}

/*
PostSetMnemonic Set mnemonic can be used to initialize the device with your own seed.
*/
func (a *Client) PostSetMnemonic(params *PostSetMnemonicParams, authInfo runtime.ClientAuthInfoWriter) (*PostSetMnemonicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSetMnemonicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSetMnemonic",
		Method:             "POST",
		PathPattern:        "/setMnemonic",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSetMnemonicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSetMnemonicOK), nil

}

/*
PostSetPinCode Configure a pin code on the device.
*/
func (a *Client) PostSetPinCode(params *PostSetPinCodeParams, authInfo runtime.ClientAuthInfoWriter) (*PostSetPinCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSetPinCodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSetPinCode",
		Method:             "POST",
		PathPattern:        "/setPinCode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSetPinCodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSetPinCodeOK), nil

}

/*
PostSignMessage Sign a message using the secret key at given index.
*/
func (a *Client) PostSignMessage(params *PostSignMessageParams, authInfo runtime.ClientAuthInfoWriter) (*PostSignMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSignMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSignMessage",
		Method:             "POST",
		PathPattern:        "/signMessage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSignMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSignMessageOK), nil

}

/*
PostTransactionSign Sign a transaction with the hardware wallet.
*/
func (a *Client) PostTransactionSign(params *PostTransactionSignParams, authInfo runtime.ClientAuthInfoWriter) (*PostTransactionSignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTransactionSignParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionSign",
		Method:             "POST",
		PathPattern:        "/transactionSign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTransactionSignReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionSignOK), nil

}

/*
PutCancel Cancels the current operation.
*/
func (a *Client) PutCancel(params *PutCancelParams, authInfo runtime.ClientAuthInfoWriter) (*PutCancelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCancelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCancel",
		Method:             "PUT",
		PathPattern:        "/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutCancelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCancelOK), nil

}

/*
PutFirmwareUpdate Update firmware
*/
func (a *Client) PutFirmwareUpdate(params *PutFirmwareUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*PutFirmwareUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFirmwareUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutFirmwareUpdate",
		Method:             "PUT",
		PathPattern:        "/firmwareUpdate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutFirmwareUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutFirmwareUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
