/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
	"os"
)


// ContactsApiService ContactsApi service
type ContactsApiService service

type ApiContactsByEmailDeleteRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	email string
}

func (r ApiContactsByEmailDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContactsByEmailDeleteExecute(r)
}

/*
ContactsByEmailDelete Delete Contact

Deletes the provided contact. Required Access Level: ModifyContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param email Proper email address.
 @return ApiContactsByEmailDeleteRequest
*/
func (a *ContactsApiService) ContactsByEmailDelete(ctx context.Context, email string) ApiContactsByEmailDeleteRequest {
	return ApiContactsByEmailDeleteRequest{
		ApiService: a,
		ctx: ctx,
		email: email,
	}
}

// Execute executes the request
func (a *ContactsApiService) ContactsByEmailDeleteExecute(r ApiContactsByEmailDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsByEmailDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterToString(r.email, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiContactsByEmailGetRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	email string
}

func (r ApiContactsByEmailGetRequest) Execute() (*Contact, *http.Response, error) {
	return r.ApiService.ContactsByEmailGetExecute(r)
}

/*
ContactsByEmailGet Load Contact

Load detailed contact information for specified email. Required Access Level: ViewContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param email Proper email address.
 @return ApiContactsByEmailGetRequest
*/
func (a *ContactsApiService) ContactsByEmailGet(ctx context.Context, email string) ApiContactsByEmailGetRequest {
	return ApiContactsByEmailGetRequest{
		ApiService: a,
		ctx: ctx,
		email: email,
	}
}

// Execute executes the request
//  @return Contact
func (a *ContactsApiService) ContactsByEmailGetExecute(r ApiContactsByEmailGetRequest) (*Contact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Contact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsByEmailGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterToString(r.email, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContactsByEmailPutRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	email string
	contactUpdatePayload *ContactUpdatePayload
}

func (r ApiContactsByEmailPutRequest) ContactUpdatePayload(contactUpdatePayload ContactUpdatePayload) ApiContactsByEmailPutRequest {
	r.contactUpdatePayload = &contactUpdatePayload
	return r
}

func (r ApiContactsByEmailPutRequest) Execute() (*Contact, *http.Response, error) {
	return r.ApiService.ContactsByEmailPutExecute(r)
}

/*
ContactsByEmailPut Update Contact

Update selected contact. Omitted contact's fields will not be changed. Required Access Level: ModifyContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param email Proper email address.
 @return ApiContactsByEmailPutRequest
*/
func (a *ContactsApiService) ContactsByEmailPut(ctx context.Context, email string) ApiContactsByEmailPutRequest {
	return ApiContactsByEmailPutRequest{
		ApiService: a,
		ctx: ctx,
		email: email,
	}
}

// Execute executes the request
//  @return Contact
func (a *ContactsApiService) ContactsByEmailPutExecute(r ApiContactsByEmailPutRequest) (*Contact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Contact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsByEmailPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterToString(r.email, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contactUpdatePayload == nil {
		return localVarReturnValue, nil, reportError("contactUpdatePayload is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.contactUpdatePayload
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContactsDeletePostRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	emailsPayload *EmailsPayload
}

// Provide either rule or a list of emails, not both.
func (r ApiContactsDeletePostRequest) EmailsPayload(emailsPayload EmailsPayload) ApiContactsDeletePostRequest {
	r.emailsPayload = &emailsPayload
	return r
}

func (r ApiContactsDeletePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContactsDeletePostExecute(r)
}

/*
ContactsDeletePost Delete Contacts Bulk

Deletes provided contacts in bulk. Required Access Level: ModifyContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContactsDeletePostRequest
*/
func (a *ContactsApiService) ContactsDeletePost(ctx context.Context) ApiContactsDeletePostRequest {
	return ApiContactsDeletePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ContactsApiService) ContactsDeletePostExecute(r ApiContactsDeletePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsDeletePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.emailsPayload == nil {
		return nil, reportError("emailsPayload is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.emailsPayload
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiContactsExportByIdStatusGetRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	id string
}

func (r ApiContactsExportByIdStatusGetRequest) Execute() (*ExportStatus, *http.Response, error) {
	return r.ApiService.ContactsExportByIdStatusGetExecute(r)
}

/*
ContactsExportByIdStatusGet Check Export Status

Check the current status of the export. Required Access Level: Export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the exported file
 @return ApiContactsExportByIdStatusGetRequest
*/
func (a *ContactsApiService) ContactsExportByIdStatusGet(ctx context.Context, id string) ApiContactsExportByIdStatusGetRequest {
	return ApiContactsExportByIdStatusGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ExportStatus
func (a *ContactsApiService) ContactsExportByIdStatusGetExecute(r ApiContactsExportByIdStatusGetRequest) (*ExportStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsExportByIdStatusGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/export/{id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContactsExportPostRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	fileFormat *ExportFileFormats
	rule *string
	emails *[]string
	compressionFormat *CompressionFormat
	fileName *string
}

// Format of the exported file
func (r ApiContactsExportPostRequest) FileFormat(fileFormat ExportFileFormats) ApiContactsExportPostRequest {
	r.fileFormat = &fileFormat
	return r
}

// Query used for filtering.
func (r ApiContactsExportPostRequest) Rule(rule string) ApiContactsExportPostRequest {
	r.rule = &rule
	return r
}

// Comma delimited list of contact emails
func (r ApiContactsExportPostRequest) Emails(emails []string) ApiContactsExportPostRequest {
	r.emails = &emails
	return r
}

// FileResponse compression format. None or Zip.
func (r ApiContactsExportPostRequest) CompressionFormat(compressionFormat CompressionFormat) ApiContactsExportPostRequest {
	r.compressionFormat = &compressionFormat
	return r
}

// Name of your file including extension.
func (r ApiContactsExportPostRequest) FileName(fileName string) ApiContactsExportPostRequest {
	r.fileName = &fileName
	return r
}

func (r ApiContactsExportPostRequest) Execute() (*ExportLink, *http.Response, error) {
	return r.ApiService.ContactsExportPostExecute(r)
}

/*
ContactsExportPost Export Contacts

Request an Export of specified Contacts. Required Access Level: Export

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContactsExportPostRequest
*/
func (a *ContactsApiService) ContactsExportPost(ctx context.Context) ApiContactsExportPostRequest {
	return ApiContactsExportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExportLink
func (a *ContactsApiService) ContactsExportPostExecute(r ApiContactsExportPostRequest) (*ExportLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExportLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsExportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fileFormat != nil {
		localVarQueryParams.Add("fileFormat", parameterToString(*r.fileFormat, ""))
	}
	if r.rule != nil {
		localVarQueryParams.Add("rule", parameterToString(*r.rule, ""))
	}
	if r.emails != nil {
		t := *r.emails
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("emails", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("emails", parameterToString(t, "multi"))
		}
	}
	if r.compressionFormat != nil {
		localVarQueryParams.Add("compressionFormat", parameterToString(*r.compressionFormat, ""))
	}
	if r.fileName != nil {
		localVarQueryParams.Add("fileName", parameterToString(*r.fileName, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContactsGetRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	limit *int32
	offset *int32
}

// Maximum number of returned items.
func (r ApiContactsGetRequest) Limit(limit int32) ApiContactsGetRequest {
	r.limit = &limit
	return r
}

// How many items should be returned ahead.
func (r ApiContactsGetRequest) Offset(offset int32) ApiContactsGetRequest {
	r.offset = &offset
	return r
}

func (r ApiContactsGetRequest) Execute() ([]Contact, *http.Response, error) {
	return r.ApiService.ContactsGetExecute(r)
}

/*
ContactsGet Load Contacts

Returns a list of contacts. Required Access Level: ViewContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContactsGetRequest
*/
func (a *ContactsApiService) ContactsGet(ctx context.Context) ApiContactsGetRequest {
	return ApiContactsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Contact
func (a *ContactsApiService) ContactsGetExecute(r ApiContactsGetRequest) ([]Contact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Contact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiContactsImportPostRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	listName *string
	encodingName *string
	fileUrl *string
	file **os.File
}

// Name of an existing list to add these contacts to
func (r ApiContactsImportPostRequest) ListName(listName string) ApiContactsImportPostRequest {
	r.listName = &listName
	return r
}

// In what encoding the file is uploaded
func (r ApiContactsImportPostRequest) EncodingName(encodingName string) ApiContactsImportPostRequest {
	r.encodingName = &encodingName
	return r
}

// Optional url of csv to import
func (r ApiContactsImportPostRequest) FileUrl(fileUrl string) ApiContactsImportPostRequest {
	r.fileUrl = &fileUrl
	return r
}

func (r ApiContactsImportPostRequest) File(file *os.File) ApiContactsImportPostRequest {
	r.file = &file
	return r
}

func (r ApiContactsImportPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContactsImportPostExecute(r)
}

/*
ContactsImportPost Upload Contacts

Upload contacts from a file. Required Access Level: ModifyContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContactsImportPostRequest
*/
func (a *ContactsApiService) ContactsImportPost(ctx context.Context) ApiContactsImportPostRequest {
	return ApiContactsImportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ContactsApiService) ContactsImportPostExecute(r ApiContactsImportPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsImportPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.listName != nil {
		localVarQueryParams.Add("listName", parameterToString(*r.listName, ""))
	}
	if r.encodingName != nil {
		localVarQueryParams.Add("encodingName", parameterToString(*r.encodingName, ""))
	}
	if r.fileUrl != nil {
		localVarQueryParams.Add("fileUrl", parameterToString(*r.fileUrl, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	var fileLocalVarFile *os.File
	if r.file != nil {
		fileLocalVarFile = *r.file
	}
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiContactsPostRequest struct {
	ctx context.Context
	ApiService *ContactsApiService
	contactPayload *[]ContactPayload
	listnames *[]string
}

func (r ApiContactsPostRequest) ContactPayload(contactPayload []ContactPayload) ApiContactsPostRequest {
	r.contactPayload = &contactPayload
	return r
}

// Names of lists to which the uploaded contacts should be added to
func (r ApiContactsPostRequest) Listnames(listnames []string) ApiContactsPostRequest {
	r.listnames = &listnames
	return r
}

func (r ApiContactsPostRequest) Execute() ([]Contact, *http.Response, error) {
	return r.ApiService.ContactsPostExecute(r)
}

/*
ContactsPost Add Contact

Add new Contacts to your Lists. Up to 1000 can be added (for more please refer to the import request). Required Access Level: ModifyContacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiContactsPostRequest
*/
func (a *ContactsApiService) ContactsPost(ctx context.Context) ApiContactsPostRequest {
	return ApiContactsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Contact
func (a *ContactsApiService) ContactsPostExecute(r ApiContactsPostRequest) ([]Contact, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Contact
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactsApiService.ContactsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/contacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contactPayload == nil {
		return localVarReturnValue, nil, reportError("contactPayload is required and must be specified")
	}

	if r.listnames != nil {
		t := *r.listnames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("listnames", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("listnames", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.contactPayload
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-ElasticEmail-ApiKey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
