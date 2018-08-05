// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	context "context"
	strconv "strconv"

	graphql "github.com/99designs/gqlgen/graphql"
	introspection "github.com/99designs/gqlgen/graphql/introspection"
	models "github.com/conradwt/gqlgen_example/models"
	gqlparser "github.com/vektah/gqlparser"
	ast "github.com/vektah/gqlparser/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
	Todo() TodoResolver
}

type DirectiveRoot struct {
}
type MutationResolver interface {
	CreateTodo(ctx context.Context, input models.NewTodo) (models.Todo, error)
}
type QueryResolver interface {
	Todos(ctx context.Context) ([]models.Todo, error)
}
type TodoResolver interface {
	User(ctx context.Context, obj *models.Todo) (models.User, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:   buf,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Mutation(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:   buf,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

var mutationImplementors = []string{"Mutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, mutationImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Mutation",
	})

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createTodo":
			out.Values[i] = ec._Mutation_createTodo(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Mutation_createTodo(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 models.NewTodo
	if tmp, ok := rawArgs["input"]; ok {
		var err error
		arg0, err = UnmarshalNewTodo(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["input"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Mutation"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.resolvers.Mutation().CreateTodo(ctx, args["input"].(models.NewTodo))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(models.Todo)
	return ec._Todo(ctx, field.Selections, &res)
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "todos":
			out.Values[i] = ec._Query_todos(ctx, field)
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Query_todos(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
		Args:   nil,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.Query().Todos(ctx)
		})
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.([]models.Todo)
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler {
				rctx := graphql.GetResolverContext(ctx)
				rctx.PushIndex(idx1)
				defer rctx.Pop()
				return ec._Todo(ctx, field.Selections, &res[idx1])
			}())
		}
		return arr1
	})
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["name"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Query"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.introspectType(args["name"].(string)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Query"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.introspectSchema(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	if res == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, field.Selections, res)
}

var todoImplementors = []string{"Todo"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Todo(ctx context.Context, sel ast.SelectionSet, obj *models.Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, todoImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			out.Values[i] = ec._Todo_id(ctx, field, obj)
		case "text":
			out.Values[i] = ec._Todo_text(ctx, field, obj)
		case "done":
			out.Values[i] = ec._Todo_done(ctx, field, obj)
		case "user":
			out.Values[i] = ec._Todo_user(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Todo_id(ctx context.Context, field graphql.CollectedField, obj *models.Todo) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Todo"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.ID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalID(res)
}

func (ec *executionContext) _Todo_text(ctx context.Context, field graphql.CollectedField, obj *models.Todo) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Todo"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Text, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) _Todo_done(ctx context.Context, field graphql.CollectedField, obj *models.Todo) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Todo"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Done, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(bool)
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) _Todo_user(ctx context.Context, field graphql.CollectedField, obj *models.Todo) graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Todo",
		Args:   nil,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.Todo().User(ctx, obj)
		})
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.(models.User)
		return ec._User(ctx, field.Selections, &res)
	})
}

var userImplementors = []string{"User"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *models.User) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, userImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "id":
			out.Values[i] = ec._User_id(ctx, field, obj)
		case "name":
			out.Values[i] = ec._User_name(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *models.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.ID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalID(res)
}

func (ec *executionContext) _User_name(ctx context.Context, field graphql.CollectedField, obj *models.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Locations, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]string)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return graphql.MarshalString(res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Args, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(bool)
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Args, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Type, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(bool)
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Type, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Types(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Directives(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Directive(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Kind(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Field(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___EnumValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func UnmarshalNewTodo(v interface{}) (models.NewTodo, error) {
	var it models.NewTodo
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "text":
			var err error
			it.Text, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "userId":
			var err error
			it.UserID, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) FieldMiddleware(ctx context.Context, next graphql.Resolver) interface{} {
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name])
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `type Todo {
  id: ID!
  text: String!
  done: Boolean!
  user: User!
}

type User {
  id: ID!
  name: String!
}

type Query {
  todos: [Todo!]!
}

input NewTodo {
  text: String!
  userId: String!
}

type Mutation {
  createTodo(input: NewTodo!): Todo!
}
`},
)
