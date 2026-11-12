package ldap

import (
	"testing"

	"github.com/go-ldap/ldap/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestOUProtectionReadsRequestDACL(t *testing.T) {
	const base = "DC=lab,DC=example"
	const dn = "OU=Test," + base
	const guid = "11111111-2222-3333-4444-555555555555"

	client := &MockOUClient{}
	assertDACL := func(args mock.Arguments) {
		request := args.Get(1).(*SearchRequest)
		control := ldap.FindControl(request.Controls, ldap.ControlTypeMicrosoftSDFlags)
		require.NotNil(t, control)
		flags, ok := control.(*ldap.ControlMicrosoftSDFlags)
		require.True(t, ok)
		require.Equal(t, int32(SDFlagsDACLSecurityInformation), flags.ControlValue)
	}
	client.On("Search", mock.Anything, mock.Anything).Run(assertDACL).Return(&SearchResult{}, nil)
	client.On("SearchWithPaging", mock.Anything, mock.Anything).Run(assertDACL).Return(&SearchResult{}, nil)
	manager := NewOUManager(t.Context(), client, base)

	_, err := manager.GetOU(guid)
	require.True(t, IsNotFoundError(err))
	_, err = manager.GetOUByDN(dn)
	require.True(t, IsNotFoundError(err))
	_, err = manager.SearchOUs(base, "")
	require.NoError(t, err)
	_, err = manager.GetOUChildren(t.Context(), dn)
	require.NoError(t, err)
}
