package middlewares

import "github.com/google/uuid"

type UuidGenerator struct {
	Middleware
}

func NewUuidGenerator() *UuidGenerator {
	return &UuidGenerator{}
}

func (s *UuidGenerator) Generate() (error, string) {
	uuidCode, err := uuid.NewUUID()
	if err != nil {
		return err, ""
	}
	return nil, uuidCode.String()
}
