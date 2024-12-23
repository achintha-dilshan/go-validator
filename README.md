# Custom JSON Validator

This utility function provides a custom JSON validation mechanism for validating API requests in Go. Unlike using external libraries such as `go-playground/validator`, this solution is designed for simplicity and flexibility, allowing you to handle validation errors and return JSON responses tailored to your needs.

## Features
- Validate required fields.
- Validate email addresses.
- Validate minimum field length.

## Example Usage
If you send the following JSON request:

```json
{
  "name": "Te",
  "email": "testtest.com",
  "password": ""
}
```

The validation function will return the following response:

```json
{
  "error": {
    "email": "Enter a valid email.",
    "name": "'name' field must be longer than 3 characters",
    "password": "This field is required."
  }
}
```

## How It Works
The function:
1. Iterates through the struct fields.
2. Reads the `validate` tags for rules and validates the field value accordingly.
3. Supports the following validation rules:
   - **`required`**: Ensures the field is not empty.
   - **`email`**: Checks if the field contains a valid email address.
   - **`min=n`**: Validates if the field's length is at least `n` characters.

### Example Struct
Here is an example of how you might define a struct with validation rules:

```go
import "myproject/validator"

// User struct for validation.
type User struct {
    Name     string `json:"name" validate:"required,min=3"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

v := validator.New()
errors := v.Validate(user)
```

## Extensibility
This package was designed with simplicity in mind and does not include many features. You are free to customize it and add new features as needed to meet your requirements.

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue for bugs, feature requests, or improvements.

---
This custom validator is a lightweight, flexible alternative to third-party solutions, perfect for projects requiring bespoke validation logic.

