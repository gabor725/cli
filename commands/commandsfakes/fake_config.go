// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/utils/configv3"
)

type FakeConfig struct {
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct{}
	aPIVersionReturns     struct {
		result1 string
	}
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
	}
	BinaryNameStub        func() string
	binaryNameMutex       sync.RWMutex
	binaryNameArgsForCall []struct{}
	binaryNameReturns     struct {
		result1 string
	}
	ColorEnabledStub        func() configv3.ColorSetting
	colorEnabledMutex       sync.RWMutex
	colorEnabledArgsForCall []struct{}
	colorEnabledReturns     struct {
		result1 configv3.ColorSetting
	}
	CurrentUserStub        func() (configv3.User, error)
	currentUserMutex       sync.RWMutex
	currentUserArgsForCall []struct{}
	currentUserReturns     struct {
		result1 configv3.User
		result2 error
	}
	ExperimentalStub        func() bool
	experimentalMutex       sync.RWMutex
	experimentalArgsForCall []struct{}
	experimentalReturns     struct {
		result1 bool
	}
	LocaleStub        func() string
	localeMutex       sync.RWMutex
	localeArgsForCall []struct{}
	localeReturns     struct {
		result1 string
	}
	PluginsStub        func() map[string]configv3.Plugin
	pluginsMutex       sync.RWMutex
	pluginsArgsForCall []struct{}
	pluginsReturns     struct {
		result1 map[string]configv3.Plugin
	}
	RefreshTokenStub        func() string
	refreshTokenMutex       sync.RWMutex
	refreshTokenArgsForCall []struct{}
	refreshTokenReturns     struct {
		result1 string
	}
	SetTargetInformationStub        func(api string, apiVersion string, auth string, loggregator string, doppler string, uaa string, routing string, skipSSLValidation bool)
	setTargetInformationMutex       sync.RWMutex
	setTargetInformationArgsForCall []struct {
		api               string
		apiVersion        string
		auth              string
		loggregator       string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}
	SetTokenInformationStub        func(accessToken string, refreshToken string, sshOAuthClient string)
	setTokenInformationMutex       sync.RWMutex
	setTokenInformationArgsForCall []struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}
	SkipSSLValidationStub        func() bool
	skipSSLValidationMutex       sync.RWMutex
	skipSSLValidationArgsForCall []struct{}
	skipSSLValidationReturns     struct {
		result1 bool
	}
	TargetStub        func() string
	targetMutex       sync.RWMutex
	targetArgsForCall []struct{}
	targetReturns     struct {
		result1 string
	}
	TargetedOrganizationStub        func() configv3.Organization
	targetedOrganizationMutex       sync.RWMutex
	targetedOrganizationArgsForCall []struct{}
	targetedOrganizationReturns     struct {
		result1 configv3.Organization
	}
	TargetedSpaceStub        func() configv3.Space
	targetedSpaceMutex       sync.RWMutex
	targetedSpaceArgsForCall []struct{}
	targetedSpaceReturns     struct {
		result1 configv3.Space
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("APIVersion", []interface{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	} else {
		return fake.aPIVersionReturns.result1
	}
}

func (fake *FakeConfig) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakeConfig) APIVersionReturns(result1 string) {
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) AccessToken() string {
	fake.accessTokenMutex.Lock()
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	} else {
		return fake.accessTokenReturns.result1
	}
}

func (fake *FakeConfig) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeConfig) AccessTokenReturns(result1 string) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) BinaryName() string {
	fake.binaryNameMutex.Lock()
	fake.binaryNameArgsForCall = append(fake.binaryNameArgsForCall, struct{}{})
	fake.recordInvocation("BinaryName", []interface{}{})
	fake.binaryNameMutex.Unlock()
	if fake.BinaryNameStub != nil {
		return fake.BinaryNameStub()
	} else {
		return fake.binaryNameReturns.result1
	}
}

func (fake *FakeConfig) BinaryNameCallCount() int {
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	return len(fake.binaryNameArgsForCall)
}

func (fake *FakeConfig) BinaryNameReturns(result1 string) {
	fake.BinaryNameStub = nil
	fake.binaryNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) ColorEnabled() configv3.ColorSetting {
	fake.colorEnabledMutex.Lock()
	fake.colorEnabledArgsForCall = append(fake.colorEnabledArgsForCall, struct{}{})
	fake.recordInvocation("ColorEnabled", []interface{}{})
	fake.colorEnabledMutex.Unlock()
	if fake.ColorEnabledStub != nil {
		return fake.ColorEnabledStub()
	} else {
		return fake.colorEnabledReturns.result1
	}
}

func (fake *FakeConfig) ColorEnabledCallCount() int {
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	return len(fake.colorEnabledArgsForCall)
}

func (fake *FakeConfig) ColorEnabledReturns(result1 configv3.ColorSetting) {
	fake.ColorEnabledStub = nil
	fake.colorEnabledReturns = struct {
		result1 configv3.ColorSetting
	}{result1}
}

