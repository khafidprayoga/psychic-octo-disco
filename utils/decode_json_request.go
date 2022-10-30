package utils

// References: https://tanmaydas.com/posts/parsing-json-request-body-in-go-fiber/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/gofiber/fiber/v2"
	"io"
	"strings"
)

type MalformedRequest struct {
	Status  int
	Message string
}

func (e *MalformedRequest) Error() string {
	return e.Message
}

func decodeJSONBody(c *fiber.Ctx, out any) error {
	dec := decoder.NewStreamDecoder(bytes.NewReader(c.Body()))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&out); err != nil {
		var (
			syntaxErr         *json.SyntaxError
			unmarshallTypeErr *json.UnmarshalTypeError
		)
		switch {
		case errors.As(err, &syntaxErr):
			return &MalformedRequest{
				Status: fiber.StatusBadRequest,
				Message: fmt.Sprintf(
					"Request body contains badly-formed JSON (at position %d)",
					syntaxErr.Offset,
				),
			}

		case errors.Is(err, io.ErrUnexpectedEOF):
			return &MalformedRequest{
				Status:  fiber.StatusBadRequest,
				Message: "Request body contains badly-formed JSON",
			}

		case errors.As(err, &unmarshallTypeErr):
			return &MalformedRequest{
				Status: fiber.StatusBadRequest,
				Message: fmt.Sprintf(
					"Request body contains an invalid value for the %q field (at position %d)",
					unmarshallTypeErr.Field,
					unmarshallTypeErr.Offset,
				),
			}

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return &MalformedRequest{
				Status:  fiber.StatusBadRequest,
				Message: fmt.Sprintf("Request body contains unknown field %s", fieldName),
			}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{
				Status:  fiber.StatusBadRequest,
				Message: msg,
			}

		case err.Error() == "http: request body too large":
			return &MalformedRequest{
				Status:  fiber.StatusRequestEntityTooLarge,
				Message: "Request body must not be larger than 1MB",
			}

		default:
			return err

		}
	}

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return &MalformedRequest{
			Status:  fiber.StatusBadRequest,
			Message: "Request body must only contain a single JSON object",
		}
	}

	return nil
}

func JSONBodyParser(ctx *fiber.Ctx, out any) error {
	var mr *MalformedRequest

	if err := decodeJSONBody(ctx, out); err != nil {
		if errors.As(err, &mr) {
			return ctx.Status(mr.Status).JSON(fiber.Map{
				"message": mr.Message,
			})
		} else {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
	}

	return nil
}
