package errors

type BadRequest struct {
	Message string
}

func (b *BadRequest) Error() string {
	if b.Message == "" {
		return "Bad Request"
	}

	return b.Message
}

type InternalServerError struct {
	Message string
}

func (i *InternalServerError) Error() string {
	if i.Message == "" {
		return "Oops! something went wrong!!!"
	}

	return i.Message
}

type Unauthorized struct {
	Message string
}

func (u *Unauthorized) Error() string {
	if u.Message == "" {
		return "Unauthorized"
	}

	return u.Message
}

type NotFound struct {
	Message string
}

func (n *NotFound) Error() string {
	if n.Message == "" {
		return "Not Found"
	}

	return n.Message
}

type TokenMalformed struct {
	Message string
}

func (t *TokenMalformed) Error() string {
	if t.Message == "" {
		return "The token is malformed"
	}

	return t.Message
}

type TokenExpired struct {
	Message string
}

func (t *TokenExpired) Error() string {
	if t.Message == "" {
		return "The token has expired"
	}

	return t.Message
}

type TokenSignatureInvalid struct {
	Message string
}

func (t *TokenSignatureInvalid) Error() string {
	if t.Message == "" {
		return "The token signature is invalid"
	}

	return t.Message
}
