package array2string

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

const (
	methodPOST  = "POST"
	methodPUT   = "PUT"
	methodPATCH = "PATCH"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{settings: s}

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
// settings : {method, uri, headers, proxy, skipSSL}
// input    : {pathParams, queryParams, headers, content}
// outputs  : {status, result}
type Activity struct {
	settings *Settings
	// containsParam bool
	// client        *http.Client
}

// Metadata Activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}
	delimeter := a.settings.Delimeter
	prefix := a.settings.Prefix
	suffix := a.settings.Suffix

	// 	uri = uri + "?" + qp.Encode()
	// }

	logger := ctx.Logger()

	if logger.DebugEnabled() {
		logger.Debugf("Eval called: [%s] %s", a.settings.Delimeter, delimeter)
		logger.Debugf("Eval called: [%s] %s", a.settings.Prefix, prefix)
		logger.Debugf("Eval called: [%s] %s", a.settings.Suffix, suffix)
	}

	var result string

	for value := range input.InputArray {
		// qp.Set(key, value)

		nval, err := coerce.ToString(value)
		if err != nil {
			return false, err
		}
		result = result + nval
	}

	// result = "Result !"
	// }

	// if logger.TraceEnabled() {
	// 	logger.Trace("Response body:", result)
	// }

	output := &Output{ResultString: result} //coerce.ToString(result)}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

// func (a *Activity) getHeaders(inputHeaders map[string]string) map[string]string {

// 	if len(inputHeaders) == 0 {
// 		return a.settings.Headers
// 	}

// 	if len(a.settings.Headers) == 0 {
// 		return inputHeaders
// 	}

// 	headers := make(map[string]string)
// 	for key, value := range a.settings.Headers {
// 		headers[key] = value
// 	}
// 	for key, value := range inputHeaders {
// 		headers[key] = value
// 	}

// 	return headers
// }

//todo just make contentType a setting
// func getContentType(replyData interface{}) string {

// 	contentType := "application/json; charset=UTF-8"

// 	switch v := replyData.(type) {
// 	case string:
// 		if !strings.HasPrefix(v, "{") && !strings.HasPrefix(v, "[") {
// 			contentType = "text/plain; charset=UTF-8"
// 		}
// 	case int, int64, float64, bool, json.Number:
// 		contentType = "text/plain; charset=UTF-8"
// 	default:
// 		contentType = "application/json; charset=UTF-8"
// 	}

// 	return contentType
// }

// BuildURI is a temporary crude URI builder
// func BuildURI(uri string, values map[string]string) string {

// 	var buffer bytes.Buffer
// 	buffer.Grow(len(uri))

// 	addrStart := strings.Index(uri, "://")

// 	i := addrStart + 3

// 	for i < len(uri) {
// 		if uri[i] == '/' {
// 			break
// 		}
// 		i++
// 	}

// 	buffer.WriteString(uri[0:i])

// 	for i < len(uri) {
// 		if uri[i] == ':' {
// 			j := i + 1
// 			for j < len(uri) && uri[j] != '/' {
// 				j++
// 			}

// 			if i+1 == j {

// 				buffer.WriteByte(uri[i])
// 				i++
// 			} else {

// 				param := uri[i+1 : j]
// 				value := values[param]
// 				buffer.WriteString(value)
// 				if j < len(uri) {
// 					buffer.WriteString("/")
// 				}
// 				i = j + 1
// 			}

// 		} else {
// 			buffer.WriteByte(uri[i])
// 			i++
// 		}
// 	}

// 	return buffer.String()
// }
