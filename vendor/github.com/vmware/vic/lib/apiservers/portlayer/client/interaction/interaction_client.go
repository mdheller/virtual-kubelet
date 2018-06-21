package interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new interaction API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for interaction API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ContainerCloseStdin closes stdin

Close a stdin for the container
*/
func (a *Client) ContainerCloseStdin(params *ContainerCloseStdinParams) (*ContainerCloseStdinOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerCloseStdinParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerCloseStdin",
		Method:             "DELETE",
		PathPattern:        "/interaction/{id}/stdin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerCloseStdinReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerCloseStdinOK), nil

}

/*
ContainerGetStderr gets stderr

Get a stderr for the container
*/
func (a *Client) ContainerGetStderr(params *ContainerGetStderrParams, writer io.Writer) (*ContainerGetStderrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerGetStderrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerGetStderr",
		Method:             "GET",
		PathPattern:        "/interaction/{id}/stderr",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerGetStderrReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerGetStderrOK), nil

}

/*
ContainerGetStdout gets stdout

Get a stdout for the container
*/
func (a *Client) ContainerGetStdout(params *ContainerGetStdoutParams, writer io.Writer) (*ContainerGetStdoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerGetStdoutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerGetStdout",
		Method:             "GET",
		PathPattern:        "/interaction/{id}/stdout",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerGetStdoutReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerGetStdoutOK), nil

}

/*
ContainerResize resizes tty session

Resize the container's tty session
*/
func (a *Client) ContainerResize(params *ContainerResizeParams) (*ContainerResizeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerResizeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerResize",
		Method:             "POST",
		PathPattern:        "/interaction/{id}/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerResizeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerResizeOK), nil

}

/*
ContainerSetStdin sets stdin

Set a stdin for the container
*/
func (a *Client) ContainerSetStdin(params *ContainerSetStdinParams) (*ContainerSetStdinOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerSetStdinParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerSetStdin",
		Method:             "POST",
		PathPattern:        "/interaction/{id}/stdin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/raw-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerSetStdinReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerSetStdinOK), nil

}

/*
InteractionBind enables bind activate the virtual device

Enable/bind/activate the VirtualDevice
*/
func (a *Client) InteractionBind(params *InteractionBindParams) (*InteractionBindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInteractionBindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InteractionBind",
		Method:             "POST",
		PathPattern:        "/interaction/binding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InteractionBindReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InteractionBindOK), nil

}

/*
InteractionJoin adds interaction capability

Adds interaction capabilities to given handle
*/
func (a *Client) InteractionJoin(params *InteractionJoinParams) (*InteractionJoinOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInteractionJoinParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InteractionJoin",
		Method:             "POST",
		PathPattern:        "/interaction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InteractionJoinReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InteractionJoinOK), nil

}

/*
InteractionUnbind disables unbind deactivate the virtual device

Disable/unbind/deactivate the VirtualDevice
*/
func (a *Client) InteractionUnbind(params *InteractionUnbindParams) (*InteractionUnbindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInteractionUnbindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InteractionUnbind",
		Method:             "DELETE",
		PathPattern:        "/interaction/binding",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InteractionUnbindReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InteractionUnbindOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}