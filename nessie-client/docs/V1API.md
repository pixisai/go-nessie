# \V1API

All URIs are relative to *http://localhost:19120*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllReferences**](V1API.md#GetAllReferences) | **Get** /v1/trees | Get all references



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