func (fake *FakeConfig) CurrentUser() (configv3.User, error) {
	fake.currentUserMutex.Lock()
	fake.currentUserArgsForCall = append(fake.currentUserArgsForCall, struct{}{})
	fake.recordInvocation("CurrentUser", []interface{}{})
	fake.currentUserMutex.Unlock()
	if fake.CurrentUserStub != nil {
		return fake.CurrentUserStub()
	} else {
		return fake.currentUserReturns.result1, fake.currentUserReturns.result2
	}
}

func (fake *FakeConfig) CurrentUserCallCount() int {
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	return len(fake.currentUserArgsForCall)
}

func (fake *FakeConfig) CurrentUserReturns(result1 configv3.User, result2 error) {
	fake.CurrentUserStub = nil
	fake.currentUserReturns = struct {
		result1 configv3.User
		result2 error
	}{result1, result2}
}

func (fake *FakeConfig) Experimental() bool {
	fake.experimentalMutex.Lock()
	fake.experimentalArgsForCall = append(fake.experimentalArgsForCall, struct{}{})
	fake.recordInvocation("Experimental", []interface{}{})
	fake.experimentalMutex.Unlock()
	if fake.ExperimentalStub != nil {
		return fake.ExperimentalStub()
	} else {
		return fake.experimentalReturns.result1
	}
}

func (fake *FakeConfig) ExperimentalCallCount() int {
	fake.experimentalMutex.RLock()
	defer fake.experimentalMutex.RUnlock()
	return len(fake.experimentalArgsForCall)
}

func (fake *FakeConfig) ExperimentalReturns(result1 bool) {
	fake.ExperimentalStub = nil
	fake.experimentalReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) Locale() string {
	fake.localeMutex.Lock()
	fake.localeArgsForCall = append(fake.localeArgsForCall, struct{}{})
	fake.recordInvocation("Locale", []interface{}{})
	fake.localeMutex.Unlock()
	if fake.LocaleStub != nil {
		return fake.LocaleStub()
	} else {
		return fake.localeReturns.result1
	}
}

func (fake *FakeConfig) LocaleCallCount() int {
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	return len(fake.localeArgsForCall)
}

func (fake *FakeConfig) LocaleReturns(result1 string) {
	fake.LocaleStub = nil
	fake.localeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Plugins() map[string]configv3.Plugin {
	fake.pluginsMutex.Lock()
	fake.pluginsArgsForCall = append(fake.pluginsArgsForCall, struct{}{})
	fake.recordInvocation("Plugins", []interface{}{})
	fake.pluginsMutex.Unlock()
	if fake.PluginsStub != nil {
		return fake.PluginsStub()
	} else {
		return fake.pluginsReturns.result1
	}
}

func (fake *FakeConfig) PluginsCallCount() int {
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	return len(fake.pluginsArgsForCall)
}

func (fake *FakeConfig) PluginsReturns(result1 map[string]configv3.Plugin) {
	fake.PluginsStub = nil
	fake.pluginsReturns = struct {
		result1 map[string]configv3.Plugin
	}{result1}
}

func (fake *FakeConfig) RefreshToken() string {
	fake.refreshTokenMutex.Lock()
	fake.refreshTokenArgsForCall = append(fake.refreshTokenArgsForCall, struct{}{})
	fake.recordInvocation("RefreshToken", []interface{}{})
	fake.refreshTokenMutex.Unlock()
	if fake.RefreshTokenStub != nil {
		return fake.RefreshTokenStub()
	} else {
		return fake.refreshTokenReturns.result1
	}
}

func (fake *FakeConfig) RefreshTokenCallCount() int {
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	return len(fake.refreshTokenArgsForCall)
}

func (fake *FakeConfig) RefreshTokenReturns(result1 string) {
	fake.RefreshTokenStub = nil
	fake.refreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) SetTargetInformation(api string, apiVersion string, auth string, loggregator string, doppler string, uaa string, routing string, skipSSLValidation bool) {
	fake.setTargetInformationMutex.Lock()
	fake.setTargetInformationArgsForCall = append(fake.setTargetInformationArgsForCall, struct {
		api               string
		apiVersion        string
		auth              string
		loggregator       string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}{api, apiVersion, auth, loggregator, doppler, uaa, routing, skipSSLValidation})
	fake.recordInvocation("SetTargetInformation", []interface{}{api, apiVersion, auth, loggregator, doppler, uaa, routing, skipSSLValidation})
	fake.setTargetInformationMutex.Unlock()
	if fake.SetTargetInformationStub != nil {
		fake.SetTargetInformationStub(api, apiVersion, auth, loggregator, doppler, uaa, routing, skipSSLValidation)
	}
}

func (fake *FakeConfig) SetTargetInformationCallCount() int {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return len(fake.setTargetInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformationArgsForCall(i int) (string, string, string, string, string, string, string, bool) {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return fake.setTargetInformationArgsForCall[i].api, fake.setTargetInformationArgsForCall[i].apiVersion, fake.setTargetInformationArgsForCall[i].auth, fake.setTargetInformationArgsForCall[i].loggregator, fake.setTargetInformationArgsForCall[i].doppler, fake.setTargetInformationArgsForCall[i].uaa, fake.setTargetInformationArgsForCall[i].routing, fake.setTargetInformationArgsForCall[i].skipSSLValidation
}

func (fake *FakeConfig) SetTokenInformation(accessToken string, refreshToken string, sshOAuthClient string) {
	fake.setTokenInformationMutex.Lock()
	fake.setTokenInformationArgsForCall = append(fake.setTokenInformationArgsForCall, struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}{accessToken, refreshToken, sshOAuthClient})
	fake.recordInvocation("SetTokenInformation", []interface{}{accessToken, refreshToken, sshOAuthClient})
	fake.setTokenInformationMutex.Unlock()
	if fake.SetTokenInformationStub != nil {
		fake.SetTokenInformationStub(accessToken, refreshToken, sshOAuthClient)
	}
}

func (fake *FakeConfig) SetTokenInformationCallCount() int {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return len(fake.setTokenInformationArgsForCall)
}

func (fake *FakeConfig) SetTokenInformationArgsForCall(i int) (string, string, string) {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return fake.setTokenInformationArgsForCall[i].accessToken, fake.setTokenInformationArgsForCall[i].refreshToken, fake.setTokenInformationArgsForCall[i].sshOAuthClient
}

func (fake *FakeConfig) SkipSSLValidation() bool {
	fake.skipSSLValidationMutex.Lock()
	fake.skipSSLValidationArgsForCall = append(fake.skipSSLValidationArgsForCall, struct{}{})
	fake.recordInvocation("SkipSSLValidation", []interface{}{})
	fake.skipSSLValidationMutex.Unlock()
	if fake.SkipSSLValidationStub != nil {
		return fake.SkipSSLValidationStub()
	} else {
		return fake.skipSSLValidationReturns.result1
	}
}

func (fake *FakeConfig) SkipSSLValidationCallCount() int {
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return len(fake.skipSSLValidationArgsForCall)
}

func (fake *FakeConfig) SkipSSLValidationReturns(result1 bool) {
	fake.SkipSSLValidationStub = nil
	fake.skipSSLValidationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) Target() string {
	fake.targetMutex.Lock()
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct{}{})
	fake.recordInvocation("Target", []interface{}{})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub()
	} else {
		return fake.targetReturns.result1
	}
}

