// dto/permission.go
package dto

type CreatePermissionRequest struct {
    Name        string  `json:"name"`
    Description *string `json:"description"`
}

type UpdatePermissionRequest struct {
    Name        string  `json:"name"`
    Description *string `json:"description"`
}
