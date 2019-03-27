package auth

import (
    "net/http"
    "github.com/labstack/echo"
)

func UrlHeader(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        url := c.Request().Header.Get("url")
        if url != "" {
            return next(c)
        } else {
            return c.String(http.StatusBadRequest, "Invalid callback url")
        }
    }
}
