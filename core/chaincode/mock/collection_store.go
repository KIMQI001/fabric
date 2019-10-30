// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/core/common/privdata"
	"github.com/hyperledger/fabric/core/ledger"
)

type CollectionStore struct {
	AccessFilterStub        func(string, *common.CollectionPolicyConfig) (privdata.Filter, error)
	accessFilterMutex       sync.RWMutex
	accessFilterArgsForCall []struct {
		arg1 string
		arg2 *common.CollectionPolicyConfig
	}
	accessFilterReturns struct {
		result1 privdata.Filter
		result2 error
	}
	accessFilterReturnsOnCall map[int]struct {
		result1 privdata.Filter
		result2 error
	}
	RetrieveCollectionStub        func(privdata.CollectionCriteria) (privdata.Collection, error)
	retrieveCollectionMutex       sync.RWMutex
	retrieveCollectionArgsForCall []struct {
		arg1 privdata.CollectionCriteria
	}
	retrieveCollectionReturns struct {
		result1 privdata.Collection
		result2 error
	}
	retrieveCollectionReturnsOnCall map[int]struct {
		result1 privdata.Collection
		result2 error
	}
	RetrieveCollectionAccessPolicyStub        func(privdata.CollectionCriteria) (privdata.CollectionAccessPolicy, error)
	retrieveCollectionAccessPolicyMutex       sync.RWMutex
	retrieveCollectionAccessPolicyArgsForCall []struct {
		arg1 privdata.CollectionCriteria
	}
	retrieveCollectionAccessPolicyReturns struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}
	retrieveCollectionAccessPolicyReturnsOnCall map[int]struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}
	RetrieveCollectionConfigStub        func(privdata.CollectionCriteria) (*common.StaticCollectionConfig, error)
	retrieveCollectionConfigMutex       sync.RWMutex
	retrieveCollectionConfigArgsForCall []struct {
		arg1 privdata.CollectionCriteria
	}
	retrieveCollectionConfigReturns struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}
	retrieveCollectionConfigReturnsOnCall map[int]struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}
	RetrieveCollectionConfigPackageStub        func(privdata.CollectionCriteria) (*common.CollectionConfigPackage, error)
	retrieveCollectionConfigPackageMutex       sync.RWMutex
	retrieveCollectionConfigPackageArgsForCall []struct {
		arg1 privdata.CollectionCriteria
	}
	retrieveCollectionConfigPackageReturns struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	retrieveCollectionConfigPackageReturnsOnCall map[int]struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	RetrieveCollectionPersistenceConfigsStub        func(privdata.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error)
	retrieveCollectionPersistenceConfigsMutex       sync.RWMutex
	retrieveCollectionPersistenceConfigsArgsForCall []struct {
		arg1 privdata.CollectionCriteria
	}
	retrieveCollectionPersistenceConfigsReturns struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}
	retrieveCollectionPersistenceConfigsReturnsOnCall map[int]struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}
	RetrieveReadWritePermissionStub        func(privdata.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) (bool, bool, error)
	retrieveReadWritePermissionMutex       sync.RWMutex
	retrieveReadWritePermissionArgsForCall []struct {
		arg1 privdata.CollectionCriteria
		arg2 *peer.SignedProposal
		arg3 ledger.QueryExecutor
	}
	retrieveReadWritePermissionReturns struct {
		result1 bool
		result2 bool
		result3 error
	}
	retrieveReadWritePermissionReturnsOnCall map[int]struct {
		result1 bool
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CollectionStore) AccessFilter(arg1 string, arg2 *common.CollectionPolicyConfig) (privdata.Filter, error) {
	fake.accessFilterMutex.Lock()
	ret, specificReturn := fake.accessFilterReturnsOnCall[len(fake.accessFilterArgsForCall)]
	fake.accessFilterArgsForCall = append(fake.accessFilterArgsForCall, struct {
		arg1 string
		arg2 *common.CollectionPolicyConfig
	}{arg1, arg2})
	fake.recordInvocation("AccessFilter", []interface{}{arg1, arg2})
	fake.accessFilterMutex.Unlock()
	if fake.AccessFilterStub != nil {
		return fake.AccessFilterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.accessFilterReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) AccessFilterCallCount() int {
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	return len(fake.accessFilterArgsForCall)
}

func (fake *CollectionStore) AccessFilterCalls(stub func(string, *common.CollectionPolicyConfig) (privdata.Filter, error)) {
	fake.accessFilterMutex.Lock()
	defer fake.accessFilterMutex.Unlock()
	fake.AccessFilterStub = stub
}

func (fake *CollectionStore) AccessFilterArgsForCall(i int) (string, *common.CollectionPolicyConfig) {
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	argsForCall := fake.accessFilterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CollectionStore) AccessFilterReturns(result1 privdata.Filter, result2 error) {
	fake.accessFilterMutex.Lock()
	defer fake.accessFilterMutex.Unlock()
	fake.AccessFilterStub = nil
	fake.accessFilterReturns = struct {
		result1 privdata.Filter
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) AccessFilterReturnsOnCall(i int, result1 privdata.Filter, result2 error) {
	fake.accessFilterMutex.Lock()
	defer fake.accessFilterMutex.Unlock()
	fake.AccessFilterStub = nil
	if fake.accessFilterReturnsOnCall == nil {
		fake.accessFilterReturnsOnCall = make(map[int]struct {
			result1 privdata.Filter
			result2 error
		})
	}
	fake.accessFilterReturnsOnCall[i] = struct {
		result1 privdata.Filter
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollection(arg1 privdata.CollectionCriteria) (privdata.Collection, error) {
	fake.retrieveCollectionMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionReturnsOnCall[len(fake.retrieveCollectionArgsForCall)]
	fake.retrieveCollectionArgsForCall = append(fake.retrieveCollectionArgsForCall, struct {
		arg1 privdata.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollection", []interface{}{arg1})
	fake.retrieveCollectionMutex.Unlock()
	if fake.RetrieveCollectionStub != nil {
		return fake.RetrieveCollectionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionCallCount() int {
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	return len(fake.retrieveCollectionArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionCalls(stub func(privdata.CollectionCriteria) (privdata.Collection, error)) {
	fake.retrieveCollectionMutex.Lock()
	defer fake.retrieveCollectionMutex.Unlock()
	fake.RetrieveCollectionStub = stub
}

func (fake *CollectionStore) RetrieveCollectionArgsForCall(i int) privdata.CollectionCriteria {
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	argsForCall := fake.retrieveCollectionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionReturns(result1 privdata.Collection, result2 error) {
	fake.retrieveCollectionMutex.Lock()
	defer fake.retrieveCollectionMutex.Unlock()
	fake.RetrieveCollectionStub = nil
	fake.retrieveCollectionReturns = struct {
		result1 privdata.Collection
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionReturnsOnCall(i int, result1 privdata.Collection, result2 error) {
	fake.retrieveCollectionMutex.Lock()
	defer fake.retrieveCollectionMutex.Unlock()
	fake.RetrieveCollectionStub = nil
	if fake.retrieveCollectionReturnsOnCall == nil {
		fake.retrieveCollectionReturnsOnCall = make(map[int]struct {
			result1 privdata.Collection
			result2 error
		})
	}
	fake.retrieveCollectionReturnsOnCall[i] = struct {
		result1 privdata.Collection
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicy(arg1 privdata.CollectionCriteria) (privdata.CollectionAccessPolicy, error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionAccessPolicyReturnsOnCall[len(fake.retrieveCollectionAccessPolicyArgsForCall)]
	fake.retrieveCollectionAccessPolicyArgsForCall = append(fake.retrieveCollectionAccessPolicyArgsForCall, struct {
		arg1 privdata.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionAccessPolicy", []interface{}{arg1})
	fake.retrieveCollectionAccessPolicyMutex.Unlock()
	if fake.RetrieveCollectionAccessPolicyStub != nil {
		return fake.RetrieveCollectionAccessPolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionAccessPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyCallCount() int {
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	return len(fake.retrieveCollectionAccessPolicyArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyCalls(stub func(privdata.CollectionCriteria) (privdata.CollectionAccessPolicy, error)) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	defer fake.retrieveCollectionAccessPolicyMutex.Unlock()
	fake.RetrieveCollectionAccessPolicyStub = stub
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyArgsForCall(i int) privdata.CollectionCriteria {
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	argsForCall := fake.retrieveCollectionAccessPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyReturns(result1 privdata.CollectionAccessPolicy, result2 error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	defer fake.retrieveCollectionAccessPolicyMutex.Unlock()
	fake.RetrieveCollectionAccessPolicyStub = nil
	fake.retrieveCollectionAccessPolicyReturns = struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyReturnsOnCall(i int, result1 privdata.CollectionAccessPolicy, result2 error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	defer fake.retrieveCollectionAccessPolicyMutex.Unlock()
	fake.RetrieveCollectionAccessPolicyStub = nil
	if fake.retrieveCollectionAccessPolicyReturnsOnCall == nil {
		fake.retrieveCollectionAccessPolicyReturnsOnCall = make(map[int]struct {
			result1 privdata.CollectionAccessPolicy
			result2 error
		})
	}
	fake.retrieveCollectionAccessPolicyReturnsOnCall[i] = struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfig(arg1 privdata.CollectionCriteria) (*common.StaticCollectionConfig, error) {
	fake.retrieveCollectionConfigMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionConfigReturnsOnCall[len(fake.retrieveCollectionConfigArgsForCall)]
	fake.retrieveCollectionConfigArgsForCall = append(fake.retrieveCollectionConfigArgsForCall, struct {
		arg1 privdata.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionConfig", []interface{}{arg1})
	fake.retrieveCollectionConfigMutex.Unlock()
	if fake.RetrieveCollectionConfigStub != nil {
		return fake.RetrieveCollectionConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionConfigCallCount() int {
	fake.retrieveCollectionConfigMutex.RLock()
	defer fake.retrieveCollectionConfigMutex.RUnlock()
	return len(fake.retrieveCollectionConfigArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionConfigCalls(stub func(privdata.CollectionCriteria) (*common.StaticCollectionConfig, error)) {
	fake.retrieveCollectionConfigMutex.Lock()
	defer fake.retrieveCollectionConfigMutex.Unlock()
	fake.RetrieveCollectionConfigStub = stub
}

func (fake *CollectionStore) RetrieveCollectionConfigArgsForCall(i int) privdata.CollectionCriteria {
	fake.retrieveCollectionConfigMutex.RLock()
	defer fake.retrieveCollectionConfigMutex.RUnlock()
	argsForCall := fake.retrieveCollectionConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionConfigReturns(result1 *common.StaticCollectionConfig, result2 error) {
	fake.retrieveCollectionConfigMutex.Lock()
	defer fake.retrieveCollectionConfigMutex.Unlock()
	fake.RetrieveCollectionConfigStub = nil
	fake.retrieveCollectionConfigReturns = struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigReturnsOnCall(i int, result1 *common.StaticCollectionConfig, result2 error) {
	fake.retrieveCollectionConfigMutex.Lock()
	defer fake.retrieveCollectionConfigMutex.Unlock()
	fake.RetrieveCollectionConfigStub = nil
	if fake.retrieveCollectionConfigReturnsOnCall == nil {
		fake.retrieveCollectionConfigReturnsOnCall = make(map[int]struct {
			result1 *common.StaticCollectionConfig
			result2 error
		})
	}
	fake.retrieveCollectionConfigReturnsOnCall[i] = struct {
		result1 *common.StaticCollectionConfig
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigPackage(arg1 privdata.CollectionCriteria) (*common.CollectionConfigPackage, error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionConfigPackageReturnsOnCall[len(fake.retrieveCollectionConfigPackageArgsForCall)]
	fake.retrieveCollectionConfigPackageArgsForCall = append(fake.retrieveCollectionConfigPackageArgsForCall, struct {
		arg1 privdata.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionConfigPackage", []interface{}{arg1})
	fake.retrieveCollectionConfigPackageMutex.Unlock()
	if fake.RetrieveCollectionConfigPackageStub != nil {
		return fake.RetrieveCollectionConfigPackageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionConfigPackageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageCallCount() int {
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	return len(fake.retrieveCollectionConfigPackageArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageCalls(stub func(privdata.CollectionCriteria) (*common.CollectionConfigPackage, error)) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	defer fake.retrieveCollectionConfigPackageMutex.Unlock()
	fake.RetrieveCollectionConfigPackageStub = stub
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageArgsForCall(i int) privdata.CollectionCriteria {
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	argsForCall := fake.retrieveCollectionConfigPackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageReturns(result1 *common.CollectionConfigPackage, result2 error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	defer fake.retrieveCollectionConfigPackageMutex.Unlock()
	fake.RetrieveCollectionConfigPackageStub = nil
	fake.retrieveCollectionConfigPackageReturns = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageReturnsOnCall(i int, result1 *common.CollectionConfigPackage, result2 error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	defer fake.retrieveCollectionConfigPackageMutex.Unlock()
	fake.RetrieveCollectionConfigPackageStub = nil
	if fake.retrieveCollectionConfigPackageReturnsOnCall == nil {
		fake.retrieveCollectionConfigPackageReturnsOnCall = make(map[int]struct {
			result1 *common.CollectionConfigPackage
			result2 error
		})
	}
	fake.retrieveCollectionConfigPackageReturnsOnCall[i] = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigs(arg1 privdata.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionPersistenceConfigsReturnsOnCall[len(fake.retrieveCollectionPersistenceConfigsArgsForCall)]
	fake.retrieveCollectionPersistenceConfigsArgsForCall = append(fake.retrieveCollectionPersistenceConfigsArgsForCall, struct {
		arg1 privdata.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionPersistenceConfigs", []interface{}{arg1})
	fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	if fake.RetrieveCollectionPersistenceConfigsStub != nil {
		return fake.RetrieveCollectionPersistenceConfigsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.retrieveCollectionPersistenceConfigsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsCallCount() int {
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	return len(fake.retrieveCollectionPersistenceConfigsArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsCalls(stub func(privdata.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error)) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	fake.RetrieveCollectionPersistenceConfigsStub = stub
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsArgsForCall(i int) privdata.CollectionCriteria {
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	argsForCall := fake.retrieveCollectionPersistenceConfigsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsReturns(result1 privdata.CollectionPersistenceConfigs, result2 error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	fake.RetrieveCollectionPersistenceConfigsStub = nil
	fake.retrieveCollectionPersistenceConfigsReturns = struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsReturnsOnCall(i int, result1 privdata.CollectionPersistenceConfigs, result2 error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	fake.RetrieveCollectionPersistenceConfigsStub = nil
	if fake.retrieveCollectionPersistenceConfigsReturnsOnCall == nil {
		fake.retrieveCollectionPersistenceConfigsReturnsOnCall = make(map[int]struct {
			result1 privdata.CollectionPersistenceConfigs
			result2 error
		})
	}
	fake.retrieveCollectionPersistenceConfigsReturnsOnCall[i] = struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveReadWritePermission(arg1 privdata.CollectionCriteria, arg2 *peer.SignedProposal, arg3 ledger.QueryExecutor) (bool, bool, error) {
	fake.retrieveReadWritePermissionMutex.Lock()
	ret, specificReturn := fake.retrieveReadWritePermissionReturnsOnCall[len(fake.retrieveReadWritePermissionArgsForCall)]
	fake.retrieveReadWritePermissionArgsForCall = append(fake.retrieveReadWritePermissionArgsForCall, struct {
		arg1 privdata.CollectionCriteria
		arg2 *peer.SignedProposal
		arg3 ledger.QueryExecutor
	}{arg1, arg2, arg3})
	fake.recordInvocation("RetrieveReadWritePermission", []interface{}{arg1, arg2, arg3})
	fake.retrieveReadWritePermissionMutex.Unlock()
	if fake.RetrieveReadWritePermissionStub != nil {
		return fake.RetrieveReadWritePermissionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.retrieveReadWritePermissionReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *CollectionStore) RetrieveReadWritePermissionCallCount() int {
	fake.retrieveReadWritePermissionMutex.RLock()
	defer fake.retrieveReadWritePermissionMutex.RUnlock()
	return len(fake.retrieveReadWritePermissionArgsForCall)
}

func (fake *CollectionStore) RetrieveReadWritePermissionCalls(stub func(privdata.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) (bool, bool, error)) {
	fake.retrieveReadWritePermissionMutex.Lock()
	defer fake.retrieveReadWritePermissionMutex.Unlock()
	fake.RetrieveReadWritePermissionStub = stub
}

func (fake *CollectionStore) RetrieveReadWritePermissionArgsForCall(i int) (privdata.CollectionCriteria, *peer.SignedProposal, ledger.QueryExecutor) {
	fake.retrieveReadWritePermissionMutex.RLock()
	defer fake.retrieveReadWritePermissionMutex.RUnlock()
	argsForCall := fake.retrieveReadWritePermissionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CollectionStore) RetrieveReadWritePermissionReturns(result1 bool, result2 bool, result3 error) {
	fake.retrieveReadWritePermissionMutex.Lock()
	defer fake.retrieveReadWritePermissionMutex.Unlock()
	fake.RetrieveReadWritePermissionStub = nil
	fake.retrieveReadWritePermissionReturns = struct {
		result1 bool
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *CollectionStore) RetrieveReadWritePermissionReturnsOnCall(i int, result1 bool, result2 bool, result3 error) {
	fake.retrieveReadWritePermissionMutex.Lock()
	defer fake.retrieveReadWritePermissionMutex.Unlock()
	fake.RetrieveReadWritePermissionStub = nil
	if fake.retrieveReadWritePermissionReturnsOnCall == nil {
		fake.retrieveReadWritePermissionReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 bool
			result3 error
		})
	}
	fake.retrieveReadWritePermissionReturnsOnCall[i] = struct {
		result1 bool
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *CollectionStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	fake.retrieveCollectionConfigMutex.RLock()
	defer fake.retrieveCollectionConfigMutex.RUnlock()
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	fake.retrieveReadWritePermissionMutex.RLock()
	defer fake.retrieveReadWritePermissionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CollectionStore) recordInvocation(key string, args []interface{}) {
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

var _ privdata.CollectionStore = new(CollectionStore)