func (fake *FakeConfig) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeConfig) TargetReturns(result1 string) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) TargetedOrganization() configv3.Organization {
	fake.targetedOrganizationMutex.Lock()
	fake.targetedOrganizationArgsForCall = append(fake.targetedOrganizationArgsForCall, struct{}{})
	fake.recordInvocation("TargetedOrganization", []interface{}{})
	fake.targetedOrganizationMutex.Unlock()
	if fake.TargetedOrganizationStub != nil {
		return fake.TargetedOrganizationStub()
	} else {
		return fake.targetedOrganizationReturns.result1
	}
}

func (fake *FakeConfig) TargetedOrganizationCallCount() int {
	fake.targetedOrganizationMutex.RLock()
	defer fake.targetedOrganizationMutex.RUnlock()
	return len(fake.targetedOrganizationArgsForCall)
}

func (fake *FakeConfig) TargetedOrganizationReturns(result1 configv3.Organization) {
	fake.TargetedOrganizationStub = nil
	fake.targetedOrganizationReturns = struct {
		result1 configv3.Organization
	}{result1}
}

func (fake *FakeConfig) TargetedSpace() configv3.Space {
	fake.targetedSpaceMutex.Lock()
	fake.targetedSpaceArgsForCall = append(fake.targetedSpaceArgsForCall, struct{}{})
	fake.recordInvocation("TargetedSpace", []interface{}{})
	fake.targetedSpaceMutex.Unlock()
	if fake.TargetedSpaceStub != nil {
		return fake.TargetedSpaceStub()
	} else {
		return fake.targetedSpaceReturns.result1
	}
}

func (fake *FakeConfig) TargetedSpaceCallCount() int {
	fake.targetedSpaceMutex.RLock()
	defer fake.targetedSpaceMutex.RUnlock()
	return len(fake.targetedSpaceArgsForCall)
}

func (fake *FakeConfig) TargetedSpaceReturns(result1 configv3.Space) {
	fake.TargetedSpaceStub = nil
	fake.targetedSpaceReturns = struct {
		result1 configv3.Space
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	fake.currentUserMutex.RLock()
	defer fake.currentUserMutex.RUnlock()
	fake.experimentalMutex.RLock()
	defer fake.experimentalMutex.RUnlock()
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	fake.targetedOrganizationMutex.RLock()
	defer fake.targetedOrganizationMutex.RUnlock()
	fake.targetedSpaceMutex.RLock()
	defer fake.targetedSpaceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ commands.Config = new(FakeConfig)
