package providers

import "danielr1996/bashdoard/sse"

type Provider interface {
	Push(s *sse.SSE)
}
