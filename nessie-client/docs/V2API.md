# \V2API

All URIs are relative to *http://localhost:19120*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReferenceV2**](V2API.md#CreateReferenceV2) | **Post** /v2/trees | Create a new branch or tag
[**GetAllReferencesV2**](V2API.md#GetAllReferencesV2) | **Get** /v2/trees | Get information about all branches and tags



## CreateReferenceV2

> SingleReferenceResponse1 CreateReferenceV2(ctx).Name(name).Type_(type_).Reference3(reference3).Execute()

Create a new branch or tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | A reference name.  Reference name must start with a letter, followed by letters, digits, one of the ./_- characters, not end with a slash or dot, not contain '..' 
	type_ := "branch" // string | Type of the reference to be created
	reference3 := openapiclient.Reference_3{Branch3: openapiclient.NewBranch3("Name_example")} // Reference3 | Source reference data from which the new reference is to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.CreateReferenceV2(context.Background()).Name(name).Type_(type_).Reference3(reference3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.CreateReferenceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReferenceV2`: SingleReferenceResponse1
	fmt.Fprintf(os.Stdout, "Response from `V2API.CreateReferenceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferenceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | A reference name.  Reference name must start with a letter, followed by letters, digits, one of the ./_- characters, not end with a slash or dot, not contain &#39;..&#39;  | 
 **type_** | **string** | Type of the reference to be created | 
 **reference3** | [**Reference3**](Reference3.md) | Source reference data from which the new reference is to be created. | 

### Return type

[**SingleReferenceResponse1**](SingleReferenceResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReferencesV2

> GetAllReferencesV2200Response GetAllReferencesV2(ctx).Fetch(fetch).Filter(filter).MaxRecords(maxRecords).PageToken(pageToken).Execute()

Get information about all branches and tags

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fetch := "fetch_example" // string | Specifies how much extra information is to be retrived from the server.  If the fetch option is set to 'ALL' the following addition information will be returned for each Branch object in the output:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits from the root to the HEAD of the branch).  The returned Tag instances will only contain the 'commitMetaOfHEAD' and 'numTotalCommits' fields.  Note that computing & fetching additional metadata might be computationally expensive on the server-side, so this flag should be used with care. (optional)
	filter := "refType == 'BRANCH'" // string | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md. Usable variables within the expression are:  - ref (Reference) describes the reference, with fields name (String), hash (String), metadata (ReferenceMetadata)  - metadata (ReferenceMetadata) shortcut to ref.metadata, never null, but possibly empty  - commit (CommitMeta) - shortcut to ref.metadata.commitMetaOfHEAD, never null, but possibly empty  - refType (String) - the reference type, either BRANCH or TAG  Note that the expression can only test attributes metadata and commit, if 'fetchOption' is set to 'ALL'. (optional)
	maxRecords := int32(56) // int32 | maximum number of entries to return, just a hint for the server (optional)
	pageToken := "pageToken_example" // string | paging continuation token, as returned in the previous value of the field 'token' in the corresponding 'EntriesResponse' or 'LogResponse' or 'ReferencesResponse' or 'RefLogResponse'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.GetAllReferencesV2(context.Background()).Fetch(fetch).Filter(filter).MaxRecords(maxRecords).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.GetAllReferencesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReferencesV2`: GetAllReferencesV2200Response
	fmt.Fprintf(os.Stdout, "Response from `V2API.GetAllReferencesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReferencesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetch** | **string** | Specifies how much extra information is to be retrived from the server.  If the fetch option is set to &#39;ALL&#39; the following addition information will be returned for each Branch object in the output:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits from the root to the HEAD of the branch).  The returned Tag instances will only contain the &#39;commitMetaOfHEAD&#39; and &#39;numTotalCommits&#39; fields.  Note that computing &amp; fetching additional metadata might be computationally expensive on the server-side, so this flag should be used with care. | 
 **filter** | **string** | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md. Usable variables within the expression are:  - ref (Reference) describes the reference, with fields name (String), hash (String), metadata (ReferenceMetadata)  - metadata (ReferenceMetadata) shortcut to ref.metadata, never null, but possibly empty  - commit (CommitMeta) - shortcut to ref.metadata.commitMetaOfHEAD, never null, but possibly empty  - refType (String) - the reference type, either BRANCH or TAG  Note that the expression can only test attributes metadata and commit, if &#39;fetchOption&#39; is set to &#39;ALL&#39;. | 
 **maxRecords** | **int32** | maximum number of entries to return, just a hint for the server | 
 **pageToken** | **string** | paging continuation token, as returned in the previous value of the field &#39;token&#39; in the corresponding &#39;EntriesResponse&#39; or &#39;LogResponse&#39; or &#39;ReferencesResponse&#39; or &#39;RefLogResponse&#39;. | 

### Return type

[**GetAllReferencesV2200Response**](GetAllReferencesV2200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

