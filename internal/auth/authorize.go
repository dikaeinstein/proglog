package auth

import (
	"fmt"

	casbin "github.com/casbin/casbin/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authorizer wraps casbin Enforcer which uses ACL model and a file policy
// adapter to authorize subjects
type Authorizer struct {
	enforcer *casbin.Enforcer
}

// New creates and initializes an Authorizer with the specified model and policy path.
// This will configure Casbin’s authorization mechanism which is ACL for the Authorizer
func New(model, policy string) *Authorizer {
	enforcer, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		panic(err)
	}

	return &Authorizer{enforcer: enforcer}
}

// Authorize defers to Casbin’s Enforce function. This function
// returns whether the given subject is permitted to run the
// given action on the given object based on the model and policy you
// configure Casbin with.
func (a *Authorizer) Authorize(subject, object, action string) error {
	ok, err := a.enforcer.Enforce(subject, object, action)
	if err != nil {
		return err
	}

	if !ok {
		msg := fmt.Sprintf(
			"%s not permitted to %s to %s",
			subject,
			action,
			object,
		)
		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}

	return nil
}
