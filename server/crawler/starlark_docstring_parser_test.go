package crawler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDocstringParser(t *testing.T) {
	docstringContent := `This is a docstring comment

	Args:
		origin_chain (string): The value of this field should be the same as the value provided to
			the --local flag when the Hyperlane contract was deployed.
		validator_key (string): 
		remote_chains (dict[string, string]): A dictionary whose keys correspond to the chains 
			specified in the --remotes flag during deploy, and whose values are the RPC URLs where 
			those chains can be reached. 
	Returns:
		Returns a bunch of stuff.
		And also this
`
	result, err := ParseRunFunctionDocstring(docstringContent)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestDocstringParser_NoArgs(t *testing.T) {
	docstring := `This is a docstring comment

	Returns:
		Returns a bunch of stuff.
`
	result, err := ParseRunFunctionDocstring(docstring)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestDocstringParser_NoReturns(t *testing.T) {
	docstring := `This is a docstring comment

	Args:
		origin_chain (string): The value of this field should be the same as the value provided to
			the --local flag when the Hyperlane contract was deployed.
		validator_key (string): 
		remote_chains (dict[string, string]): A dictionary whose keys correspond to the chains 
			specified in the --remotes flag during deploy, and whose values are the RPC URLs where 
			those chains can be reached.
`
	result, err := ParseRunFunctionDocstring(docstring)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestDocstringParser_DescriptionOnly(t *testing.T) {
	docstring := `This is a docstring comment
`
	result, err := ParseRunFunctionDocstring(docstring)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestDocstringParser_NoDescription(t *testing.T) {
	docstring := `Args:
		origin_chain (string): The value of this field should be the same as the value provided to
			the --local flag when the Hyperlane contract was deployed.
		validator_key (string): 
		remote_chains (dict[string, string]): A dictionary whose keys correspond to the chains 
			specified in the --remotes flag during deploy, and whose values are the RPC URLs where 
			those chains can be reached. 
	Returns:
		Returns a bunch of stuff.
`
	result, err := ParseRunFunctionDocstring(docstring)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestDocstringParser_EmptyDescription(t *testing.T) {
	docstring := `
    Args:
		origin_chain (string): The value of this field should be the same as the value provided to
			the --local flag when the Hyperlane contract was deployed.
		validator_key (string): 
		remote_chains (dict[string, string]): A dictionary whose keys correspond to the chains 
			specified in the --remotes flag during deploy, and whose values are the RPC URLs where 
			those chains can be reached. 
	Returns:
		Returns a bunch of stuff.
`
	result, err := ParseRunFunctionDocstring(docstring)
	require.NoError(t, err)
	require.NotNil(t, result)
}
