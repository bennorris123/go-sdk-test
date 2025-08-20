// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest

import (
	"github.com/bennorris123/go-sdk-test/internal/apierror"
	"github.com/bennorris123/go-sdk-test/packages/param"
	"github.com/bennorris123/go-sdk-test/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type OpenAICompletionTokensDetails = shared.OpenAICompletionTokensDetails

// This is an alias to an internal type.
type OpenAIPromptTokensDetails = shared.OpenAIPromptTokensDetails

// This is an alias to an internal type.
type OpenAIUsage = shared.OpenAIUsage
