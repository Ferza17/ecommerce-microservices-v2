// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	gen "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/payment/v1"
	"github.com/vektah/gqlparser/v2/ast"
	"golang.org/x/sync/semaphore"
)

// region    ************************** generated!.gotpl **************************

type ProviderResolver interface {
	CreatedAt(ctx context.Context, obj *gen.Provider) (*time.Time, error)
	UpdatedAt(ctx context.Context, obj *gen.Provider) (*time.Time, error)
	DiscardedAt(ctx context.Context, obj *gen.Provider) (*time.Time, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _FindPaymentProvidersResponse_data(ctx context.Context, field graphql.CollectedField, obj *gen.FindPaymentProvidersResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_FindPaymentProvidersResponse_data(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Data, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*gen.Provider)
	fc.Result = res
	return ec.marshalNProvider2ᚕᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProviderᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_FindPaymentProvidersResponse_data(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "FindPaymentProvidersResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Provider_id(ctx, field)
			case "name":
				return ec.fieldContext_Provider_name(ctx, field)
			case "method":
				return ec.fieldContext_Provider_method(ctx, field)
			case "createdAt":
				return ec.fieldContext_Provider_createdAt(ctx, field)
			case "updatedAt":
				return ec.fieldContext_Provider_updatedAt(ctx, field)
			case "discardedAt":
				return ec.fieldContext_Provider_discardedAt(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Provider", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Provider_id(ctx context.Context, field graphql.CollectedField, obj *gen.Provider) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Provider_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Id, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Provider_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Provider",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Provider_name(ctx context.Context, field graphql.CollectedField, obj *gen.Provider) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Provider_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Provider_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Provider",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Provider_method(ctx context.Context, field graphql.CollectedField, obj *gen.Provider) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Provider_method(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Method, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(gen.ProviderMethod)
	fc.Result = res
	return ec.marshalNProviderMethod2githubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProviderMethod(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Provider_method(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Provider",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ProviderMethod does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Provider_createdAt(ctx context.Context, field graphql.CollectedField, obj *gen.Provider) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Provider_createdAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Provider().CreatedAt(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Provider_createdAt(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Provider",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Provider_updatedAt(ctx context.Context, field graphql.CollectedField, obj *gen.Provider) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Provider_updatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Provider().UpdatedAt(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Provider_updatedAt(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Provider",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Provider_discardedAt(ctx context.Context, field graphql.CollectedField, obj *gen.Provider) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Provider_discardedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Provider().DiscardedAt(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Provider_discardedAt(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Provider",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputFindPaymentProvidersRequest(ctx context.Context, obj any) (gen.FindPaymentProvidersRequest, error) {
	var it gen.FindPaymentProvidersRequest
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalOString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var findPaymentProvidersResponseImplementors = []string{"FindPaymentProvidersResponse"}

func (ec *executionContext) _FindPaymentProvidersResponse(ctx context.Context, sel ast.SelectionSet, obj *gen.FindPaymentProvidersResponse) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, findPaymentProvidersResponseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("FindPaymentProvidersResponse")
		case "data":
			out.Values[i] = ec._FindPaymentProvidersResponse_data(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var providerImplementors = []string{"Provider"}

func (ec *executionContext) _Provider(ctx context.Context, sel ast.SelectionSet, obj *gen.Provider) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, providerImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Provider")
		case "id":
			out.Values[i] = ec._Provider_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "name":
			out.Values[i] = ec._Provider_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "method":
			out.Values[i] = ec._Provider_method(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "createdAt":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Provider_createdAt(ctx, field, obj)
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "updatedAt":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Provider_updatedAt(ctx, field, obj)
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "discardedAt":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Provider_discardedAt(ctx, field, obj)
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNProvider2ᚕᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProviderᚄ(ctx context.Context, sel ast.SelectionSet, v []*gen.Provider) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	sm := semaphore.NewWeighted(1000)
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer func() {
					sm.Release(1)
					wg.Done()
				}()
			}
			ret[i] = ec.marshalNProvider2ᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProvider(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			if err := sm.Acquire(ctx, 1); err != nil {
				ec.Error(ctx, ctx.Err())
			} else {
				go f(i)
			}
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNProvider2ᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProvider(ctx context.Context, sel ast.SelectionSet, v *gen.Provider) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Provider(ctx, sel, v)
}

func (ec *executionContext) unmarshalNProviderMethod2githubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProviderMethod(ctx context.Context, v any) (gen.ProviderMethod, error) {
	var res gen.ProviderMethod
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNProviderMethod2githubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProviderMethod(ctx context.Context, sel ast.SelectionSet, v gen.ProviderMethod) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOFindPaymentProvidersRequest2ᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐFindPaymentProvidersRequest(ctx context.Context, v any) (*gen.FindPaymentProvidersRequest, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputFindPaymentProvidersRequest(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOFindPaymentProvidersResponse2ᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐFindPaymentProvidersResponse(ctx context.Context, sel ast.SelectionSet, v *gen.FindPaymentProvidersResponse) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._FindPaymentProvidersResponse(ctx, sel, v)
}

func (ec *executionContext) marshalOProvider2ᚖgithubᚗcomᚋferza17ᚋecommerceᚑmicroservicesᚑv2ᚋapiᚑgatewayᚋmodelᚋrpcᚋgenᚋpaymentᚋv1ᚐProvider(ctx context.Context, sel ast.SelectionSet, v *gen.Provider) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Provider(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
