# \V1API

All URIs are relative to *http://localhost:19120*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignReference**](V1API.md#AssignReference) | **Put** /v1/trees/{referenceType}/{referenceName} | Set a named reference to a specific hash via a named-reference.
[**CommitMultipleOperations**](V1API.md#CommitMultipleOperations) | **Post** /v1/trees/branch/{branchName}/commit | Commit multiple operations against the given branch expecting that branch to have the given hash as its latest commit. The hash in the successful response contains the hash of the commit that contains the operations of the invocation.
[**CreateNamespace**](V1API.md#CreateNamespace) | **Put** /v1/namespaces/namespace/{ref}/{name} | Creates a Namespace
[**CreateReference**](V1API.md#CreateReference) | **Post** /v1/trees/tree | Create a new reference
[**DeleteNamespace**](V1API.md#DeleteNamespace) | **Delete** /v1/namespaces/namespace/{ref}/{name} | Deletes a Namespace
[**DeleteReference**](V1API.md#DeleteReference) | **Delete** /v1/trees/{referenceType}/{referenceName} | Delete a reference endpoint
[**GetAllReferences**](V1API.md#GetAllReferences) | **Get** /v1/trees | Get all references
[**GetCommitLog**](V1API.md#GetCommitLog) | **Get** /v1/trees/tree/{ref}/log | Get commit log for a reference
[**GetConfig**](V1API.md#GetConfig) | **Get** /v1/config | List all configuration settings
[**GetContent**](V1API.md#GetContent) | **Get** /v1/contents/{key} | Get object content associated with a key.
[**GetDefaultBranch**](V1API.md#GetDefaultBranch) | **Get** /v1/trees/tree | Get default branch for commits and reads
[**GetDiff**](V1API.md#GetDiff) | **Get** /v1/diffs/{fromRefWithHash}...{toRefWithHash} | Get a diff for two given references
[**GetEntries**](V1API.md#GetEntries) | **Get** /v1/trees/tree/{ref}/entries | Fetch all entries for a given reference
[**GetMultipleContents**](V1API.md#GetMultipleContents) | **Post** /v1/contents | Get multiple objects&#39; content.
[**GetNamespace**](V1API.md#GetNamespace) | **Get** /v1/namespaces/namespace/{ref}/{name} | Retrieves a Namespace
[**GetNamespaces**](V1API.md#GetNamespaces) | **Get** /v1/namespaces/{ref} | 
[**GetReferenceByName**](V1API.md#GetReferenceByName) | **Get** /v1/trees/tree/{ref} | Fetch details of a reference
[**MergeRefIntoBranch**](V1API.md#MergeRefIntoBranch) | **Post** /v1/trees/branch/{branchName}/merge | Merge commits from &#39;mergeRef&#39; onto &#39;branchName&#39;.
[**TransplantCommitsIntoBranch**](V1API.md#TransplantCommitsIntoBranch) | **Post** /v1/trees/branch/{branchName}/transplant | Transplant commits from &#39;transplant&#39; onto &#39;branchName&#39;
[**UpdateProperties**](V1API.md#UpdateProperties) | **Post** /v1/namespaces/namespace/{ref}/{name} | 



## AssignReference

> AssignReference(ctx, referenceName, referenceType).ExpectedHash(expectedHash).Reference2(reference2).Execute()

Set a named reference to a specific hash via a named-reference.



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
	referenceName := "main" // string | Reference name to reassign
	referenceType := "branch" // string | Reference type to reassign
	expectedHash := "1122334455667788112233445566778811223344556677881122334455667788" // string | Expected previous hash of reference
	reference2 := openapiclient.Reference_2{Branch2: openapiclient.NewBranch2("Name_example")} // Reference2 | Reference hash to which 'referenceName' shall be assigned to. This must be either a 'Transaction', 'Branch' or 'Tag' via which the hash is visible to the caller.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V1API.AssignReference(context.Background(), referenceName, referenceType).ExpectedHash(expectedHash).Reference2(reference2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AssignReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceName** | **string** | Reference name to reassign | 
**referenceType** | **string** | Reference type to reassign | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expectedHash** | **string** | Expected previous hash of reference | 
 **reference2** | [**Reference2**](Reference2.md) | Reference hash to which &#39;referenceName&#39; shall be assigned to. This must be either a &#39;Transaction&#39;, &#39;Branch&#39; or &#39;Tag&#39; via which the hash is visible to the caller. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommitMultipleOperations

> Branch1 CommitMultipleOperations(ctx, branchName).Operations1(operations1).ExpectedHash(expectedHash).Execute()

Commit multiple operations against the given branch expecting that branch to have the given hash as its latest commit. The hash in the successful response contains the hash of the commit that contains the operations of the invocation.

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
	branchName := "main" // string | Branch to change, defaults to default branch.
	operations1 := *openapiclient.NewOperations1(*openapiclient.NewCommitMeta2([]string{"Authors_example"}, []string{"AllSignedOffBy_example"}, "Message_example", map[string]string{"key": "Inner_example"}, map[string][]string{"key": []string{"Inner_example"}}, []string{"ParentCommitHashes_example"}), []openapiclient.Operation1{openapiclient.Operation_1{DeleteContentOperationForAContentKey: openapiclient.NewDeleteContentOperationForAContentKey(*openapiclient.NewGetMultipleContentsRequest1RequestedKeysInner([]string{"Elements_example"}))}}) // Operations1 | Operations
	expectedHash := "1122334455667788112233445566778811223344556677881122334455667788" // string | Expected hash of branch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CommitMultipleOperations(context.Background(), branchName).Operations1(operations1).ExpectedHash(expectedHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CommitMultipleOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommitMultipleOperations`: Branch1
	fmt.Fprintf(os.Stdout, "Response from `V1API.CommitMultipleOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchName** | **string** | Branch to change, defaults to default branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitMultipleOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operations1** | [**Operations1**](Operations1.md) | Operations | 
 **expectedHash** | **string** | Expected hash of branch. | 

### Return type

[**Branch1**](Branch1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespace

> Content1AnyOf2 CreateNamespace(ctx, name, ref).Content1AnyOf2(content1AnyOf2).HashOnRef(hashOnRef).Execute()

Creates a Namespace

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
	name := *openapiclient.NewContent1AnyOf2([]string{"Elements_example"}, map[string]string{"key": "Inner_example"}) // Content1AnyOf2 | the name of the namespace
	ref := "main" // string | name of ref to fetch
	content1AnyOf2 := *openapiclient.NewContent1AnyOf2([]string{"Elements_example"}, map[string]string{"key": "Inner_example"}) // Content1AnyOf2 | 
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateNamespace(context.Background(), name, ref).Content1AnyOf2(content1AnyOf2).HashOnRef(hashOnRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespace`: Content1AnyOf2
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**Content1AnyOf2**](.md) | the name of the namespace | 
**ref** | **string** | name of ref to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content1AnyOf2** | [**Content1AnyOf2**](Content1AnyOf2.md) |  | 
 **hashOnRef** | **string** | a particular hash on the given ref | 

### Return type

[**Content1AnyOf2**](Content1AnyOf2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReference

> Reference1 CreateReference(ctx).Reference2(reference2).SourceRefName(sourceRefName).Execute()

Create a new reference



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
	reference2 := openapiclient.Reference_2{Branch2: openapiclient.NewBranch2("Name_example")} // Reference2 | Reference to create.
	sourceRefName := "sourceRefName_example" // string | Source named reference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.CreateReference(context.Background()).Reference2(reference2).SourceRefName(sourceRefName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.CreateReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReference`: Reference1
	fmt.Fprintf(os.Stdout, "Response from `V1API.CreateReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reference2** | [**Reference2**](Reference2.md) | Reference to create. | 
 **sourceRefName** | **string** | Source named reference | 

### Return type

[**Reference1**](Reference1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespace

> DeleteNamespace(ctx, name, ref).HashOnRef(hashOnRef).Execute()

Deletes a Namespace

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
	name := *openapiclient.NewContent1AnyOf2([]string{"Elements_example"}, map[string]string{"key": "Inner_example"}) // Content1AnyOf2 | the name of the namespace
	ref := "main" // string | name of ref to fetch
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V1API.DeleteNamespace(context.Background(), name, ref).HashOnRef(hashOnRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.DeleteNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**Content1AnyOf2**](.md) | the name of the namespace | 
**ref** | **string** | name of ref to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hashOnRef** | **string** | a particular hash on the given ref | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReference

> DeleteReference(ctx, referenceName, referenceType).ExpectedHash(expectedHash).Execute()

Delete a reference endpoint

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
	referenceName := "main" // string | Reference name to delete
	referenceType := "branch" // string | Reference type to delete
	expectedHash := "1122334455667788112233445566778811223344556677881122334455667788" // string | Expected hash of tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V1API.DeleteReference(context.Background(), referenceName, referenceType).ExpectedHash(expectedHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.DeleteReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceName** | **string** | Reference name to delete | 
**referenceType** | **string** | Reference type to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expectedHash** | **string** | Expected hash of tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReferences

> GetAllReferences200Response GetAllReferences(ctx).Fetch(fetch).Filter(filter).MaxRecords(maxRecords).PageToken(pageToken).Execute()

Get all references

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
	fetch := "fetch_example" // string | Specify how much information to be returned. Will fetch additional metadata for references if set to 'ALL'.  A returned Branch instance will have the following information:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits in this reference).  A returned Tag instance will only contain the 'commitMetaOfHEAD' and 'numTotalCommits' fields.  Note that computing & fetching additional metadata might be computationally expensive on the server-side, so this flag should be used with care. (optional)
	filter := "refType == 'BRANCH'" // string | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md. Usable variables within the expression are:  - ref (Reference) describes the reference, with fields name (String), hash (String), metadata (ReferenceMetadata)  - metadata (ReferenceMetadata) shortcut to ref.metadata, never null, but possibly empty  - commit (CommitMeta) - shortcut to ref.metadata.commitMetaOfHEAD, never null, but possibly empty  - refType (String) - the reference type, either BRANCH or TAG  Note that the expression can only test attributes metadata and commit, if 'fetchOption' is set to 'ALL'. (optional)
	maxRecords := int32(56) // int32 | maximum number of entries to return, just a hint for the server (optional)
	pageToken := "pageToken_example" // string | paging continuation token, as returned in the previous value of the field 'token' in the corresponding 'EntriesResponse' or 'LogResponse' or 'ReferencesResponse' or 'RefLogResponse'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetAllReferences(context.Background()).Fetch(fetch).Filter(filter).MaxRecords(maxRecords).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetAllReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReferences`: GetAllReferences200Response
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetAllReferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetch** | **string** | Specify how much information to be returned. Will fetch additional metadata for references if set to &#39;ALL&#39;.  A returned Branch instance will have the following information:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits in this reference).  A returned Tag instance will only contain the &#39;commitMetaOfHEAD&#39; and &#39;numTotalCommits&#39; fields.  Note that computing &amp; fetching additional metadata might be computationally expensive on the server-side, so this flag should be used with care. | 
 **filter** | **string** | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md. Usable variables within the expression are:  - ref (Reference) describes the reference, with fields name (String), hash (String), metadata (ReferenceMetadata)  - metadata (ReferenceMetadata) shortcut to ref.metadata, never null, but possibly empty  - commit (CommitMeta) - shortcut to ref.metadata.commitMetaOfHEAD, never null, but possibly empty  - refType (String) - the reference type, either BRANCH or TAG  Note that the expression can only test attributes metadata and commit, if &#39;fetchOption&#39; is set to &#39;ALL&#39;. | 
 **maxRecords** | **int32** | maximum number of entries to return, just a hint for the server | 
 **pageToken** | **string** | paging continuation token, as returned in the previous value of the field &#39;token&#39; in the corresponding &#39;EntriesResponse&#39; or &#39;LogResponse&#39; or &#39;ReferencesResponse&#39; or &#39;RefLogResponse&#39;. | 

### Return type

[**GetAllReferences200Response**](GetAllReferences200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommitLog

> LogResponse1 GetCommitLog(ctx, ref).EndHash(endHash).Fetch(fetch).Filter(filter).MaxRecords(maxRecords).PageToken(pageToken).StartHash(startHash).Execute()

Get commit log for a reference



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
	ref := "main" // string | ref to show log from
	endHash := "endHash_example" // string | Hash on the given ref to end at (in chronological sense), the 'near' end of the commit log, returned 'early' in the result. (optional)
	fetch := "fetch_example" // string | Specify how much information to be returned. Will fetch additional metadata such as parent commit hash and operations in a commit, for each commit if set to 'ALL'. (optional)
	filter := "commit.author=='nessie_author'" // string | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md.  Usable variables within the expression are:  - 'commit' with fields 'author' (string), 'committer' (string), 'commitTime' (timestamp), 'hash' (string), ',message' (string), 'properties' (map)  - 'operations' (list), each operation has the fields 'type' (string, either 'PUT' or 'DELETE'), 'key' (string, namespace + table name), 'keyElements' (list of strings), 'namespace' (string), 'namespaceElements' (list of strings) and 'name' (string, the \"simple\" table name)  Note that the expression can only test against 'operations', if 'fetch' is set to 'ALL'.  Hint: when filtering commits, you can determine whether commits are \"missing\" (filtered) by checking whether 'LogEntry.parentCommitHash' is different from the hash of the previous commit in the log response. (optional)
	maxRecords := int32(56) // int32 | maximum number of entries to return, just a hint for the server (optional)
	pageToken := "pageToken_example" // string | paging continuation token, as returned in the previous value of the field 'token' in the corresponding 'EntriesResponse' or 'LogResponse' or 'ReferencesResponse' or 'RefLogResponse'. (optional)
	startHash := "startHash_example" // string | Hash on the given ref to start from (in chronological sense), the 'far' end of the commit log, returned 'late' in the result. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetCommitLog(context.Background(), ref).EndHash(endHash).Fetch(fetch).Filter(filter).MaxRecords(maxRecords).PageToken(pageToken).StartHash(startHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetCommitLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommitLog`: LogResponse1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetCommitLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | ref to show log from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endHash** | **string** | Hash on the given ref to end at (in chronological sense), the &#39;near&#39; end of the commit log, returned &#39;early&#39; in the result. | 
 **fetch** | **string** | Specify how much information to be returned. Will fetch additional metadata such as parent commit hash and operations in a commit, for each commit if set to &#39;ALL&#39;. | 
 **filter** | **string** | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md.  Usable variables within the expression are:  - &#39;commit&#39; with fields &#39;author&#39; (string), &#39;committer&#39; (string), &#39;commitTime&#39; (timestamp), &#39;hash&#39; (string), &#39;,message&#39; (string), &#39;properties&#39; (map)  - &#39;operations&#39; (list), each operation has the fields &#39;type&#39; (string, either &#39;PUT&#39; or &#39;DELETE&#39;), &#39;key&#39; (string, namespace + table name), &#39;keyElements&#39; (list of strings), &#39;namespace&#39; (string), &#39;namespaceElements&#39; (list of strings) and &#39;name&#39; (string, the \&quot;simple\&quot; table name)  Note that the expression can only test against &#39;operations&#39;, if &#39;fetch&#39; is set to &#39;ALL&#39;.  Hint: when filtering commits, you can determine whether commits are \&quot;missing\&quot; (filtered) by checking whether &#39;LogEntry.parentCommitHash&#39; is different from the hash of the previous commit in the log response. | 
 **maxRecords** | **int32** | maximum number of entries to return, just a hint for the server | 
 **pageToken** | **string** | paging continuation token, as returned in the previous value of the field &#39;token&#39; in the corresponding &#39;EntriesResponse&#39; or &#39;LogResponse&#39; or &#39;ReferencesResponse&#39; or &#39;RefLogResponse&#39;. | 
 **startHash** | **string** | Hash on the given ref to start from (in chronological sense), the &#39;far&#39; end of the commit log, returned &#39;late&#39; in the result. | 

### Return type

[**LogResponse1**](LogResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> NessieConfiguration1 GetConfig(ctx).Execute()

List all configuration settings

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: NessieConfiguration1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


### Return type

[**NessieConfiguration1**](NessieConfiguration1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContent

> Content1 GetContent(ctx, key).HashOnRef(hashOnRef).Ref(ref).Execute()

Get object content associated with a key.



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
	key := "example.key" // string | object name to search for
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)
	ref := "main" // string | Reference to use. Defaults to default branch if not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetContent(context.Background(), key).HashOnRef(hashOnRef).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContent`: Content1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | object name to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hashOnRef** | **string** | a particular hash on the given ref | 
 **ref** | **string** | Reference to use. Defaults to default branch if not provided. | 

### Return type

[**Content1**](Content1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultBranch

> Branch1 GetDefaultBranch(ctx).Execute()

Get default branch for commits and reads

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetDefaultBranch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetDefaultBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultBranch`: Branch1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetDefaultBranch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultBranchRequest struct via the builder pattern


### Return type

[**Branch1**](Branch1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiff

> DiffResponse1 GetDiff(ctx, fromRefWithHash, toRefWithHash).Execute()

Get a diff for two given references



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
	fromRefWithHash := "main" // string | The 'from' reference (and optional hash) to start the diff from
	toRefWithHash := "main" // string | The 'to' reference (and optional hash) to end the diff at.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetDiff(context.Background(), fromRefWithHash, toRefWithHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiff`: DiffResponse1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromRefWithHash** | **string** | The &#39;from&#39; reference (and optional hash) to start the diff from | 
**toRefWithHash** | **string** | The &#39;to&#39; reference (and optional hash) to end the diff at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DiffResponse1**](DiffResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntries

> GetEntries200Response GetEntries(ctx, ref).Filter(filter).HashOnRef(hashOnRef).MaxRecords(maxRecords).NamespaceDepth(namespaceDepth).PageToken(pageToken).Execute()

Fetch all entries for a given reference



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
	ref := "main" // string | name of ref to fetch from
	filter := "entry.namespace.startsWith('a.b.c')" // string | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md. Usable variables within the expression are 'entry.namespace' (string) & 'entry.contentType' (string) (optional)
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)
	maxRecords := int32(56) // int32 | maximum number of entries to return, just a hint for the server (optional)
	namespaceDepth := int32(56) // int32 | If set > 0 will filter the results to only return namespaces/tables to the depth of namespaceDepth. If not set or <=0 has no effect Setting this parameter > 0 will turn off paging. (optional)
	pageToken := "pageToken_example" // string | paging continuation token, as returned in the previous value of the field 'token' in the corresponding 'EntriesResponse' or 'LogResponse' or 'ReferencesResponse' or 'RefLogResponse'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetEntries(context.Background(), ref).Filter(filter).HashOnRef(hashOnRef).MaxRecords(maxRecords).NamespaceDepth(namespaceDepth).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntries`: GetEntries200Response
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | name of ref to fetch from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A Common Expression Language (CEL) expression. An intro to CEL can be found at https://github.com/google/cel-spec/blob/master/doc/intro.md. Usable variables within the expression are &#39;entry.namespace&#39; (string) &amp; &#39;entry.contentType&#39; (string) | 
 **hashOnRef** | **string** | a particular hash on the given ref | 
 **maxRecords** | **int32** | maximum number of entries to return, just a hint for the server | 
 **namespaceDepth** | **int32** | If set &gt; 0 will filter the results to only return namespaces/tables to the depth of namespaceDepth. If not set or &lt;&#x3D;0 has no effect Setting this parameter &gt; 0 will turn off paging. | 
 **pageToken** | **string** | paging continuation token, as returned in the previous value of the field &#39;token&#39; in the corresponding &#39;EntriesResponse&#39; or &#39;LogResponse&#39; or &#39;ReferencesResponse&#39; or &#39;RefLogResponse&#39;. | 

### Return type

[**GetEntries200Response**](GetEntries200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleContents

> GetMultipleContentsResponse1 GetMultipleContents(ctx).GetMultipleContentsRequest1(getMultipleContentsRequest1).HashOnRef(hashOnRef).Ref(ref).Execute()

Get multiple objects' content.



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
	getMultipleContentsRequest1 := *openapiclient.NewGetMultipleContentsRequest1([]openapiclient.GetMultipleContentsRequest1RequestedKeysInner{*openapiclient.NewGetMultipleContentsRequest1RequestedKeysInner([]string{"Elements_example"})}) // GetMultipleContentsRequest1 | Keys to retrieve.
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)
	ref := "main" // string | Reference to use. Defaults to default branch if not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetMultipleContents(context.Background()).GetMultipleContentsRequest1(getMultipleContentsRequest1).HashOnRef(hashOnRef).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetMultipleContents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleContents`: GetMultipleContentsResponse1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetMultipleContents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getMultipleContentsRequest1** | [**GetMultipleContentsRequest1**](GetMultipleContentsRequest1.md) | Keys to retrieve. | 
 **hashOnRef** | **string** | a particular hash on the given ref | 
 **ref** | **string** | Reference to use. Defaults to default branch if not provided. | 

### Return type

[**GetMultipleContentsResponse1**](GetMultipleContentsResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespace

> Content1AnyOf2 GetNamespace(ctx, name, ref).HashOnRef(hashOnRef).Execute()

Retrieves a Namespace

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
	name := *openapiclient.NewContent1AnyOf2([]string{"Elements_example"}, map[string]string{"key": "Inner_example"}) // Content1AnyOf2 | the name of the namespace
	ref := "main" // string | name of ref to fetch
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetNamespace(context.Background(), name, ref).HashOnRef(hashOnRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespace`: Content1AnyOf2
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**Content1AnyOf2**](.md) | the name of the namespace | 
**ref** | **string** | name of ref to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hashOnRef** | **string** | a particular hash on the given ref | 

### Return type

[**Content1AnyOf2**](Content1AnyOf2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaces

> GetNamespaces200Response GetNamespaces(ctx, ref).HashOnRef(hashOnRef).Name(name).Execute()



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
	ref := "main" // string | name of ref to fetch
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)
	name := *openapiclient.NewContent1AnyOf2([]string{"Elements_example"}, map[string]string{"key": "Inner_example"}) // Content1AnyOf2 | the name of the namespace (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetNamespaces(context.Background(), ref).HashOnRef(hashOnRef).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaces`: GetNamespaces200Response
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | name of ref to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hashOnRef** | **string** | a particular hash on the given ref | 
 **name** | [**Content1AnyOf2**](Content1AnyOf2.md) | the name of the namespace | 

### Return type

[**GetNamespaces200Response**](GetNamespaces200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferenceByName

> Reference1 GetReferenceByName(ctx, ref).Fetch(fetch).Execute()

Fetch details of a reference

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
	ref := "main" // string | name of ref to fetch
	fetch := "fetch_example" // string | Specify how much information to be returned. Will fetch additional metadata for references if set to 'ALL'.  A returned Branch instance will have the following information:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits in this reference).  A returned Tag instance will only contain the 'commitMetaOfHEAD' and 'numTotalCommits' fields.  Note that computing & fetching additional metadata might be computationally expensive on the server-side, so this flag should be used with care. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.GetReferenceByName(context.Background(), ref).Fetch(fetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.GetReferenceByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferenceByName`: Reference1
	fmt.Fprintf(os.Stdout, "Response from `V1API.GetReferenceByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | name of ref to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferenceByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetch** | **string** | Specify how much information to be returned. Will fetch additional metadata for references if set to &#39;ALL&#39;.  A returned Branch instance will have the following information:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits in this reference).  A returned Tag instance will only contain the &#39;commitMetaOfHEAD&#39; and &#39;numTotalCommits&#39; fields.  Note that computing &amp; fetching additional metadata might be computationally expensive on the server-side, so this flag should be used with care. | 

### Return type

[**Reference1**](Reference1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeRefIntoBranch

> MergeResponse MergeRefIntoBranch(ctx, branchName).ExpectedHash(expectedHash).MergeOperation(mergeOperation).Execute()

Merge commits from 'mergeRef' onto 'branchName'.



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
	branchName := "main" // string | Branch to merge into
	expectedHash := "1122334455667788112233445566778811223344556677881122334455667788" // string | Expected current HEAD of 'branchName'
	mergeOperation := *openapiclient.NewMergeOperation("FromRefName_example", "FromHash_example") // MergeOperation | Merge operation that defines the source reference name and an optional hash. If 'fromHash' is not present, the current 'sourceRef's HEAD will be used.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MergeRefIntoBranch(context.Background(), branchName).ExpectedHash(expectedHash).MergeOperation(mergeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MergeRefIntoBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeRefIntoBranch`: MergeResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.MergeRefIntoBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchName** | **string** | Branch to merge into | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeRefIntoBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expectedHash** | **string** | Expected current HEAD of &#39;branchName&#39; | 
 **mergeOperation** | [**MergeOperation**](MergeOperation.md) | Merge operation that defines the source reference name and an optional hash. If &#39;fromHash&#39; is not present, the current &#39;sourceRef&#39;s HEAD will be used. | 

### Return type

[**MergeResponse**](MergeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransplantCommitsIntoBranch

> MergeResponse TransplantCommitsIntoBranch(ctx, branchName).ExpectedHash(expectedHash).Transplant1(transplant1).Message(message).Execute()

Transplant commits from 'transplant' onto 'branchName'



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
	branchName := "main" // string | Branch to transplant into
	expectedHash := "1122334455667788112233445566778811223344556677881122334455667788" // string | Expected hash of tag.
	transplant1 := *openapiclient.NewTransplant1("FromRefName_example", []string{"HashesToTransplant_example"}) // Transplant1 | Hashes to transplant
	message := "Example Commit Message" // string | commit message (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.TransplantCommitsIntoBranch(context.Background(), branchName).ExpectedHash(expectedHash).Transplant1(transplant1).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.TransplantCommitsIntoBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransplantCommitsIntoBranch`: MergeResponse
	fmt.Fprintf(os.Stdout, "Response from `V1API.TransplantCommitsIntoBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchName** | **string** | Branch to transplant into | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransplantCommitsIntoBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expectedHash** | **string** | Expected hash of tag. | 
 **transplant1** | [**Transplant1**](Transplant1.md) | Hashes to transplant | 
 **message** | **string** | commit message | 

### Return type

[**MergeResponse**](MergeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProperties

> UpdateProperties(ctx, name, ref).UpdatePropertiesRequest(updatePropertiesRequest).HashOnRef(hashOnRef).Execute()



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
	name := *openapiclient.NewContent1AnyOf2([]string{"Elements_example"}, map[string]string{"key": "Inner_example"}) // Content1AnyOf2 | the name of the namespace
	ref := "main" // string | name of ref to fetch
	updatePropertiesRequest := *openapiclient.NewUpdatePropertiesRequest() // UpdatePropertiesRequest | Namespace properties to update/delete.
	hashOnRef := "hashOnRef_example" // string | a particular hash on the given ref (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V1API.UpdateProperties(context.Background(), name, ref).UpdatePropertiesRequest(updatePropertiesRequest).HashOnRef(hashOnRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.UpdateProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**Content1AnyOf2**](.md) | the name of the namespace | 
**ref** | **string** | name of ref to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePropertiesRequest** | [**UpdatePropertiesRequest**](UpdatePropertiesRequest.md) | Namespace properties to update/delete. | 
 **hashOnRef** | **string** | a particular hash on the given ref | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

