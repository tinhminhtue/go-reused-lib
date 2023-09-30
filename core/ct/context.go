package ct

type key int

const CtxVerKey key = iota

// type CtxVerKey struct{} // or exported to use outside the package
// const ForceSampleKey CtxVerKey = new CtxVerKey{}

// usage
// ctx = context.WithValue(ctx, ctxKey{}, 123)
// fmt.Println(ctx.Value(ctxKey{}).(int) == 123) // true
