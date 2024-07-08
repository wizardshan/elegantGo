package request

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validate = binding.Validator.Engine().(*validator.Validate)
